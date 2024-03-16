package testutils // import "github.com/joaopgrassi/otel-recipes/internal/common/testutils"

import (
	otlpcommon "go.opentelemetry.io/proto/otlp/common/v1"
)

func StringAttribute(key, value string) *otlpcommon.KeyValue {
	return &otlpcommon.KeyValue{Key: key, Value: &otlpcommon.AnyValue{Value: &otlpcommon.AnyValue_StringValue{StringValue: value}}}
}
