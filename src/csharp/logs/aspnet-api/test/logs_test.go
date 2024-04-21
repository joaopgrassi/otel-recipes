package test

import (
	"testing"

	tu "github.com/joaopgrassi/otel-recipes/internal/common/testutils"
)

func TestLogGeneratedFromSample(t *testing.T) {
	tu.InvokeSampleApi(t, "http://localhost:8080/helloworld")

	tc := tu.NewLogTestCase("csharp.aspnetapi.logs", "Information", "This is a info message {foo}", true, tu.StringAttribute("foo", "bar"))

	tu.AssertLogWithAttributeExists(t, tc)
}
