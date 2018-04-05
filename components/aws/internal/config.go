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
		config.NewVariable(aws.ConfigKey, config.ValueTypeString).
			WithUsage("Access key ID").
			WithGroup("AWS config").
			WithEditable(true),
		config.NewVariable(aws.ConfigSecret, config.ValueTypeString).
			WithUsage("Secret access key").
			WithGroup("AWS config").
			WithEditable(true).
			WithView([]string{config.ViewPassword}),
		config.NewVariable(aws.ConfigRegion, config.ValueTypeString).
			WithUsage("Region").
			WithGroup("AWS config").
			WithEditable(true).
			WithView([]string{config.ViewEnum}).
			WithViewOptions(map[string]interface{}{
				config.ViewOptionEnumOptions: [][]interface{}{
					{"us-east-1", "US East (N. Virginia)"},
					{"us-east-2", "US East (Ohio)"},
					{"us-west-1", "US West (N. California)"},
					{"us-west-2", "US West (Oregon)"},
					{"ca-central-1", "Canada (Central)"},
					{"eu-central-1", "EU (Frankfurt)"},
					{"eu-west-1", "EU (Ireland)"},
					{"eu-west-2", "EU (London)"},
					{"eu-west-3", "EU (Paris)"},
					{"ap-northeast-1", "Asia Pacific (Tokyo)"},
					{"ap-northeast-2", "Asia Pacific (Seoul)"},
					{"ap-northeast-3", "Asia Pacific (Osaka-Local)"},
					{"ap-southeast-1", "Asia Pacific (Singapore)"},
					{"ap-southeast-2", "Asia Pacific (Sydney)"},
					{"ap-south-1", "Asia Pacific (Mumbai)"},
					{"sa-east-1", "South America (São Paulo)"},
				},
			}),
		config.NewVariable(aws.ConfigLogLevel, config.ValueTypeUint).
			WithUsage("Log level").
			WithGroup("AWS config").
			WithEditable(true).
			WithDefault(sdk.LogOff).
			WithView([]string{config.ViewEnum}).
			WithViewOptions(map[string]interface{}{
				config.ViewOptionEnumOptions: [][]interface{}{
					{sdk.LogOff, "LogOff"},
					{sdk.LogDebug, "LogDebug"},
					{sdk.LogDebugWithSigning, "LogDebugWithSigning"},
					{sdk.LogDebugWithHTTPBody, "LogDebugWithHTTPBody"},
					{sdk.LogDebugWithRequestRetries, "LogDebugWithRequestRetries"},
					{sdk.LogDebugWithRequestErrors, "LogDebugWithRequestErrors"},
				},
			}),
		config.NewVariable(aws.ConfigRunUpdatersOnStartup, config.ValueTypeBool).
			WithUsage("Run updater on startup").
			WithGroup("Updaters").
			WithDefault(true),
		config.NewVariable(aws.ConfigUpdaterApplicationsDuration, config.ValueTypeDuration).
			WithUsage("Duration for SNS applications updater").
			WithGroup("Updaters").
			WithEditable(true).
			WithDefault("10m"),
		config.NewVariable(aws.ConfigUpdaterSubscriptionsDuration, config.ValueTypeDuration).
			WithUsage("Duration for SNS subscriptions updater").
			WithGroup("Updaters").
			WithEditable(true).
			WithDefault("10m"),
		config.NewVariable(aws.ConfigUpdaterTopicsDuration, config.ValueTypeDuration).
			WithUsage("Duration for SNS topics updater").
			WithGroup("Updaters").
			WithEditable(true).
			WithDefault("10m"),
		config.NewVariable(aws.ConfigSesFromEmail, config.ValueTypeString).
			WithUsage("Mail from address").
			WithGroup("SES").
			WithEditable(true),
		config.NewVariable(aws.ConfigSesFromName, config.ValueTypeString).
			WithUsage("Mail from name").
			WithGroup("SES").
			WithEditable(true),
	}
}

func (c *Component) ConfigWatchers() []config.Watcher {
	return []config.Watcher{
		config.NewWatcher([]string{aws.ConfigKey, aws.ConfigSecret}, c.watchCredentials),
		config.NewWatcher([]string{aws.ConfigRegion}, c.watchRegion),
		config.NewWatcher([]string{aws.ConfigLogLevel}, c.watchLogLevel),
		config.NewWatcher([]string{aws.ConfigUpdaterApplicationsDuration}, c.watchUpdaterApplicationsDuration),
		config.NewWatcher([]string{aws.ConfigUpdaterSubscriptionsDuration}, c.watchUpdaterSubscriptionsDuration),
		config.NewWatcher([]string{aws.ConfigUpdaterTopicsDuration}, c.watchUpdaterTopicsDuration),
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
