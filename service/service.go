package service

import (
	"sync"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow-aws/resource"
	r "github.com/kihamo/shadow/resource"
)

const (
	updateEndpointsBulk = 5
)

type AwsSnsApplication struct {
	Arn                       string
	EndpointsCount            int
	EndpointsEnabledCount     int
	CertificateExpirationDate *time.Time
	LastUpdate                time.Time
}

type AwsService struct {
	application *shadow.Application

	Aws    *resource.Aws
	SNS    *sns.SNS
	logger *logrus.Entry

	mutex sync.RWMutex

	applications  map[string]AwsSnsApplication
	subscriptions []*sns.Subscription
	topics        []*sns.Topic
}

func (s *AwsService) GetName() string {
	return "aws"
}

func (s *AwsService) Init(a *shadow.Application) error {
	s.application = a

	resourceLogger, err := a.GetResource("logger")
	if err != nil {
		return err
	}
	s.logger = resourceLogger.(*r.Logger).Get(s.GetName())

	resourceAws, err := a.GetResource("aws")
	if err != nil {
		return err
	}
	s.Aws = resourceAws.(*resource.Aws)

	s.SNS = s.Aws.GetSNS()

	return nil
}

func (s *AwsService) Run() error {
	s.applications = map[string]AwsSnsApplication{}

	if s.application.HasResource("workers") {
		resourceWorkers, _ := s.application.GetResource("workers")
		workers := resourceWorkers.(*r.Workers)

		workers.AddNamedTaskByFunc("aws.updater.applications", s.getApplicationsJob)
		workers.AddNamedTaskByFunc("aws.updater.subscriptions", s.getSubscriptionsJob)
		workers.AddNamedTaskByFunc("aws.updater.topics", s.getTopicsJob)
	}

	return nil
}

func (s *AwsService) getApplicationsJob(attempts int64, _ chan bool, args ...interface{}) (int64, time.Duration, interface{}, error) {
	lastUpdate := time.Now()
	params := &sns.ListPlatformApplicationsInput{}

	err := s.SNS.ListPlatformApplicationsPages(params, func(p *sns.ListPlatformApplicationsOutput, lastPage bool) bool {
		for _, a := range p.PlatformApplications {
			app := AwsSnsApplication{
				Arn:                   aws.StringValue(a.PlatformApplicationArn),
				EndpointsCount:        -1,
				EndpointsEnabledCount: -1,
				LastUpdate:            lastUpdate,
			}

			if dateRaw, ok := a.Attributes["AppleCertificateExpirationDate"]; ok {
				if dateValue, err := time.Parse(time.RFC3339, aws.StringValue(dateRaw)); err == nil {
					app.CertificateExpirationDate = &dateValue
				}
			}

			s.mutex.Lock()
			s.applications[app.Arn] = app
			s.mutex.Unlock()
		}

		return !lastPage
	})

	if err != nil {
		s.logger.Errorf("Update applications error %s", err.Error())
	}

	s.mutex.Lock()
	for key, application := range s.applications {
		if application.LastUpdate.Before(lastUpdate) {
			delete(s.applications, key)
		}
	}
	s.mutex.Unlock()

	if attempts == 1 {
		resourceWorkers, _ := s.application.GetResource("workers")
		resourceWorkers.(*r.Workers).AddNamedTaskByFunc("aws.updater.endpoints", s.getEndpointsConsolidatedJob)
	}

	return -1, time.Minute * 10, nil, err
}

func (s *AwsService) getEndpointsConsolidatedJob(attempts int64, _ chan bool, args ...interface{}) (int64, time.Duration, interface{}, error) {
	applications := s.GetApplications()

	resourceWorkers, _ := s.application.GetResource("workers")
	workers := resourceWorkers.(*r.Workers)

	for i := 0; i < len(applications); i += updateEndpointsBulk {
		stop := i + updateEndpointsBulk
		if stop > len(applications) {
			stop = len(applications)
		}

		workers.AddTaskByFunc(s.getEndpointsJob, applications[i:stop])
	}

	return -1, time.Minute * 60, nil, nil
}

func (s *AwsService) getEndpointsJob(attempts int64, _ chan bool, args ...interface{}) (int64, time.Duration, interface{}, error) {
	for _, app := range args[0].([]AwsSnsApplication) {
		params := &sns.ListEndpointsByPlatformApplicationInput{
			PlatformApplicationArn: aws.String(app.Arn),
		}

		s.mutex.RLock()
		_, ok := s.applications[app.Arn]
		s.mutex.RUnlock()
		if !ok {
			continue
		}

		app.EndpointsCount = 0
		app.EndpointsEnabledCount = 0
		app.LastUpdate = time.Now()

		err := s.SNS.ListEndpointsByPlatformApplicationPages(params, func(p *sns.ListEndpointsByPlatformApplicationOutput, lastPage bool) bool {
			app.EndpointsCount += len(p.Endpoints)

			for _, point := range p.Endpoints {
				if enabled, ok := point.Attributes["Enabled"]; ok && aws.StringValue(enabled) == "true" {
					app.EndpointsEnabledCount++
				}
			}

			return !lastPage
		})

		s.mutex.Lock()

		if _, ok := s.applications[app.Arn]; ok {
			s.applications[app.Arn] = app
		}

		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"application.ednpoints":         app.EndpointsCount,
				"application.ednpoints-enabled": app.EndpointsEnabledCount,
				"error": err.Error(),
			}).Errorf("Update apn %s", app.Arn)
		} else {
			s.logger.WithFields(logrus.Fields{
				"application.ednpoints":         app.EndpointsCount,
				"application.ednpoints-enabled": app.EndpointsEnabledCount,
			}).Infof("Update apn %s", app.Arn)
		}

		s.mutex.Unlock()
	}

	return 1, 0, nil, nil
}

func (s *AwsService) getSubscriptionsJob(attempts int64, _ chan bool, args ...interface{}) (int64, time.Duration, interface{}, error) {
	subscriptions := []*sns.Subscription{}
	params := &sns.ListSubscriptionsInput{}

	err := s.SNS.ListSubscriptionsPages(params, func(p *sns.ListSubscriptionsOutput, lastPage bool) bool {
		subscriptions = append(subscriptions, p.Subscriptions...)
		return !lastPage
	})

	if err != nil {
		s.logger.Errorf("Update subscriptions error %s", err.Error())
	}

	s.mutex.Lock()
	s.subscriptions = subscriptions
	s.mutex.Unlock()

	return -1, time.Minute * 10, nil, err
}

func (s *AwsService) getTopicsJob(attempts int64, _ chan bool, args ...interface{}) (int64, time.Duration, interface{}, error) {
	topics := []*sns.Topic{}
	params := &sns.ListTopicsInput{}

	err := s.SNS.ListTopicsPages(params, func(p *sns.ListTopicsOutput, lastPage bool) bool {
		topics = append(topics, p.Topics...)
		return !lastPage
	})

	if err != nil {
		s.logger.Errorf("Update topics error %s", err.Error())
	}

	s.mutex.Lock()
	s.topics = topics
	s.mutex.Unlock()

	return -1, time.Minute * 10, nil, err
}

func (s *AwsService) GetApplications() []AwsSnsApplication {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	applications := make([]AwsSnsApplication, len(s.applications))

	i := 0
	for _, application := range s.applications {
		applications[i] = application
		i++
	}

	return applications
}

func (s *AwsService) GetSubscriptions() []*sns.Subscription {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.subscriptions
}

func (s *AwsService) GetTopics() []*sns.Topic {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.topics
}
