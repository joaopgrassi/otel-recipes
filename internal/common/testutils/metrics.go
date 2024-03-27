package testutils // import "github.com/joaopgrassi/otel-recipes/internal/common/testutils"

import (
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	otlpmetrics "go.opentelemetry.io/proto/otlp/metrics/v1"
	"google.golang.org/protobuf/proto"
)

func AssertCounter[T Number](t *testing.T, tc *MetricTestCase[T], actualMetrics []*otlpmetrics.Metric) {
	// find metric by name
	m := findMetric(t, actualMetrics, tc.metricName)

	// assert
	assert.Equal(t, tc.description, m.GetDescription())
	assert.Equal(t, tc.unit, m.GetUnit())
	s := m.GetData().(*otlpmetrics.Metric_Sum)
	dp := s.Sum.DataPoints[0]
	assert.Equal(t, tc.value, dp.GetAsInt())

	for _, exp := range tc.attributes {
		assert.Contains(t, dp.Attributes, exp)
	}
}

func AssertGauge[T Number](t *testing.T, tc *MetricTestCase[T], actualMetrics []*otlpmetrics.Metric) {
	// find metric by name
	m := findMetric(t, actualMetrics, tc.metricName)

	// assert
	assert.Equal(t, tc.description, m.GetDescription())
	assert.Equal(t, tc.unit, m.GetUnit())
	g := m.GetData().(*otlpmetrics.Metric_Gauge)
	dp := g.Gauge.DataPoints[0]
	assert.Equal(t, tc.value, dp.GetAsDouble())

	for _, exp := range tc.attributes {
		assert.Contains(t, dp.Attributes, exp)
	}
}

func findMetric(t *testing.T, metrics []*otlpmetrics.Metric, name string) *otlpmetrics.Metric {
	for _, m := range metrics {
		if m.GetName() == name {
			return m
		}
	}
	t.Fatalf("Could not find metric with name: %s", name)
	return nil
}

func GetMetricsWithRetry(t *testing.T, serviceName string) *otlpmetrics.ResourceMetrics {
	backoffSchedule := []time.Duration{
		1 * time.Second,
		3 * time.Second,
		10 * time.Second,
		15 * time.Second,
		20 * time.Second,
		30 * time.Second,
	}

	var rm *otlpmetrics.ResourceMetrics

	// do some retries until we backend has it
	for _, backoff := range backoffSchedule {
		rm = GetMetric(t, serviceName)

		if rm != nil {
			break
		}

		t.Logf("Metrics not found yet, retrying in %v\n", backoff)
		time.Sleep(backoff)
	}

	return rm
}

func GetMetric(t *testing.T, serviceName string) *otlpmetrics.ResourceMetrics {
	t.Logf("Going to call OTLP backend to fetch metrics for sample: %s", serviceName)
	r, err := http.Get(fmt.Sprintf("%s/getotlp?signal=metrics&servicename=%s", OtlpBackendUri, serviceName))
	if err != nil {
		t.Fatalf("Failed getting trace from OTLP backend: %v", err)
	}

	t.Log("Received 200 response from OTLP backend")

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		t.Fatalf("Error reading payload from OTLP backend: %v", err)
	}

	if len(body) == 0 {
		return nil
	}

	rm := &otlpmetrics.ResourceMetrics{}
	err = proto.Unmarshal(body, rm)
	if err != nil {
		t.Fatalf("Error reading metrics payload from OTLP backend: %v", err)
	}
	return rm
}
