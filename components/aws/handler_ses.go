package aws

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/kihamo/shadow/components/dashboard"
)

type SESHandler struct {
	dashboard.Handler

	component *Component
}

func (h *SESHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	config := dashboard.ConfigFromContext(r.Context())

	errors := []string{}
	vars := map[string]interface{}{
		"sent":          0,
		"maxSendRate":   0,
		"max24HourSend": 0,
		"remaining":     0,
		"sentPercent":   0,
		"stats":         nil,
		"errors":        nil,
		"message":       nil,

		"sendFrom":    config.GetString(ConfigAwsSesFromEmail),
		"sendTo":      "",
		"sendSubject": "",
		"sendMessage": "",
		"sendType":    "html",
	}

	name := config.GetString(ConfigAwsSesFromName)
	if name != "" {
		vars["sendFrom"] = fmt.Sprintf("\"%s\" <%s>", name, vars["sendFrom"])
	}

	if h.IsPost(r) {
		var (
			text, html string
		)

		vars["sendFrom"] = r.FormValue("from")
		vars["sendTo"] = r.FormValue("to")
		vars["sendSubject"] = r.FormValue("subject")
		vars["sendMessage"] = r.FormValue("message")
		vars["sendType"] = r.FormValue("type")

		if vars["sendType"] == "html" {
			html = vars["sendMessage"].(string)
		} else {
			text = vars["sendMessage"].(string)
		}

		to := strings.Split(vars["sendTo"].(string), ",")

		if err := h.component.SendEmail(to, vars["sendSubject"].(string), text, html, vars["sendFrom"].(string)); err != nil {
			errors = append(errors, err.Error())
		} else {
			vars["message"] = "Message send success"
		}
	}

	service := h.component.GetSES()

	// quota
	quota, err := service.GetSendQuota(&ses.GetSendQuotaInput{})
	if err == nil {
		vars["sent"] = quota.SentLast24Hours
		vars["maxSendRate"] = quota.MaxSendRate
		vars["max24HourSend"] = quota.Max24HourSend
		vars["remaining"] = *quota.Max24HourSend - *quota.SentLast24Hours
		vars["sentPercent"] = (*quota.SentLast24Hours * 100.0) / *quota.Max24HourSend
	} else {
		errors = append(errors, err.Error())
	}

	// stats
	stats, err := service.GetSendStatistics(&ses.GetSendStatisticsInput{})
	if err == nil {
		vars["stats"] = stats.SendDataPoints
	} else {
		errors = append(errors, err.Error())
	}

	vars["errors"] = errors

	h.Render(r.Context(), ComponentName, "ses", vars)
}
