package aws

import (
	"time"
)

type SnsApplication struct {
	Arn                       string
	AwsAttributes             map[string]*string
	Enabled                   bool
	EndpointsCount            int
	EndpointsEnabledCount     int
	CertificateExpirationDate *time.Time
	LastUpdate                time.Time
}

func (a SnsApplication) IsIAM() bool {
	if _, ok := a.AwsAttributes["SuccessFeedbackRoleArn"]; !ok {
		return false
	}

	if _, ok := a.AwsAttributes["FailureFeedbackRoleArn"]; !ok {
		return false
	}

	return true
}

func (a SnsApplication) GetEnabledCount() int {
	if a.EndpointsEnabledCount <= 0 {
		return 0
	}

	return a.EndpointsEnabledCount
}

func (a SnsApplication) GetEnabledPercent() int {
	if a.EndpointsCount <= 0 {
		return 0
	}

	return (100 * a.EndpointsEnabledCount) / a.EndpointsCount
}

func (a SnsApplication) GetDisabledCount() int {
	if a.EndpointsCount <= 0 || a.EndpointsEnabledCount <= 0 {
		return 0
	}

	return a.EndpointsCount - a.EndpointsEnabledCount
}

func (a SnsApplication) GetDisabledPercent() int {
	if a.EndpointsCount <= 0 {
		return 0
	}

	return 100 - a.GetEnabledPercent()
}
