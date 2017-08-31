package aws

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/kihamo/shadow/components/dashboard"
)

type SESHandler struct {
	dashboard.Handler

	component *Component
}

func (h *SESHandler) ServeHTTP(_ *dashboard.Response, r *dashboard.Request) {
	errors := []string{}
	vars := map[string]interface{}{
		"sent":          float64(0),
		"maxSendRate":   float64(0),
		"max24HourSend": float64(0),
		"remaining":     float64(0),
		"sentPercent":   float64(0),
		"stats":         nil,
		"statsStart":    time.Now().Add(-time.Hour * 4),
		"errors":        nil,
		"message":       nil,

		"sendFrom":    r.Config().GetString(ConfigSesFromEmail),
		"sendTo":      "",
		"sendSubject": "",
		"sendMessage": "",
		"sendType":    "html",
	}

	name := r.Config().GetString(ConfigSesFromName)
	if name != "" {
		vars["sendFrom"] = fmt.Sprintf("\"%s\" <%s>", name, vars["sendFrom"])
	}

	if r.IsPost() {
		var (
			text, html string
		)

		vars["sendFrom"] = r.Original().FormValue("from")
		vars["sendTo"] = r.Original().FormValue("to")
		vars["sendSubject"] = r.Original().FormValue("subject")
		vars["sendMessage"] = r.Original().FormValue("message")
		vars["sendType"] = r.Original().FormValue("type")

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
