package service

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/kihamo/shadow/service/frontend"
)

type IndexHandler struct {
	frontend.AbstractFrontendHandler
}

func (h *IndexHandler) Handle() {
	h.SetTemplate("index.tpl.html")
	h.SetPageTitle("Aws")
	h.SetPageHeader("Amazon Web Service")

	service := h.Service.(*AwsService)

	h.SetVar("Services", service.aws.GetServices())
	h.SetVar("Applications", service.GetApplications())
	h.SetVar("Subscriptions", service.GetSubscriptions())
	h.SetVar("Topics", service.GetTopics())
	h.SetVar("SDKVersion", aws.SDKVersion)
}
