package aws

import (
	"fmt"
	"time"

	sdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/kihamo/shadow/components/config"
)

const (
	ConfigKey                          = ComponentName + ".key"
	ConfigSecret                       = ComponentName + ".secret"
	ConfigRegion                       = ComponentName + ".region"
	ConfigLogLevel                     = ComponentName + ".log_level"
	ConfigRunUpdatersOnStartup         = ComponentName + ".run_updaters_on_startup"
	ConfigUpdaterApplicationsDuration  = ComponentName + ".updater_applications_duration"
	ConfigUpdaterSubscriptionsDuration = ComponentName + ".updater_subscriptions_duration"
	ConfigUpdaterTopicsDuration        = ComponentName + ".updater_topics_duration"
	ConfigSesFromEmail                 = ComponentName + ".ses.from_email"
	ConfigSesFromName                  = ComponentName + ".ses.from_name"
)

func (c *Component) GetConfigVariables() []config.Variable {
	return []config.Variable{
		{
			Key:      ConfigKey,
			Usage:    "AWS access key ID",
			Type:     config.ValueTypeString,
			Editable: true,
		},
		{
			Key:      ConfigSecret,
			Usage:    "AWS secret access key",
			Type:     config.ValueTypeString,
			Editable: true,
		},
		{
			Key:      ConfigRegion,
			Usage:    "AWS region",
			Type:     config.ValueTypeString,
			Editable: true,
		},
		{
			Key:     ConfigLogLevel,
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
			Key:     ConfigRunUpdatersOnStartup,
			Default: true,
			Usage:   "Run updater jobs on startup",
			Type:    config.ValueTypeBool,
		},
		{
			Key:      ConfigUpdaterApplicationsDuration,
			Default:  "10m",
			Usage:    "Duration for AWS applications updater",
			Type:     config.ValueTypeDuration,
			Editable: true,
		},
		{
			Key:      ConfigUpdaterSubscriptionsDuration,
			Default:  "10m",
			Usage:    "Duration for AWS subscriptions updater",
			Type:     config.ValueTypeDuration,
			Editable: true,
		},
		{
			Key:      ConfigUpdaterTopicsDuration,
			Default:  "10m",
			Usage:    "Duration for AWS topics updater",
			Type:     config.ValueTypeDuration,
			Editable: true,
		},
		{
			Key:      ConfigSesFromEmail,
			Usage:    "Email for from field in letters",
			Type:     config.ValueTypeString,
			Editable: true,
		},
		{
			Key:      ConfigSesFromName,
			Usage:    "Name for from field in letters",
			Type:     config.ValueTypeString,
			Editable: true,
		},
	}
}

func (c *Component) GetConfigWatchers() map[string][]config.Watcher {
	return map[string][]config.Watcher{
		ConfigKey:                          {c.watchCredentials},
		ConfigSecret:                       {c.watchCredentials},
		ConfigRegion:                       {c.watchRegion},
		ConfigLogLevel:                     {c.watchLogLevel},
		ConfigUpdaterApplicationsDuration:  {c.watchUpdaterApplicationsDuration},
		ConfigUpdaterSubscriptionsDuration: {c.watchUpdaterSubscriptionsDuration},
		ConfigUpdaterTopicsDuration:        {c.watchUpdaterTopicsDuration},
	}
}

func (c *Component) watchCredentials(_ string, newValue interface{}, _ interface{}) {
	config := c.getAwsConfig().WithCredentials(credentials.NewStaticCredentials(c.config.GetString(ConfigKey), c.config.GetString(ConfigSecret), ""))
	c.initAwsConfig(config)
}

func (c *Component) watchRegion(_ string, newValue interface{}, _ interface{}) {
	config := c.getAwsConfig().WithRegion(newValue.(string))
	c.initAwsConfig(config)
}

func (c *Component) watchLogLevel(_ string, newValue interface{}, _ interface{}) {
	config := c.getAwsConfig().WithLogLevel(sdk.LogLevelType(newValue.(uint)))
	c.initAwsConfig(config)
}

func (c *Component) watchUpdaterApplicationsDuration(_ string, newValue interface{}, _ interface{}) {
	c.applicationsTicker <- newValue.(time.Duration)
}

func (c *Component) watchUpdaterSubscriptionsDuration(_ string, newValue interface{}, _ interface{}) {
	c.subscriptionsTicker <- newValue.(time.Duration)
}

func (c *Component) watchUpdaterTopicsDuration(_ string, newValue interface{}, _ interface{}) {
	c.topicsTicker <- newValue.(time.Duration)
}
