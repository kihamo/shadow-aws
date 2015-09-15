package service

import (
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
	h.View.Context["Applications"] = service.applications
	h.View.Context["Subscriptions"] = service.subscriptions
	h.View.Context["Topics"] = service.topics
}
