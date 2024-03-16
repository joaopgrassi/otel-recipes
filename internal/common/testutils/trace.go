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

func AssertSpanWithAttributeExists(t *testing.T, spanTest *SpanTest) {
	rs := GetTraceWithRetry(t, spanTest.serviceName)

	var span *otlptrace.Span

	for _, ss := range rs.ScopeSpans {
		for _, s := range ss.Spans {
			if s.Name == spanTest.spanName {
				span = s
			}
		}
	}

	assert.NotNil(t, span)
	assert.Equal(t, spanTest.attributes[0], span.Attributes[0])
	assert.Contains(t, span.Attributes, spanTest.attributes[0])

	for _, exp := range spanTest.attributes {
		assert.Contains(t, span.Attributes, exp)
	}
}

func GetTraceWithRetry(t *testing.T, serviceName string) *otlptrace.ResourceSpans {
	backoffSchedule := []time.Duration{
		1 * time.Second,
		3 * time.Second,
		10 * time.Second,
	}

	var rs *otlptrace.ResourceSpans

	// do some retries until we Jaeger has it
	for _, backoff := range backoffSchedule {
		rs = GetTrace(t, serviceName)

		if rs != nil {
			break
		}

		t.Logf("Trace not found yet, retrying in %v\n", backoff)
		time.Sleep(backoff)
	}

	// All retries failed
	if rs == nil {
		t.Fatalf("Failed getting trace from OTLP backend")
	}

	return rs
}

func GetTrace(t *testing.T, serviceName string) *otlptrace.ResourceSpans {
	t.Logf("Going to call OTLP backend to fetch trace for sample: %s", serviceName)
	r, err := http.Get(fmt.Sprintf("%s/getotlp?signal=trace&serviceName=%s", OtlpBackendUri, serviceName))
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

	rs := &otlptrace.ResourceSpans{}
	err = proto.Unmarshal(body, rs)
	if err != nil {
		t.Fatalf("Error reading payload from OTLP backend: %v", err)
	}
	return rs
}

func InvokeSampleApi(t *testing.T, url string) string {
	t.Logf("Going to call the sample API to generate trace for URL: %s", url)
	r, err := http.Get(url)
	if err != nil {
		t.Fatalf("Failed calling the sample endpoint in the sample API: %v", err)
	}

	t.Log("Received 200 response from the sample API")

	defer r.Body.Close()

	//We Read the response body on the line below.
	body, err := io.ReadAll(io.Reader(r.Body))
	if err != nil {
		t.Fatalf("Failed reading response body from the sample API: %v", err)
	}

	return string(body)
}
