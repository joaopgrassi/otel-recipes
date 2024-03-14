package testutils // import "github.com/joaopgrassi/otel-recipes/internal/common/testutils"

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"
)

type JaegerResponse struct {
	Traces []Trace `json:"data"`
}

type Trace struct {
	TraceID string `json:"traceID"`
	Spans   []Span `json:"spans"`
}

type Span struct {
	TraceID       string `json:"traceID"`
	SpanID        string `json:"spanID"`
	OperationName string `json:"operationName"`
	Tags          []Tag  `json:"tags"`
}

type Tag struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

func GetTrace(t *testing.T, serviceName string) *Trace {
	t.Logf("Going to call Jaeger to fetch trace for sample: %s", serviceName)
	r, err := http.Get("http://localhost:16686/api/traces?service=" + serviceName)
	if err != nil {
		t.Fatalf("Failed getting trace from Jaeger: %v", err)
	}

	t.Log("Received 200 response from Jaeger")

	defer r.Body.Close()
	var data JaegerResponse

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		t.Fatalf("Failed decoding json response from Jaeger: %v", err)
	}

	// useful for CI runs
	json, _ := json.MarshalIndent(data, "", "  ")
	t.Logf("Data received from Jaeger: \n%s\n", json)

	if len(data.Traces) == 0 {
		return nil
	}

	return &data.Traces[0]
}

func GetTraceWithRetry(t *testing.T, serviceName string) *Trace {
	backoffSchedule := []time.Duration{
		1 * time.Second,
		3 * time.Second,
		10 * time.Second,
	}

	var trace *Trace

	// do some retries until we Jaeger has it
	for _, backoff := range backoffSchedule {
		trace = GetTrace(t, serviceName)

		if trace != nil {
			break
		}

		t.Logf("Trace not found yet, retrying in %v\n", backoff)
		time.Sleep(backoff)
	}

	// All retries failed
	if trace == nil {
		t.Fatalf("Failed getting trace from Jaeger")
	}

	return trace
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
