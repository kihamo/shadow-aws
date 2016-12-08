package aws

import (
	"github.com/kihamo/shadow/resource/config"
)

func (r *Aws) GetConfigVariables() []config.Variable {
	return []config.Variable{
		{
			Key:   "aws.key",
			Value: "",
			Usage: "AWS access key ID",
		},
		{
			Key:   "aws.secret",
			Value: "",
			Usage: "AWS secret access key",
		},
		{
			Key:   "aws.region",
			Value: "",
			Usage: "AWS region",
		},
	}
}
