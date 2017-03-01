package aws

import (
	"net/http"

	sdk "github.com/aws/aws-sdk-go/aws"
	"github.com/kihamo/shadow/components/dashboard"
)

type IndexHandler struct {
	dashboard.Handler

	component *Component
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.IsPost(r) {
		r.ParseForm()
		updater := r.PostForm.Get("updater")

		switch updater {
		case "applications":
			h.component.applicationsRun <- true
			dashboard.LoggerFromContext(r.Context()).Info("Run updater applications manually")
		case "subscriptions":
			h.component.subscriptionsRun <- true
			dashboard.LoggerFromContext(r.Context()).Info("Run updater subscriptions manually")
		case "topics":
			h.component.topicsRun <- true
			dashboard.LoggerFromContext(r.Context()).Info("Run updater topics manually")
		}

		h.Redirect(r.RequestURI, http.StatusFound, w, r)
		return
	}

	h.Render(r.Context(), ComponentName, "index", map[string]interface{}{
		"services":      h.component.GetServices(),
		"applications":  h.component.GetApplications(),
		"subscriptions": h.component.GetSubscriptions(),
		"topics":        h.component.GetTopics(),
		"sdkVersion":    sdk.SDKVersion,
	})
}
