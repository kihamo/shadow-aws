package aws

import (
	"sync"
	"time"

	sdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow/components/config"
	"github.com/kihamo/shadow/components/logger"
)

type Component struct {
	application shadow.Application

	config *config.Component
	logger logger.Logger

	services  map[string]interface{}
	awsConfig *sdk.Config
	mutex     sync.RWMutex

	applications       map[string]AwsSnsApplication
	applicationsTicker chan time.Duration
	applicationsRun    chan bool

	subscriptions       []*sns.Subscription
	subscriptionsTicker chan time.Duration
	subscriptionsRun    chan bool

	topics       []*sns.Topic
	topicsTicker chan time.Duration
	topicsRun    chan bool
}

func (c *Component) GetName() string {
	return "aws"
}

func (c *Component) GetVersion() string {
	return "1.0.1"
}

func (c *Component) Init(a shadow.Application) error {
	c.application = a

	cmpConfig, err := a.GetComponent("config")
	if err != nil {
		return err
	}
	c.config = cmpConfig.(*config.Component)

	c.services = map[string]interface{}{}

	c.applicationsTicker = make(chan time.Duration)
	c.applicationsRun = make(chan bool)

	c.subscriptionsTicker = make(chan time.Duration)
	c.subscriptionsRun = make(chan bool)

	c.topicsTicker = make(chan time.Duration)
	c.topicsRun = make(chan bool)

	c.applications = map[string]AwsSnsApplication{}
	c.subscriptions = []*sns.Subscription{}
	c.topics = []*sns.Topic{}

	return nil
}

func (c *Component) Run() error {
	c.logger = logger.NewOrNop(c.GetName(), c.application)

	awsConfig := sdk.NewConfig().
		WithCredentials(credentials.NewStaticCredentials(c.config.GetString(ConfigAwsKey), c.config.GetString(ConfigAwsSecret), "")).
		WithRegion(c.config.GetString(ConfigAwsRegion)).
		WithLogLevel(sdk.LogLevelType(c.config.GetUint(ConfigAwsLogLevel))).
		WithLogger(c.logger)

	fields := map[string]interface{}{
		"region": *awsConfig.Region,
	}

	credentials, err := awsConfig.Credentials.Get()
	if err == nil {
		fields["key"] = credentials.AccessKeyID
		fields["secret"] = credentials.SecretAccessKey
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

func (c *Component) GetApplications() []AwsSnsApplication {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	applications := make([]AwsSnsApplication, len(c.applications))

	i := 0
	for _, application := range c.applications {
		applications[i] = application
		i++
	}

	return applications
}

func (c *Component) GetSubscriptions() []*sns.Subscription {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	subscriptions := make([]*sns.Subscription, len(c.subscriptions))
	copy(subscriptions, c.subscriptions)

	return subscriptions
}

func (c *Component) GetTopics() []*sns.Topic {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	topics := make([]*sns.Topic, len(c.topics))
	copy(topics, c.topics)

	return topics
}

func (c *Component) GetSNS() *sns.SNS {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.services["sns"]; !ok {
		c.services["sns"] = sns.New(session.New(c.awsConfig))
	}

	return c.services["sns"].(*sns.SNS)
}

func (c *Component) GetServices() map[string]interface{} {
	c.mutex.RLock()
	c.mutex.RUnlock()

	return c.services
}
