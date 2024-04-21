package test

import (
	"testing"

	tu "github.com/joaopgrassi/otel-recipes/internal/common/testutils"
)

func TestLogGeneratedFromSample(t *testing.T) {
	tc := tu.NewLogTestCase("java.console.logs", "INFO", "This is a info message", false, tu.StringAttribute("log4j.map_message.foo", "bar"))

	tu.AssertLogWithAttributeExists(t, tc)
}
