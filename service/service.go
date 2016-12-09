package service

import (
	"sync"
	"time"

	sdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/kihamo/go-workers/task"
	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow-aws/resource/aws"
	"github.com/kihamo/shadow/resource/config"
	"github.com/kihamo/shadow/resource/logger"
	"github.com/kihamo/shadow/resource/workers"
	"github.com/rs/xlog"
)

type AwsSnsApplication struct {
	Arn                       string
	AwsAttributes             map[string]*string
	Enabled                   bool
	EndpointsCount            int
	EndpointsEnabledCount     int
	CertificateExpirationDate *time.Time
	LastUpdate                time.Time
}

func (a AwsSnsApplication) IsIAM() bool {
	if _, ok := a.AwsAttributes["SuccessFeedbackRoleArn"]; !ok {
		return false
	}

	if _, ok := a.AwsAttributes["FailureFeedbackRoleArn"]; !ok {
		return false
	}

	return true
}

func (a AwsSnsApplication) GetEnabledCount() int {
	if a.EndpointsEnabledCount <= 0 {
		return 0
	}

	return a.EndpointsEnabledCount
}

func (a AwsSnsApplication) GetEnabledPercent() int {
	if a.EndpointsCount <= 0 {
		return 0
	}

	return (100 * a.EndpointsEnabledCount) / a.EndpointsCount
}

func (a AwsSnsApplication) GetDisabledCount() int {
	if a.EndpointsCount <= 0 || a.EndpointsEnabledCount <= 0 {
		return 0
	}

	return a.EndpointsCount - a.EndpointsEnabledCount
}

func (a AwsSnsApplication) GetDisabledPercent() int {
	if a.EndpointsCount <= 0 {
		return 0
	}

	return 100 - a.GetEnabledPercent()
}

type AwsService struct {
	application *shadow.Application
	config      *config.Resource
	workers     *workers.Resource
	logger      xlog.Logger

	aws   *aws.Resource
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

	resourceConfig, err := a.GetResource("config")
	if err != nil {
		return err
	}
	s.config = resourceConfig.(*config.Resource)

	resourceAws, err := a.GetResource("aws")
	if err != nil {
		return err
	}
	s.aws = resourceAws.(*aws.Resource)

	return nil
}

func (s *AwsService) Run() error {
	if resourceLogger, err := s.application.GetResource("logger"); err == nil {
		s.logger = resourceLogger.(*logger.Resource).Get(s.GetName())
	} else {
		s.logger = xlog.NopLogger
	}

	s.applications = map[string]AwsSnsApplication{}

	if s.application.HasResource("workers") {
		resourceWorkers, _ := s.application.GetResource("workers")
		s.workers = resourceWorkers.(*workers.Resource)
		runOnStartUp := s.config.GetBool("aws.run_updater_on_startup")

		var t *task.Task

		t = task.NewTask(s.getApplicationsJob)
		t.SetName("aws.updater.applications")
		if !runOnStartUp {
			t.SetDuration(s.config.GetDuration("aws.updater_applications_duration"))
		}
		s.workers.AddTask(t)

		t = task.NewTask(s.getSubscriptionsJob)
		t.SetName("aws.updater.subscriptions")
		if !runOnStartUp {
			t.SetDuration(s.config.GetDuration("aws.updater_subscriptions_duration"))
		}
		s.workers.AddTask(t)

		t = task.NewTask(s.getTopicsJob)
		t.SetName("aws.updater.topics")
		if !runOnStartUp {
			t.SetDuration(s.config.GetDuration("aws.updater_topics_duration"))
		}
		s.workers.AddTask(t)
	}

	return nil
}

func (s *AwsService) getApplicationsJob(attempts int64, _ chan bool, args ...interface{}) (int64, time.Duration, interface{}, error) {
	lastUpdate := time.Now().UTC()
	params := &sns.ListPlatformApplicationsInput{}

	err := s.aws.GetSNS().ListPlatformApplicationsPages(params, func(p *sns.ListPlatformApplicationsOutput, lastPage bool) bool {
		s.mutex.Lock()
		defer s.mutex.Unlock()

		var app AwsSnsApplication

		for _, a := range p.PlatformApplications {
			arn := sdk.StringValue(a.PlatformApplicationArn)

			if _, ok := s.applications[arn]; ok {
				app = s.applications[arn]
				app.LastUpdate = lastUpdate
			} else {
				app = AwsSnsApplication{
					Arn:                   arn,
					AwsAttributes:         a.Attributes,
					Enabled:               true,
					EndpointsCount:        -1,
					EndpointsEnabledCount: -1,
					LastUpdate:            lastUpdate,
				}
			}

			if dateRaw, ok := a.Attributes["AppleCertificateExpirationDate"]; ok {
				if dateValue, err := time.Parse(time.RFC3339, sdk.StringValue(dateRaw)); err == nil {
					app.CertificateExpirationDate = &dateValue
				}
			}

			if dateRaw, ok := a.Attributes["Enabled"]; ok && sdk.StringValue(dateRaw) == "false" {
				app.Enabled = false
			}

			s.applications[arn] = app
		}

		return !lastPage
	})

	if err != nil {
		s.logger.Errorf("Update applications error %s", err.Error())
	}

	s.mutex.Lock()

	if metricApplicationsTotal != nil {
		metricApplicationsTotal.Set(float64(len(s.applications)))
	}

	for key, application := range s.applications {
		if application.LastUpdate.Before(lastUpdate) {
			delete(s.applications, key)
		}
	}
	s.mutex.Unlock()

	if attempts == 1 {
		s.workers.AddNamedTaskByFunc("aws.updater.endpoints", s.getEndpointsConsolidatedJob)
	}

	return -1, s.config.GetDuration("aws.updater_applications_duration"), nil, err
}

func (s *AwsService) getEndpointsConsolidatedJob(attempts int64, _ chan bool, args ...interface{}) (int64, time.Duration, interface{}, error) {
	applications := s.GetApplications()
	bulkCount := s.config.GetInt("aws.updater_endpoints_bulk")

	for i := 0; i < len(applications); i += bulkCount {
		stop := i + bulkCount
		if stop > len(applications) {
			stop = len(applications)
		}

		s.workers.AddTaskByFunc(s.getEndpointsJob, applications[i:stop])
	}

	return -1, s.config.GetDuration("aws.updater_endpoints_duration"), nil, nil
}

func (s *AwsService) getEndpointsJob(attempts int64, _ chan bool, args ...interface{}) (int64, time.Duration, interface{}, error) {
	for _, app := range args[0].([]AwsSnsApplication) {
		params := &sns.ListEndpointsByPlatformApplicationInput{
			PlatformApplicationArn: sdk.String(app.Arn),
		}

		s.mutex.RLock()
		_, ok := s.applications[app.Arn]
		s.mutex.RUnlock()
		if !ok {
			continue
		}

		app.EndpointsCount = 0
		app.EndpointsEnabledCount = 0
		app.LastUpdate = time.Now().UTC()

		err := s.aws.GetSNS().ListEndpointsByPlatformApplicationPages(params, func(p *sns.ListEndpointsByPlatformApplicationOutput, lastPage bool) bool {
			app.EndpointsCount += len(p.Endpoints)

			for _, point := range p.Endpoints {
				if enabled, ok := point.Attributes["Enabled"]; ok && sdk.StringValue(enabled) == "true" {
					app.EndpointsEnabledCount++
				}
			}

			return !lastPage
		})

		if err == nil {
			if metricEndpointsTotal != nil {
				metricEndpointsTotal.With("arn", app.Arn).Set(float64(app.EndpointsCount))
			}

			if metricEndpointsEnabled != nil {
				metricEndpointsEnabled.With("arn", app.Arn).Set(float64(app.EndpointsEnabledCount))
			}
		}

		s.mutex.Lock()

		if _, ok := s.applications[app.Arn]; ok {
			s.applications[app.Arn] = app
		}

		if err != nil {
			s.logger.Errorf("Update apn %s", app.Arn, xlog.F{
				"application.ednpoints":         app.EndpointsCount,
				"application.ednpoints-enabled": app.EndpointsEnabledCount,
				"error": err.Error(),
			})
		} else {
			s.logger.Infof("Update apn %s", app.Arn, xlog.F{
				"application.ednpoints":         app.EndpointsCount,
				"application.ednpoints-enabled": app.EndpointsEnabledCount,
			})
		}

		s.mutex.Unlock()
	}

	return 1, 0, nil, nil
}

func (s *AwsService) getSubscriptionsJob(attempts int64, _ chan bool, args ...interface{}) (int64, time.Duration, interface{}, error) {
	subscriptions := []*sns.Subscription{}
	params := &sns.ListSubscriptionsInput{}

	err := s.aws.GetSNS().ListSubscriptionsPages(params, func(p *sns.ListSubscriptionsOutput, lastPage bool) bool {
		subscriptions = append(subscriptions, p.Subscriptions...)
		return !lastPage
	})

	if err != nil {
		s.logger.Errorf("Update subscriptions error %s", err.Error())
	} else if metricSubscriptionsTotal != nil {
		metricSubscriptionsTotal.Set(float64(len(subscriptions)))
	}

	s.mutex.Lock()
	s.subscriptions = subscriptions
	s.mutex.Unlock()

	return -1, s.config.GetDuration("aws.updater_subscriptions_duration"), nil, err
}

func (s *AwsService) getTopicsJob(attempts int64, _ chan bool, args ...interface{}) (int64, time.Duration, interface{}, error) {
	topics := []*sns.Topic{}
	params := &sns.ListTopicsInput{}

	err := s.aws.GetSNS().ListTopicsPages(params, func(p *sns.ListTopicsOutput, lastPage bool) bool {
		topics = append(topics, p.Topics...)
		return !lastPage
	})

	if err != nil {
		s.logger.Errorf("Update topics error %s", err.Error())
	} else if metricTopicsTotal != nil {
		metricTopicsTotal.Set(float64(len(topics)))
	}

	s.mutex.Lock()
	s.topics = topics
	s.mutex.Unlock()

	return -1, s.config.GetDuration("aws.updater_topics_duration"), nil, err
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
