package metrics

import (
	"flag"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	common "go.opentelemetry.io/proto/otlp/common/v1"
	v1 "go.opentelemetry.io/proto/otlp/metrics/v1"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

var sample = flag.String("sample", "none", "The name of the sample app used to query traces from Jaeger")
var expDim = &common.KeyValue{Key: "foo", Value: &common.AnyValue{Value: &common.AnyValue_StringValue{StringValue: "bar"}}}

func TestMetricGeneratedFromSample(t *testing.T) {
	sm := getScopeMetricsWithRetry(t)
	metrics := sm.GetMetrics()

	// 1 counter, 1 gauge
	assert.Equal(t, len(metrics), 2)

	assertCounter(t, metrics)
	assertGauge(t, metrics)
}

func assertCounter(t *testing.T, metrics []*v1.Metric) {
	m := findMetric(t, metrics, "counter")
	s := m.GetData().(*v1.Metric_Sum)
	dp := s.Sum.DataPoints[0]
	assert.Equal(t, int64(3), dp.GetAsInt())
	assert.Contains(t, dp.Attributes, expDim, "Metric does not contain dimension 'foo:bar'")
}

func assertGauge(t *testing.T, metrics []*v1.Metric) {
	m := findMetric(t, metrics, "gauge")
	g := m.GetData().(*v1.Metric_Gauge)
	dp := g.Gauge.DataPoints[0]
	assert.Equal(t, 3.5, dp.GetAsDouble())
	assert.Contains(t, dp.Attributes, expDim, "Metric does not contain dimension 'foo:bar'")
}

func findMetric(t *testing.T, metrics []*v1.Metric, mt string) *v1.Metric {
	for _, m := range metrics {
		if strings.Contains(strings.ToLower(m.GetName()), mt) {
			return m
		}
	}
	t.Fatalf("Could not find metric with type: %s", mt)
	return nil
}

func getScopeMetricsWithRetry(t *testing.T) *v1.ScopeMetrics {
	backoffSchedule := []time.Duration{
		1 * time.Second,
		3 * time.Second,
		10 * time.Second,
	}

	var sm *v1.ScopeMetrics

	// do some retries until we Jaeger has it
	for _, backoff := range backoffSchedule {
		sm = getScopeMetrics(t)

		if sm != nil {
			break
		}

		t.Logf("Metrics not found yet, retrying in %v\n", backoff)
		time.Sleep(backoff)
	}

	// All retries failed
	if sm == nil {
		t.Fatalf("Failed getting metrics from metrics server")
	}

	return sm
}

func getScopeMetrics(t *testing.T) *v1.ScopeMetrics {
	t.Logf("Going to call the metrics server to fetch metrics for sample: %s", *sample)
	r, err := http.Get("http://localhost:4319/getMetric?scopeName=" + *sample)
	if err != nil {
		t.Fatalf("Failed getting metrics from server: %v", err)
	}

	t.Log("Received 200 response from metrics server")

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		t.Fatalf("Error reading payload from metrics server: %v", err)
	}

	sm := &v1.ScopeMetrics{}
	err = proto.Unmarshal(body, sm)
	if err != nil {
		t.Fatalf("Error reading payload from metrics server: %v", err)
	}

	return sm
}
