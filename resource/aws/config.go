package aws

import (
	"github.com/kihamo/shadow/resource/config"
)

const (
	ConfigAwsKey    = "aws.key"
	ConfigAwsSecret = "aws.secret"
	ConfigAwsRegion = "aws.region"
)

func (r *Resource) GetConfigVariables() []config.Variable {
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
	}
}
