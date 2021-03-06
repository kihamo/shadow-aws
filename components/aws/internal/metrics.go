package internal

import (
	"github.com/kihamo/shadow-aws/components/aws"
	"github.com/kihamo/snitch"
)

const (
	MetricApplicationsTotal  = aws.ComponentName + "_applications_total"
	MetricSubscriptionsTotal = aws.ComponentName + "_subscriptions_total"
	MetricTopicsTotal        = aws.ComponentName + "_topics_total"
	MetricEndpointsTotal     = aws.ComponentName + "_endpoints_total"
	MetricSesEmailTotal      = aws.ComponentName + "_ses_email_total"
)

var (
	metricApplicationsTotal  = snitch.NewGauge(MetricApplicationsTotal, "Number SNS applications")
	metricSubscriptionsTotal = snitch.NewGauge(MetricSubscriptionsTotal, "Number SNS subscriptions")
	metricTopicsTotal        = snitch.NewGauge(MetricTopicsTotal, "Number SNS topics")
	metricEndpointsTotal     = snitch.NewGauge(MetricEndpointsTotal, "Number SNS endpoints")
	metricSesEmailTotal      = snitch.NewCounter(MetricSesEmailTotal, "Number of SES mail")
)

type metricsCollector struct {
}

func (c *metricsCollector) Describe(ch chan<- *snitch.Description) {
	metricApplicationsTotal.Describe(ch)
	metricSubscriptionsTotal.Describe(ch)
	metricTopicsTotal.Describe(ch)
	metricEndpointsTotal.Describe(ch)
	metricSesEmailTotal.Describe(ch)
}

func (c *metricsCollector) Collect(ch chan<- snitch.Metric) {
	metricApplicationsTotal.Collect(ch)
	metricSubscriptionsTotal.Collect(ch)
	metricTopicsTotal.Collect(ch)
	metricEndpointsTotal.Collect(ch)
	metricSesEmailTotal.Collect(ch)
}

func (c *Component) Metrics() snitch.Collector {
	return &metricsCollector{}
}
