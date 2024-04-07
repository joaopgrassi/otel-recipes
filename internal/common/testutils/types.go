package testutils // import "github.com/joaopgrassi/otel-recipes/internal/common/testutils"

import (
	otlpcommon "go.opentelemetry.io/proto/otlp/common/v1"
)

type SpanTest struct {
	serviceName string
	spanName    string
	attributes  []*otlpcommon.KeyValue
}

type SpanTestOptions struct {
	ServiceName string
	SpanName    string
	Attributes  []*otlpcommon.KeyValue
}

type SpanTestOption func(*SpanTestOptions)

func WithServiceName(name string) SpanTestOption {
	return func(s *SpanTestOptions) {
		s.ServiceName = name
	}
}

func WithSpanName(name string) SpanTestOption {
	return func(s *SpanTestOptions) {
		s.SpanName = name
	}
}

func WithAttributes(attributes ...*otlpcommon.KeyValue) SpanTestOption {
	return func(s *SpanTestOptions) {
		s.Attributes = attributes
	}
}

func NewSpanTest(options ...SpanTestOption) *SpanTest {
	opts := &SpanTestOptions{}
	for _, option := range options {
		option(opts)
	}

	return &SpanTest{
		serviceName: opts.ServiceName,
		spanName:    opts.SpanName,
		attributes:  opts.Attributes,
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
}

func NewLogTestCase(serviceName, severity, body string, attributes ...*otlpcommon.KeyValue) *LogTestCase {
	return &LogTestCase{
		serviceName: serviceName,
		severity:    severity,
		body:        body,
		attributes:  attributes,
	}
}
