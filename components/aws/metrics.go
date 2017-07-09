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
)

var (
	metricApplicationsTotal  snitch.Gauge
	metricSubscriptionsTotal snitch.Gauge
	metricTopicsTotal        snitch.Gauge
	metricEndpointsTotal     snitch.Gauge
	metricEndpointsEnabled   snitch.Gauge
)

type metricsCollector struct {
}

func (c *metricsCollector) Describe(ch chan<- *snitch.Description) {
	ch <- metricApplicationsTotal.Description()
	ch <- metricSubscriptionsTotal.Description()
	ch <- metricTopicsTotal.Description()
	ch <- metricEndpointsTotal.Description()
	ch <- metricEndpointsEnabled.Description()
}

func (c *metricsCollector) Collect(ch chan<- snitch.Metric) {
	ch <- metricApplicationsTotal
	ch <- metricSubscriptionsTotal
	ch <- metricTopicsTotal
	ch <- metricEndpointsTotal
	ch <- metricEndpointsEnabled
}

func (c *Component) Metrics() snitch.Collector {
	metricApplicationsTotal = snitch.NewGauge(MetricApplicationsTotal, "Number SNS applications")
	metricSubscriptionsTotal = snitch.NewGauge(MetricSubscriptionsTotal, "Number SNS subscriptions")
	metricTopicsTotal = snitch.NewGauge(MetricTopicsTotal, "Number SNS topics")
	metricEndpointsTotal = snitch.NewGauge(MetricEndpointsTotal, "Number SNS endpoints")
	metricEndpointsEnabled = snitch.NewGauge(MetricEndpointsEnabled, "Number SNS enabled endpoints")

	return &metricsCollector{}
}
