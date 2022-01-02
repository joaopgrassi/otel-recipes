package trace

import (
	"encoding/json"
	"flag"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type JaegerResponse struct {
	Traces []Trace `json:"data"`
}

type Trace struct {
	TraceID string `json:"traceID"`
	Spans   []struct {
		TraceID       string `json:"traceID"`
		SpanID        string `json:"spanID"`
		OperationName string `json:"operationName"`
		Tags          []Tag  `json:"tags"`
	} `json:"spans"`
}

type Tag struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

var sample = flag.String("sample", "none", "The name of the sample app used to query traces from Jaeger")

func TestTraceGeneratedFromSample(t *testing.T) {
	trace := getTrace(t)
	if trace == nil {
		t.Fatalf("error")
	}

	assert.NotNil(t, trace.TraceID)
	assert.Equal(t, 1, len(trace.Spans))

	span := trace.Spans[0]
	assert.Equal(t, "HelloWorldSpan", span.OperationName)
	assert.Contains(t, span.Tags, Tag{Key: "foo", Value: "bar"}, "Span does not contain tag 'foo:bar'")
}

func getTrace(t *testing.T) *Trace {
	t.Logf("Going to call Jaeger to fetch traces for sample: %s", *sample)
	r, err := http.Get("http://localhost:16686/api/traces?service=" + *sample)
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

	return &data.Traces[0]
}
