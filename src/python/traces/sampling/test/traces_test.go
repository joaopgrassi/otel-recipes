package test

import (
	"testing"

	tu "github.com/joaopgrassi/otel-recipes/internal/common/testutils"
)

func TestSamplingSpans(t *testing.T) {
	tc := tu.NewTraceTestCase("python.traces.sampling", "SamplingSpan", tu.StringAttribute("sampled", "1"), tu.StringAttribute("sampler", "traceidratio"))
	tu.AssertSpanWithAttributeExists(t, tc)
}
