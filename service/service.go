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

type AwsSnsApplication struct {
	Arn                       string
	EndpointsCount            int64
	CertificateExpirationDate *time.Time
}

type AwsService struct {
	application *shadow.Application

	Aws    *resource.Aws
	SNS    *sns.SNS
	logger *logrus.Entry

	mutex sync.RWMutex

	applications  []AwsSnsApplication
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
	if s.application.HasResource("workers") {
		workers, _ := s.application.GetResource("workers")
		workers.(*r.Workers).AddNamedTaskByFunc("aws.updater.applications", s.getApplicationsJob)
		workers.(*r.Workers).AddNamedTaskByFunc("aws.updater.subscriptions", s.getSubscriptionsJob)
		workers.(*r.Workers).AddNamedTaskByFunc("aws.updater.topics", s.getTopicsJob)
	}

	return nil
}

func (s *AwsService) getApplicationsJob(attempts int64, _ chan bool, args ...interface{}) (int64, time.Duration, interface{}, error) {
	applications := []AwsSnsApplication{}
	params := &sns.ListPlatformApplicationsInput{}

	err := s.SNS.ListPlatformApplicationsPages(params, func(p *sns.ListPlatformApplicationsOutput, lastPage bool) bool {
		for _, a := range p.PlatformApplications {
			application := AwsSnsApplication{
				Arn:            aws.StringValue(a.PlatformApplicationArn),
				EndpointsCount: -1,
			}

			if dateRaw, ok := a.Attributes["AppleCertificateExpirationDate"]; ok {
				if dateValue, err := time.Parse(time.RFC3339, aws.StringValue(dateRaw)); err == nil {
					application.CertificateExpirationDate = &dateValue
				}
			}

			applications = append(applications, application)
		}

		return !lastPage
	})

	s.mutex.Lock()
	s.applications = applications
	s.mutex.Unlock()

	for i, application := range applications {
		params := &sns.ListEndpointsByPlatformApplicationInput{
			PlatformApplicationArn: aws.String(application.Arn),
		}

		s.SNS.ListEndpointsByPlatformApplicationPages(params, func(p *sns.ListEndpointsByPlatformApplicationOutput, lastPage bool) bool {
			applications[i].EndpointsCount += int64(len(p.Endpoints))
			return !lastPage
		})
	}

	s.mutex.Lock()
	s.applications = applications
	s.mutex.Unlock()

	return -1, time.Minute * 10, nil, err
}

func (s *AwsService) getSubscriptionsJob(attempts int64, _ chan bool, args ...interface{}) (int64, time.Duration, interface{}, error) {
	subscriptions := []*sns.Subscription{}
	params := &sns.ListSubscriptionsInput{}

	err := s.SNS.ListSubscriptionsPages(params, func(p *sns.ListSubscriptionsOutput, lastPage bool) bool {
		subscriptions = append(subscriptions, p.Subscriptions...)
		return !lastPage
	})

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

	s.mutex.Lock()
	s.topics = topics
	s.mutex.Unlock()

	return -1, time.Minute * 10, nil, err
}

func (s *AwsService) GetApplications() []AwsSnsApplication {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.applications
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
