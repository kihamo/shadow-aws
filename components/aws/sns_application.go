package aws

import (
	"time"
)

type AwsSnsApplication struct {
	Arn                       string
	AwsAttributes             map[string]*string
	Enabled                   bool
	EndpointsCount            int
	EndpointsEnabledCount     int
	CertificateExpirationDate *time.Time
	LastUpdate                time.Time
}

func (a AwsSnsApplication) IsIAM() bool {
	if _, ok := a.AwsAttributes["SuccessFeedbackRoleArn"]; !ok {
		return false
	}

	if _, ok := a.AwsAttributes["FailureFeedbackRoleArn"]; !ok {
		return false
	}

	return true
}

func (a AwsSnsApplication) GetEnabledCount() int {
	if a.EndpointsEnabledCount <= 0 {
		return 0
	}

	return a.EndpointsEnabledCount
}

func (a AwsSnsApplication) GetEnabledPercent() int {
	if a.EndpointsCount <= 0 {
		return 0
	}

	return (100 * a.EndpointsEnabledCount) / a.EndpointsCount
}

func (a AwsSnsApplication) GetDisabledCount() int {
	if a.EndpointsCount <= 0 || a.EndpointsEnabledCount <= 0 {
		return 0
	}

	return a.EndpointsCount - a.EndpointsEnabledCount
}

func (a AwsSnsApplication) GetDisabledPercent() int {
	if a.EndpointsCount <= 0 {
		return 0
	}

	return 100 - a.GetEnabledPercent()
}
