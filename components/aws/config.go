package aws

import (
	"fmt"
	"time"

	sdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/kihamo/shadow/components/config"
)

const (
	ConfigAwsKey                          = "aws.key"
	ConfigAwsSecret                       = "aws.secret"
	ConfigAwsRegion                       = "aws.region"
	ConfigAwsLogLevel                     = "aws.log_level"
	ConfigAwsRunUpdatersOnStartup         = "aws.run_updaters_on_startup"
	ConfigAwsUpdaterApplicationsDuration  = "aws.updater_applications_duration"
	ConfigAwsUpdaterSubscriptionsDuration = "aws.updater_subscriptions_duration"
	ConfigAwsUpdaterTopicsDuration        = "aws.updater_topics_duration"
	ConfigAwsSesFromEmail                 = "aws.ses.from_email"
)

func (c *Component) GetConfigVariables() []config.Variable {
	return []config.Variable{
		{
			Key:      ConfigAwsKey,
			Usage:    "AWS access key ID",
			Type:     config.ValueTypeString,
			Editable: true,
		},
		{
			Key:      ConfigAwsSecret,
			Usage:    "AWS secret access key",
			Type:     config.ValueTypeString,
			Editable: true,
		},
		{
			Key:      ConfigAwsRegion,
			Usage:    "AWS region",
			Type:     config.ValueTypeString,
			Editable: true,
		},
		{
			Key:     ConfigAwsLogLevel,
			Default: sdk.LogOff,
			Usage: fmt.Sprintf("AWS log level: %d - LogOff, %d - LogDebug, %d - LogDebugWithSigning, %d - LogDebugWithHTTPBody, %d - LogDebugWithRequestRetries, %d - LogDebugWithRequestErrors",
				sdk.LogOff,
				sdk.LogDebug,
				sdk.LogDebugWithSigning,
				sdk.LogDebugWithHTTPBody,
				sdk.LogDebugWithRequestRetries,
				sdk.LogDebugWithRequestErrors),
			Type:     config.ValueTypeUint,
			Editable: true,
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
		{
			Key:      ConfigAwsSesFromEmail,
			Usage:    "Email for from field in letters",
			Type:     config.ValueTypeString,
			Editable: true,
		},
	}
}

func (c *Component) GetConfigWatchers() map[string][]config.Watcher {
	return map[string][]config.Watcher{
		ConfigAwsKey:                          {c.watchAwsCredentials},
		ConfigAwsSecret:                       {c.watchAwsCredentials},
		ConfigAwsRegion:                       {c.watchAwsRegion},
		ConfigAwsLogLevel:                     {c.watchAwsLogLevel},
		ConfigAwsUpdaterApplicationsDuration:  {c.watchAwsUpdaterApplicationsDuration},
		ConfigAwsUpdaterSubscriptionsDuration: {c.watchAwsUpdaterSubscriptionsDuration},
		ConfigAwsUpdaterTopicsDuration:        {c.watchAwsUpdaterTopicsDuration},
	}
}

func (c *Component) watchAwsCredentials(_ string, newValue interface{}, _ interface{}) {
	config := c.getAwsConfig().WithCredentials(credentials.NewStaticCredentials(c.config.GetString(ConfigAwsKey), c.config.GetString(ConfigAwsSecret), ""))
	c.initAwsConfig(config)
}

func (c *Component) watchAwsRegion(_ string, newValue interface{}, _ interface{}) {
	config := c.getAwsConfig().WithRegion(newValue.(string))
	c.initAwsConfig(config)
}

func (c *Component) watchAwsLogLevel(_ string, newValue interface{}, _ interface{}) {
	config := c.getAwsConfig().WithLogLevel(sdk.LogLevelType(newValue.(uint)))
	c.initAwsConfig(config)
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
