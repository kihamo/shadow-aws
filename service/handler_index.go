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
	h.View.Context["PageTitle"] = "Aws"
	h.View.Context["PageHeader"] = "Amazon Web Service"

	service := h.Service.(*AwsService)
	h.View.Context["Services"] = service.Aws.GetServices()
	h.View.Context["Applications"] = service.GetApplications()
	h.View.Context["Subscriptions"] = service.GetSubscriptions()
	h.View.Context["Topics"] = service.GetTopics()
	h.View.Context["SDKVersion"] = aws.SDKVersion
}
