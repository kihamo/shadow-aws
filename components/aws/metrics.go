package aws

import (
	"github.com/kihamo/snitch"
)

const (
	MetricApplicationsTotal  = ComponentName + ".applications.total"
	MetricSubscriptionsTotal = ComponentName + ".subscriptions.total"
	MetricTopicsTotal        = ComponentName + ".topics.total"
	MetricEndpointsTotal     = ComponentName + ".endpoints.total"
	MetricEndpointsEnabled   = ComponentName + ".endpoints.enabled"
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
	metricApplicationsTotal = snitch.NewGauge(MetricApplicationsTotal)
	metricSubscriptionsTotal = snitch.NewGauge(MetricSubscriptionsTotal)
	metricTopicsTotal = snitch.NewGauge(MetricTopicsTotal)
	metricEndpointsTotal = snitch.NewGauge(MetricEndpointsTotal)
	metricEndpointsEnabled = snitch.NewGauge(MetricEndpointsEnabled)

	return &metricsCollector{}
}
