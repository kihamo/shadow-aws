package aws

import (
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/kihamo/shadow"
)

type Component interface {
	shadow.Component

	GetSNS() *sns.SNS
	GetServices() map[string]interface{}
	GetApplications() []SnsApplication
	GetSubscriptions() []*sns.Subscription
	GetTopics() []*sns.Topic
	RunApplicationsUpdater()
	RunSubscriptionsUpdater()
	RunTopicsUpdater()

	GetSES() *ses.SES
	SendEmail(to []string, subject string, text string, html string, from string) error
}
