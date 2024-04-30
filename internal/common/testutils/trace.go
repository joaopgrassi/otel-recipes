package testutils // import "github.com/joaopgrassi/otel-recipes/internal/common/testutils"

import (
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	otlptrace "go.opentelemetry.io/proto/otlp/trace/v1"
	"google.golang.org/protobuf/proto"
)

func AssertSpanWithAttributeExists(t *testing.T, tc *TraceTestCase) {
	backoffSchedule := []time.Duration{
		1 * time.Second,
		3 * time.Second,
		10 * time.Second,
		15 * time.Second,
		20 * time.Second,
		30 * time.Second,
	}

	// do some retries until we backend has it
	var span *otlptrace.Span
found:
	for _, backoff := range backoffSchedule {
		rs := GetTraceWithRetry(t, tc.serviceName)
		for _, ss := range rs.ScopeSpans {
			for _, s := range ss.Spans {
				if s.Name == tc.spanName {
					span = s
					break found
				}
			}
		}
		t.Logf("Trace not found yet, retrying in %v\n", backoff)
		time.Sleep(backoff)
	}

	assert.NotNil(t, span)
	assert.Equal(t, tc.attributes[0], span.Attributes[0])
	assert.Contains(t, span.Attributes, tc.attributes[0])

	for _, exp := range tc.attributes {
		assert.Contains(t, span.Attributes, exp)
	}
}

func GetTraceWithRetry(t *testing.T, serviceName string) *otlptrace.ResourceSpans {
	backoffSchedule := []time.Duration{
		1 * time.Second,
		3 * time.Second,
		10 * time.Second,
		15 * time.Second,
		20 * time.Second,
		30 * time.Second,
	}

	var rs *otlptrace.ResourceSpans

	// do some retries until we backend has it
	for _, backoff := range backoffSchedule {
		rs = GetTrace(t, serviceName)

		if rs != nil {
			break
		}

		t.Logf("Trace not found yet, retrying in %v\n", backoff)
		time.Sleep(backoff)
	}

	return rs
}

func GetTrace(t *testing.T, serviceName string) *otlptrace.ResourceSpans {
	t.Logf("Going to call OTLP backend to fetch trace for sample: %s", serviceName)
	r, err := http.Get(fmt.Sprintf("%s/getotlp?signal=trace&servicename=%s", OtlpBackendUri, serviceName))
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

	rs := &otlptrace.ResourceSpans{}
	err = proto.Unmarshal(body, rs)
	if err != nil {
		t.Fatalf("Error reading payload from OTLP backend: %v", err)
	}
	return rs
}
