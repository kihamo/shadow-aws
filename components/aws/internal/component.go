package internal

import (
	"sync"
	"time"

	sdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow-aws/components/aws"
	"github.com/kihamo/shadow/components/config"
	"github.com/kihamo/shadow/components/dashboard"
	"github.com/kihamo/shadow/components/i18n"
	"github.com/kihamo/shadow/components/logger"
	"github.com/kihamo/shadow/components/metrics"
)

const (
	ServiceSNS = "sns"
	ServiceSES = "ses"
)

type Component struct {
	application shadow.Application

	config config.Component
	logger logger.Logger

	routes []dashboard.Route

	services       map[string]interface{}
	awsConfig      *sdk.Config
	mutex          sync.RWMutex
	metricsEnabled bool

	applications       map[string]aws.SnsApplication
	applicationsTicker chan time.Duration
	applicationsRun    chan struct{}

	subscriptions       []*sns.Subscription
	subscriptionsTicker chan time.Duration
	subscriptionsRun    chan struct{}

	topics       []*sns.Topic
	topicsTicker chan time.Duration
	topicsRun    chan struct{}
}

func (c *Component) Name() string {
	return aws.ComponentName
}

func (c *Component) Version() string {
	return aws.ComponentVersion + "/" + sdk.SDKVersion
}

func (c *Component) Dependencies() []shadow.Dependency {
	return []shadow.Dependency{
		{
			Name:     config.ComponentName,
			Required: true,
		},
		{
			Name: i18n.ComponentName,
		},
		{
			Name: logger.ComponentName,
		},
		{
			Name: metrics.ComponentName,
		},
	}
}

func (c *Component) Init(a shadow.Application) error {
	c.application = a

	c.config = a.GetComponent(config.ComponentName).(config.Component)

	c.services = map[string]interface{}{}
	c.metricsEnabled = a.HasComponent(metrics.ComponentName)

	c.applicationsTicker = make(chan time.Duration)
	c.applicationsRun = make(chan struct{})

	c.subscriptionsTicker = make(chan time.Duration)
	c.subscriptionsRun = make(chan struct{})

	c.topicsTicker = make(chan time.Duration)
	c.topicsRun = make(chan struct{})

	c.applications = map[string]aws.SnsApplication{}
	c.subscriptions = []*sns.Subscription{}
	c.topics = []*sns.Topic{}

	return nil
}

func (c *Component) Run() error {
	c.logger = logger.NewOrNop(c.Name(), c.application)

	awsConfig := sdk.NewConfig().
		WithCredentials(credentials.NewStaticCredentials(c.config.String(aws.ConfigKey), c.config.String(aws.ConfigSecret), "")).
		WithRegion(c.config.String(aws.ConfigRegion)).
		WithLogLevel(sdk.LogLevelType(c.config.Uint(aws.ConfigLogLevel))).
		WithLogger(c.logger)

	fields := map[string]interface{}{
		"region": *awsConfig.Region,
	}

	creds, err := awsConfig.Credentials.Get()
	if err == nil {
		fields["key"] = creds.AccessKeyID
		fields["secret"] = creds.SecretAccessKey
	}

	c.logger.Info("Connect AWS", fields)

	c.initAwsConfig(awsConfig)
	c.loadUpdaters()

	return nil
}

func (c *Component) initAwsConfig(config *sdk.Config) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.awsConfig = config
	c.services = map[string]interface{}{}
}

func (c *Component) getAwsConfig() *sdk.Config {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.awsConfig
}

func (c *Component) GetServices() map[string]interface{} {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.services
}
