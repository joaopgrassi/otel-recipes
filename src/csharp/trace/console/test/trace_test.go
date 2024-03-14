package tracetest

import (
	"testing"

	"github.com/joaopgrassi/otel-recipes/internal/common/testutils"
	"github.com/stretchr/testify/assert"
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

const expectedSpanName = "HelloWorldSpan"

func TestTraceGeneratedFromSample(t *testing.T) {
	trace := testutils.GetTrace(t, "")

	assert.NotNil(t, trace.TraceID)
	assert.Equal(t, 1, len(trace.Spans))

	span := trace.Spans[0]
	assert.Equal(t, expectedSpanName, span.OperationName)
	assert.Contains(t, span.Tags, Tag{Key: "foo", Value: "bar"}, "Span does not contain tag 'foo:bar'")
}
