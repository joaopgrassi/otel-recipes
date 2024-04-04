package test

import (
	"testing"

	tu "github.com/joaopgrassi/otel-recipes/internal/common/testutils"
)

func TestMetricsGeneratedFromSample(t *testing.T) {
	rm := tu.GetMetricsWithRetry(t, "js.console.app")
	m := rm.GetScopeMetrics()[0].Metrics

	// Counter metric
	ctc := tu.NewMetricTestCase("myCounter", "I count things", "1", int64(3), tu.StringAttribute("foo", "bar"))
	tu.AssertCounter(t, ctc, m)

	// Gauge metric
	ctg := tu.NewMetricTestCase("myGauge", "I gauge things", "1", float64(3.5), tu.StringAttribute("foo", "bar"))
	tu.AssertGauge(t, ctg, m)
}
