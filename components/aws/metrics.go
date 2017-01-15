package aws

import (
	kit "github.com/go-kit/kit/metrics"
	"github.com/kihamo/shadow/components/metrics"
)

const (
	MetricApplicationsTotal  = "aws.applications.total"
	MetricSubscriptionsTotal = "aws.subscriptions.total"
	MetricTopicsTotal        = "aws.topics.total"
	MetricEndpointsTotal     = "aws.endpoints.total"
	MetricEndpointsEnabled   = "aws.endpoints.enabled"
)

var (
	metricApplicationsTotal  kit.Gauge
	metricSubscriptionsTotal kit.Gauge
	metricTopicsTotal        kit.Gauge
	metricEndpointsTotal     kit.Gauge
	metricEndpointsEnabled   kit.Gauge
)

func (c *Component) MetricsRegister(m *metrics.Component) {
	metricApplicationsTotal = m.NewGauge(MetricApplicationsTotal)
	metricSubscriptionsTotal = m.NewGauge(MetricSubscriptionsTotal)
	metricTopicsTotal = m.NewGauge(MetricTopicsTotal)
	metricEndpointsTotal = m.NewGauge(MetricEndpointsTotal)
	metricEndpointsEnabled = m.NewGauge(MetricEndpointsEnabled)
}
