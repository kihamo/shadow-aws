package aws

import (
	"strings"
)

type Arn struct {
	Arn          string
	Partition    string
	Service      string
	Region       string
	Account      string
	Resource     string
	ResourceType string
}

func ParseArn(arn string) *Arn {
	// http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-arns

	parts := strings.Split(arn, ":")
	result := Arn{
		Arn:       parts[0],
		Partition: parts[1],
		Service:   parts[2],
		Region:    parts[3],
		Account:   parts[4],
	}

	if len(parts) > 6 {
		result.Resource = parts[5]
		result.ResourceType = parts[6]
	} else {
		path := strings.Split(parts[5], "/")

		result.Resource = path[0]
		result.ResourceType = strings.Join(path[1:], "/")
	}

	return &result
}
