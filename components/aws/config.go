package aws

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
