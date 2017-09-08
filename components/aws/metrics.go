package aws

import (
	"github.com/kihamo/snitch"
)

const (
	MetricApplicationsTotal  = ComponentName + "_applications_total"
	MetricSubscriptionsTotal = ComponentName + "_subscriptions_total"
	MetricTopicsTotal        = ComponentName + "_topics_total"
	MetricEndpointsTotal     = ComponentName + "_endpoints_total"
	MetricSesEmailTotal      = ComponentName + "_ses_email_total"
)

var (
	metricApplicationsTotal  snitch.Gauge
	metricSubscriptionsTotal snitch.Gauge
	metricTopicsTotal        snitch.Gauge
	metricEndpointsTotal     snitch.Gauge
	metricSesEmailTotal      snitch.Counter
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
	metricApplicationsTotal = snitch.NewGauge(MetricApplicationsTotal, "Number SNS applications")
	metricSubscriptionsTotal = snitch.NewGauge(MetricSubscriptionsTotal, "Number SNS subscriptions")
	metricTopicsTotal = snitch.NewGauge(MetricTopicsTotal, "Number SNS topics")
	metricEndpointsTotal = snitch.NewGauge(MetricEndpointsTotal, "Number SNS endpoints")
	metricSesEmailTotal = snitch.NewCounter(MetricSesEmailTotal, "Number of SES mail")

	return &metricsCollector{}
}
