package test

import (
	"testing"

	tu "github.com/joaopgrassi/otel-recipes/internal/common/testutils"
)

func TestLogGeneratedFromSample(t *testing.T) {
	tu.InvokeSampleApi(t, "http://localhost:8080/helloworld")

	tc := tu.NewLogTestCase("csharp.aspnet.api", "Information", "This is a info message {foo}", tu.StringAttribute("foo", "bar"))

	tu.AssertLogWithAttributeExists(t, tc)
}