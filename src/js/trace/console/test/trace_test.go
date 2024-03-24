package test

import (
	"testing"

	tu "github.com/joaopgrassi/otel-recipes/internal/common/testutils"
)

func TestTraceGeneratedFromSample(t *testing.T) {
	tc := tu.NewSpanTest(
		tu.WithServiceName("js.console.app"),
		tu.WithSpanName("HelloWorldSpan"),
		tu.WithAttributes(tu.StringAttribute("foo", "bar")),
	)

	tu.AssertSpanWithAttributeExists(t, tc)
}
