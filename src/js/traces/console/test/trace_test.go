package test

import (
	"testing"

	tu "github.com/joaopgrassi/otel-recipes/internal/common/testutils"
)

func TestTraceGeneratedFromSample(t *testing.T) {
	tc := tu.NewTraceTestCase("js.console.traces", "HelloWorldSpan", tu.StringAttribute("foo", "bar"))

	tu.AssertSpanWithAttributeExists(t, tc)
}
