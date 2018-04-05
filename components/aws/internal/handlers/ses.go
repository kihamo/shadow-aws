package handlers

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/kihamo/shadow-aws/components/aws"
	"github.com/kihamo/shadow/components/dashboard"
	"github.com/kihamo/shadow/components/i18n"
)

type SESHandler struct {
	dashboard.Handler
}

func (h *SESHandler) ServeHTTP(_ *dashboard.Response, r *dashboard.Request) {
	errors := make([]string, 0, 0)
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

		"sendFrom":    r.Config().String(aws.ConfigSesFromEmail),
		"sendTo":      "",
		"sendSubject": "",
		"sendMessage": "",
		"sendType":    "html",
	}

	name := r.Config().String(aws.ConfigSesFromName)
	if name != "" {
		vars["sendFrom"] = fmt.Sprintf("\"%s\" <%s>", name, vars["sendFrom"])
	}

	component := r.Component().(aws.Component)

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

		if err := component.SendEmail(to, vars["sendSubject"].(string), text, html, vars["sendFrom"].(string)); err != nil {
			errors = append(errors, err.Error())
		} else {
			vars["message"] = i18n.NewOrNopFromRequest(r).Translate(r.Component().Name(), "Message send success", "")
		}
	}

	service := component.GetSES()

	// quota
	quota, err := service.GetSendQuota(&ses.GetSendQuotaInput{})
	if err == nil {
		vars["sent"] = *quota.SentLast24Hours
		vars["maxSendRate"] = *quota.MaxSendRate
		vars["max24HourSend"] = *quota.Max24HourSend
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

	h.Render(r.Context(), "ses", vars)
}
