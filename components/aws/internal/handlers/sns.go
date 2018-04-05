package handlers

import (
	"net/http"

	"github.com/kihamo/shadow-aws/components/aws"
	"github.com/kihamo/shadow/components/dashboard"
)

type SNSHandler struct {
	dashboard.Handler
}

func (h *SNSHandler) ServeHTTP(w *dashboard.Response, r *dashboard.Request) {
	component := r.Component().(aws.Component)

	if r.IsPost() {
		r.Original().ParseForm()
		updater := r.Original().PostForm.Get("updater")

		switch updater {
		case "applications":
			component.RunApplicationsUpdater()
			r.Logger().Info("Run updater applications manually")
		case "subscriptions":
			component.RunSubscriptionsUpdater()
			r.Logger().Info("Run updater subscriptions manually")
		case "topics":
			component.RunTopicsUpdater()
			r.Logger().Info("Run updater topics manually")
		}

		h.Redirect(r.Original().RequestURI, http.StatusFound, w, r)
		return
	}

	h.Render(r.Context(), "sns", map[string]interface{}{
		"services":      component.GetServices(),
		"applications":  component.GetApplications(),
		"subscriptions": component.GetSubscriptions(),
		"topics":        component.GetTopics(),
	})
}
