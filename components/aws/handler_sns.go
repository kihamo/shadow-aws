package aws

import (
	"net/http"

	"github.com/kihamo/shadow/components/dashboard"
)

type SNSHandler struct {
	dashboard.Handler

	component *Component
}

func (h *SNSHandler) ServeHTTP(w *dashboard.Response, r *dashboard.Request) {
	if r.IsPost() {
		r.Original().ParseForm()
		updater := r.Original().PostForm.Get("updater")

		switch updater {
		case "applications":
			h.component.applicationsRun <- true
			r.Logger().Info("Run updater applications manually")
		case "subscriptions":
			h.component.subscriptionsRun <- true
			r.Logger().Info("Run updater subscriptions manually")
		case "topics":
			h.component.topicsRun <- true
			r.Logger().Info("Run updater topics manually")
		}

		h.Redirect(r.Original().RequestURI, http.StatusFound, w, r)
		return
	}

	h.Render(r.Context(), ComponentName, "sns", map[string]interface{}{
		"services":      h.component.GetServices(),
		"applications":  h.component.GetApplications(),
		"subscriptions": h.component.GetSubscriptions(),
		"topics":        h.component.GetTopics(),
	})
}
