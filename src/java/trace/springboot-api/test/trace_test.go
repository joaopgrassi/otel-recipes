package test

import (
	"testing"

	tu "github.com/joaopgrassi/otel-recipes/internal/common/testutils"
)

func TestTraceGeneratedFromSample(t *testing.T) {
	tu.InvokeSampleApi(t, "http://localhost:8080/helloworld")

	tc := tu.NewSpanTest(
		tu.WithServiceName("java.springboot.api"),
		tu.WithSpanName("HelloWorldSpan"),
		tu.WithAttributes(tu.StringAttribute("foo", "bar")),
	)

	tu.AssertSpanWithAttributeExists(t, tc)
}
