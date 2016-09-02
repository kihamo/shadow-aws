package service

import (
	"github.com/kihamo/shadow/resource"
)

func (s *AwsService) GetConfigVariables() []resource.ConfigVariable {
	return []resource.ConfigVariable{
		resource.ConfigVariable{
			Key:   "aws.updater_applications_duration",
			Value: "10m",
			Usage: "Duration for AWS applications updater",
		},
		resource.ConfigVariable{
			Key:   "aws.updater_endpoints_bulk",
			Value: 5,
			Usage: "Bulk size for AWS endpoints updater",
		},
		resource.ConfigVariable{
			Key:   "aws.updater_endpoints_duration",
			Value: "1h",
			Usage: "Duration for AWS endpoints updater",
		},
		resource.ConfigVariable{
			Key:   "aws.updater_subscriptions_duration",
			Value: "10m",
			Usage: "Duration for AWS subscriptions updater",
		},
		resource.ConfigVariable{
			Key:   "aws.updater_topics_duration",
			Value: "10m",
			Usage: "Duration for AWS topics updater",
		},
	}
}
