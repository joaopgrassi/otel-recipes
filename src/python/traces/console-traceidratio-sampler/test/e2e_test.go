package test

import (
	"testing"

	tu "github.com/joaopgrassi/otel-recipes/internal/common/testutils"
	"github.com/stretchr/testify/assert"
	otlptrace "go.opentelemetry.io/proto/otlp/trace/v1"
)

const (
	serviceName string = "python.traces.traceidratio.sampler"
	spanName    string = "Sampling"
)

var samplerAttribute = tu.StringAttribute("sampler", "traceidratiobased")

func GetSpansByName(t *testing.T, serviceName string, spanName string) []*otlptrace.Span {
	var sl []*otlptrace.Span
	ss := tu.GetTraceWithRetry(t, serviceName).GetScopeSpans()
	for _, scope := range ss {
		spans := scope.GetSpans()
		for _, s := range spans {
			if s.Name == spanName {
				sl = append(sl, s)
			}
		}
	}
	return sl
}

func TestSampledSpansCount(t *testing.T) {
	spans := GetSpansByName(t, serviceName, spanName)
	assert.Equal(t, 1, len(spans))
}

func TestSampledSpansAttributes(t *testing.T) {
	spans := GetSpansByName(t, serviceName, spanName)
	for _, s := range spans {
		assert.NotNil(t, s)
		assert.Equal(t, samplerAttribute, s.Attributes[0])
		assert.Contains(t, s.Attributes, samplerAttribute)
	}
}
