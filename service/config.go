package service

import (
	"github.com/kihamo/shadow/resource/config"
)

func (s *AwsService) GetConfigVariables() []config.Variable {
	return []config.Variable{
		{
			Key:   "aws.run_updater_on_startup",
			Value: true,
			Usage: "Run updater jobs on startup",
		},
		{
			Key:   "aws.updater_applications_duration",
			Value: "10m",
			Usage: "Duration for AWS applications updater",
		},
		{
			Key:   "aws.updater_endpoints_bulk",
			Value: 5,
			Usage: "Bulk size for AWS endpoints updater",
		},
		{
			Key:   "aws.updater_endpoints_duration",
			Value: "1h",
			Usage: "Duration for AWS endpoints updater",
		},
		{
			Key:   "aws.updater_subscriptions_duration",
			Value: "10m",
			Usage: "Duration for AWS subscriptions updater",
		},
		{
			Key:   "aws.updater_topics_duration",
			Value: "10m",
			Usage: "Duration for AWS topics updater",
		},
	}
}
