package handlers

import (
	"net/http"

	"github.com/kihamo/shadow-aws/components/aws"
	"github.com/kihamo/shadow/components/dashboard"
)

type SNSHandler struct {
	dashboard.Handler

	Component aws.Component
}

func (h *SNSHandler) ServeHTTP(w *dashboard.Response, r *dashboard.Request) {
	if r.IsPost() {
		r.Original().ParseForm()
		updater := r.Original().PostForm.Get("updater")

		switch updater {
		case "applications":
			h.Component.RunApplicationsUpdater()
			r.Logger().Info("Run updater applications manually")
		case "subscriptions":
			h.Component.RunSubscriptionsUpdater()
			r.Logger().Info("Run updater subscriptions manually")
		case "topics":
			h.Component.RunTopicsUpdater()
			r.Logger().Info("Run updater topics manually")
		}

		h.Redirect(r.Original().RequestURI, http.StatusFound, w, r)
		return
	}

	h.Render(r.Context(), h.Component.GetName(), "sns", map[string]interface{}{
		"services":      h.Component.GetServices(),
		"applications":  h.Component.GetApplications(),
		"subscriptions": h.Component.GetSubscriptions(),
		"topics":        h.Component.GetTopics(),
	})
}
