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
		workers.(*r.Workers).AddNamedTaskByFunc("aws.updater", s.getStatsJob)
	}

	return nil
}

func (s *AwsService) getStatsJob(attempts int64, _ chan bool, args ...interface{}) (int64, time.Duration, interface{}, error) {
	var stop bool

	// applications
	applications := []AwsSnsApplication{}
	paramsApplications := &sns.ListPlatformApplicationsInput{}
	for !stop {
		responseApps, err := s.SNS.ListPlatformApplications(paramsApplications)
		if err == nil {
			for _, a := range responseApps.PlatformApplications {
				snsApplication := AwsSnsApplication{
					Arn: aws.StringValue(a.PlatformApplicationArn),
				}

				if dateRaw, ok := a.Attributes["AppleCertificateExpirationDate"]; ok {
					if dateValue, err := time.Parse(time.RFC3339, aws.StringValue(dateRaw)); err == nil {
						snsApplication.CertificateExpirationDate = &dateValue
					}
				}

				applications = append(applications, snsApplication)
			}

			if responseApps.NextToken != nil {
				paramsApplications.NextToken = responseApps.NextToken
			} else {
				stop = true
			}
		} else {
			return -1, time.Minute * 10, nil, err
		}
	}

	// subscriptions
	stop = false
	subscriptions := []*sns.Subscription{}
	paramsSubscriptions := &sns.ListSubscriptionsInput{}
	for !stop {
		responseSubscriptions, err := s.SNS.ListSubscriptions(paramsSubscriptions)
		if err == nil {
			subscriptions = append(subscriptions, responseSubscriptions.Subscriptions...)

			if responseSubscriptions.NextToken != nil {
				paramsSubscriptions.NextToken = responseSubscriptions.NextToken
			} else {
				stop = true
			}
		} else {
			return -1, time.Minute * 10, nil, err
		}
	}

	// topics
	stop = false
	topics := []*sns.Topic{}
	paramsTopics := &sns.ListTopicsInput{}
	for !stop {
		responseTopics, err := s.SNS.ListTopics(paramsTopics)
		if err == nil {
			topics = append(topics, responseTopics.Topics...)

			if responseTopics.NextToken != nil {
				paramsTopics.NextToken = responseTopics.NextToken
			} else {
				stop = true
			}
		} else {
			return -1, time.Minute * 10, nil, err
		}
	}

	s.mutex.Lock()
	s.applications = applications
	s.subscriptions = subscriptions
	s.topics = topics
	s.mutex.Unlock()

	return -1, time.Hour, nil, nil
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
