package internal

import (
	"time"

	sdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/kihamo/shadow-aws/components/aws"
	"github.com/kihamo/shadow/components/config"
)

func (c *Component) ConfigVariables() []config.Variable {
	return []config.Variable{
		config.NewVariable(
			aws.ConfigKey,
			config.ValueTypeString,
			nil,
			"AWS access key ID",
			true,
			"AWS config",
			nil,
			nil),
		config.NewVariable(
			aws.ConfigSecret,
			config.ValueTypeString,
			nil,
			"AWS secret access key",
			true,
			"AWS config",
			nil,
			nil),
		config.NewVariable(
			aws.ConfigRegion,
			config.ValueTypeString,
			nil,
			"AWS region",
			true,
			"AWS config",
			nil,
			nil),
		config.NewVariable(
			aws.ConfigLogLevel,
			config.ValueTypeUint,
			sdk.LogOff,
			"AWS log level",
			true,
			"AWS config",
			[]string{config.ViewEnum},
			map[string]interface{}{
				config.ViewOptionEnumOptions: [][]interface{}{
					{sdk.LogOff, "LogOff"},
					{sdk.LogDebug, "LogDebug"},
					{sdk.LogDebugWithSigning, "LogDebugWithSigning"},
					{sdk.LogDebugWithHTTPBody, "LogDebugWithHTTPBody"},
					{sdk.LogDebugWithRequestRetries, "LogDebugWithRequestRetries"},
					{sdk.LogDebugWithRequestErrors, "LogDebugWithRequestErrors"},
				},
			}),
		config.NewVariable(
			aws.ConfigRunUpdatersOnStartup,
			config.ValueTypeBool,
			true,
			"Run updater jobs on startup",
			false,
			"Updaters",
			nil,
			nil),
		config.NewVariable(
			aws.ConfigUpdaterApplicationsDuration,
			config.ValueTypeDuration,
			"10m",
			"Duration for AWS applications updater",
			true,
			"Updaters",
			nil,
			nil),
		config.NewVariable(
			aws.ConfigUpdaterSubscriptionsDuration,
			config.ValueTypeDuration,
			"10m",
			"Duration for AWS subscriptions updater",
			true,
			"Updaters",
			nil,
			nil),
		config.NewVariable(
			aws.ConfigUpdaterTopicsDuration,
			config.ValueTypeDuration,
			"10m",
			"Duration for AWS topics updater",
			true,
			"Updaters",
			nil,
			nil),
		config.NewVariable(
			aws.ConfigSesFromEmail,
			config.ValueTypeString,
			nil,
			"Email for from field in letters",
			true,
			"SES",
			nil,
			nil),
		config.NewVariable(
			aws.ConfigSesFromName,
			config.ValueTypeString,
			nil,
			"Name for from field in letters",
			true,
			"SES",
			nil,
			nil),
	}
}

func (c *Component) ConfigWatchers() []config.Watcher {
	return []config.Watcher{
		config.NewWatcher(aws.ComponentName, []string{aws.ConfigKey, aws.ConfigSecret}, c.watchCredentials),
		config.NewWatcher(aws.ComponentName, []string{aws.ConfigRegion}, c.watchRegion),
		config.NewWatcher(aws.ComponentName, []string{aws.ConfigLogLevel}, c.watchLogLevel),
		config.NewWatcher(aws.ComponentName, []string{aws.ConfigUpdaterApplicationsDuration}, c.watchUpdaterApplicationsDuration),
		config.NewWatcher(aws.ComponentName, []string{aws.ConfigUpdaterSubscriptionsDuration}, c.watchUpdaterSubscriptionsDuration),
		config.NewWatcher(aws.ComponentName, []string{aws.ConfigUpdaterTopicsDuration}, c.watchUpdaterTopicsDuration),
	}
}

func (c *Component) watchCredentials(_ string, newValue interface{}, _ interface{}) {
	cfg := c.getAwsConfig().WithCredentials(credentials.NewStaticCredentials(c.config.String(aws.ConfigKey), c.config.String(aws.ConfigSecret), ""))
	c.initAwsConfig(cfg)
}

func (c *Component) watchRegion(_ string, newValue interface{}, _ interface{}) {
	cfg := c.getAwsConfig().WithRegion(newValue.(string))
	c.initAwsConfig(cfg)
}

func (c *Component) watchLogLevel(_ string, newValue interface{}, _ interface{}) {
	cfg := c.getAwsConfig().WithLogLevel(sdk.LogLevelType(newValue.(uint)))
	c.initAwsConfig(cfg)
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
