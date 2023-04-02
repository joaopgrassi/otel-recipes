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
var expDim = &common.KeyValue{Key: "foo", Value: &common.AnyValue{Value: &common.AnyValue_StringValue{StringValue: "testValue"}}}

func TestMetricGeneratedFromSample(t *testing.T) {
	m := getMetricWithRetry(t)
	var dims []*common.KeyValue

	// The metric name contains the "instrument type" which we can use to extract the data
	if strings.Contains(strings.ToLower(*sample), "counter") {
		s := m.GetData().(*v1.Metric_Sum)
		dp := s.Sum.DataPoints[0]
		dims = dp.Attributes
		assert.Equal(t, 3, dp.GetAsInt())
	} else if strings.Contains(strings.ToLower(*sample), "gauge") {
		g := m.GetData().(*v1.Metric_Gauge)
		dp := g.Gauge.DataPoints[0]
		dims = dp.Attributes
		assert.Equal(t, 5.7, dp.GetAsDouble())
	} else if strings.Contains(strings.ToLower(*sample), "histogram") {
		h := m.GetData().(*v1.Metric_Histogram)
		dp := h.Histogram.DataPoints[0]
		dims = dp.Attributes
		assert.Equal(t, 3, dp.GetCount())
	}

	assert.Contains(t, dims, expDim, "Metric does not contain dimension 'foo:bar'")
}

func getMetric(t *testing.T) *v1.Metric {
	t.Logf("Going to call the metrics server to fetch metric for sample: %s", *sample)
	r, err := http.Get("http://localhost:8090/getMetric?metricName=" + *sample)
	if err != nil {
		t.Fatalf("Failed getting metric from server: %v", err)
	}

	t.Log("Received 200 response from metrics server")

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		t.Fatalf("Error reading payload from metrics server: %v", err)
	}

	m := &v1.Metric{}
	err = proto.Unmarshal(body, m)
	if err != nil {
		t.Fatalf("Error reading payload from metrics server: %v", err)
	}

	return m
}

func getMetricWithRetry(t *testing.T) *v1.Metric {
	backoffSchedule := []time.Duration{
		1 * time.Second,
		3 * time.Second,
		10 * time.Second,
	}

	var metric *v1.Metric

	// do some retries until we Jaeger has it
	for _, backoff := range backoffSchedule {
		metric = getMetric(t)

		if metric != nil {
			break
		}

		t.Logf("Metric not found yet, retrying in %v\n", backoff)
		time.Sleep(backoff)
	}

	// All retries failed
	if metric == nil {
		t.Fatalf("Failed getting metric from metrics server")
	}

	return metric
}
