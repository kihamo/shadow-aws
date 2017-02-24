package aws

import (
	"time"

	"github.com/kihamo/shadow/components/config"
)

const (
	ConfigAwsKey                          = "aws.key"
	ConfigAwsSecret                       = "aws.secret"
	ConfigAwsRegion                       = "aws.region"
	ConfigAwsRunUpdatersOnStartup         = "aws.run_updaters_on_startup"
	ConfigAwsUpdaterApplicationsDuration  = "aws.updater_applications_duration"
	ConfigAwsUpdaterSubscriptionsDuration = "aws.updater_subscriptions_duration"
	ConfigAwsUpdaterTopicsDuration        = "aws.updater_topics_duration"
)

func (c *Component) GetConfigVariables() []config.Variable {
	return []config.Variable{
		{
			Key:   ConfigAwsKey,
			Usage: "AWS access key ID",
			Type:  config.ValueTypeString,
		},
		{
			Key:   ConfigAwsSecret,
			Usage: "AWS secret access key",
			Type:  config.ValueTypeString,
		},
		{
			Key:   ConfigAwsRegion,
			Usage: "AWS region",
			Type:  config.ValueTypeString,
		},
		{
			Key:     ConfigAwsRunUpdatersOnStartup,
			Default: true,
			Usage:   "Run updater jobs on startup",
			Type:    config.ValueTypeBool,
		},
		{
			Key:      ConfigAwsUpdaterApplicationsDuration,
			Default:  "10m",
			Usage:    "Duration for AWS applications updater",
			Type:     config.ValueTypeDuration,
			Editable: true,
		},
		{
			Key:      ConfigAwsUpdaterSubscriptionsDuration,
			Default:  "10m",
			Usage:    "Duration for AWS subscriptions updater",
			Type:     config.ValueTypeDuration,
			Editable: true,
		},
		{
			Key:      ConfigAwsUpdaterTopicsDuration,
			Default:  "10m",
			Usage:    "Duration for AWS topics updater",
			Type:     config.ValueTypeDuration,
			Editable: true,
		},
	}
}

func (c *Component) GetConfigWatchers() map[string][]config.Watcher {
	return map[string][]config.Watcher{
		ConfigAwsUpdaterApplicationsDuration:  {c.watchAwsUpdaterApplicationsDuration},
		ConfigAwsUpdaterSubscriptionsDuration: {c.watchAwsUpdaterSubscriptionsDuration},
		ConfigAwsUpdaterTopicsDuration:        {c.watchAwsUpdaterTopicsDuration},
	}
}

func (c *Component) watchAwsUpdaterApplicationsDuration(_ string, newValue interface{}, _ interface{}) {
	c.applicationsTicker <- newValue.(time.Duration)
}

func (c *Component) watchAwsUpdaterSubscriptionsDuration(_ string, newValue interface{}, _ interface{}) {
	c.subscriptionsTicker <- newValue.(time.Duration)
}

func (c *Component) watchAwsUpdaterTopicsDuration(_ string, newValue interface{}, _ interface{}) {
	c.topicsTicker <- newValue.(time.Duration)
}
