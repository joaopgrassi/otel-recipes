package test

import (
	"testing"

	tu "github.com/joaopgrassi/otel-recipes/internal/common/testutils"
)

func TestTraceGeneratedFromSample(t *testing.T) {
	tu.InvokeSampleApi(t, "http://localhost:8080/helloworld")

	tc := tu.NewTraceTestCase("java.springboot.api", "HelloWorldSpan", tu.StringAttribute("foo", "bar"))

	tu.AssertSpanWithAttributeExists(t, tc)
}
