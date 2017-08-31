package aws

import (
	"github.com/kihamo/snitch"
)

const (
	MetricApplicationsTotal  = ComponentName + "_applications_total"
	MetricSubscriptionsTotal = ComponentName + "_subscriptions_total"
	MetricTopicsTotal        = ComponentName + "_topics_total"
	MetricEndpointsTotal     = ComponentName + "_endpoints_total"
	MetricEndpointsEnabled   = ComponentName + "_endpoints_enabled"
	MetricSesEmailTotal      = ComponentName + "_ses_email_total"
)

var (
	metricApplicationsTotal    snitch.Gauge
	metricSubscriptionsTotal   snitch.Gauge
	metricTopicsTotal          snitch.Gauge
	metricEndpointsTotal       snitch.Gauge
	metricEndpointsEnabled     snitch.Gauge
	metricSesEmailTotalSuccess snitch.Counter
	metricSesEmailTotalFailed  snitch.Counter
)

type metricsCollector struct {
}

func (c *metricsCollector) Describe(ch chan<- *snitch.Description) {
	ch <- metricApplicationsTotal.Description()
	ch <- metricSubscriptionsTotal.Description()
	ch <- metricTopicsTotal.Description()
	ch <- metricEndpointsTotal.Description()
	ch <- metricEndpointsEnabled.Description()
	ch <- metricSesEmailTotalSuccess.Description()
	ch <- metricSesEmailTotalFailed.Description()
}

func (c *metricsCollector) Collect(ch chan<- snitch.Metric) {
	ch <- metricApplicationsTotal
	ch <- metricSubscriptionsTotal
	ch <- metricTopicsTotal
	ch <- metricEndpointsTotal
	ch <- metricEndpointsEnabled
	ch <- metricSesEmailTotalSuccess
	ch <- metricSesEmailTotalFailed
}

func (c *Component) Metrics() snitch.Collector {
	metricApplicationsTotal = snitch.NewGauge(MetricApplicationsTotal, "Number SNS applications")
	metricSubscriptionsTotal = snitch.NewGauge(MetricSubscriptionsTotal, "Number SNS subscriptions")
	metricTopicsTotal = snitch.NewGauge(MetricTopicsTotal, "Number SNS topics")
	metricEndpointsTotal = snitch.NewGauge(MetricEndpointsTotal, "Number SNS endpoints")
	metricEndpointsEnabled = snitch.NewGauge(MetricEndpointsEnabled, "Number SNS enabled endpoints")
	metricSesEmailTotalSuccess = snitch.NewCounter(MetricSesEmailTotal, "Number of SES mail with success status", "status", "success")
	metricSesEmailTotalFailed = snitch.NewCounter(MetricSesEmailTotal, "Number of SES mail with failed status", "status", "failed")

	return &metricsCollector{}
}
