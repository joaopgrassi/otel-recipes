package testutils // import "github.com/joaopgrassi/otel-recipes/internal/common/testutils"

import (
	otlpcommon "go.opentelemetry.io/proto/otlp/common/v1"
)

type TraceTestCase struct {
	serviceName string
	spanName    string
	attributes  []*otlpcommon.KeyValue
}

func NewTraceTestCase(serviceName, spanName string, attributes ...*otlpcommon.KeyValue) *TraceTestCase {
	return &TraceTestCase{
		serviceName: serviceName,
		spanName:    spanName,
		attributes:  attributes,
	}
}

type Number interface {
	int | int64 | float64
}

type MetricTestCase[T Number] struct {
	metricName  string
	description string
	unit        string
	value       T
	attributes  []*otlpcommon.KeyValue
}

func NewMetricTestCase[T Number](name, description, unit string, value T, attributes ...*otlpcommon.KeyValue) *MetricTestCase[T] {
	return &MetricTestCase[T]{
		metricName:  name,
		description: description,
		unit:        unit,
		value:       value,
		attributes:  attributes,
	}
}

type LogTestCase struct {
	serviceName string
	severity    string
	body        string
	attributes  []*otlpcommon.KeyValue
	withTrace   bool
}

func NewLogTestCase(serviceName, severity, body string, withTrace bool, attributes ...*otlpcommon.KeyValue) *LogTestCase {
	return &LogTestCase{
		serviceName: serviceName,
		severity:    severity,
		body:        body,
		withTrace:   withTrace,
		attributes:  attributes,
	}
}
