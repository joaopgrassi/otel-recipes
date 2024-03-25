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

type MetricTestCase struct {
	metricName  string
	description string
	unit        string
	value       float64
	attributes  []*otlpcommon.KeyValue
}
type MetricTestCaseOptions struct {
	MetricName  string
	Description string
	Unit        string
	Value       float64
	Attributes  []*otlpcommon.KeyValue
}

type MetricTestCaseOption func(*MetricTestCaseOptions)

func WithMetricName(name string) MetricTestCaseOption {
	return func(s *MetricTestCaseOptions) {
		s.MetricName = name
	}
}
func WithMetricDescription(description string) MetricTestCaseOption {
	return func(s *MetricTestCaseOptions) {
		s.Description = description
	}
}
func WithMetricUnit(unit string) MetricTestCaseOption {
	return func(s *MetricTestCaseOptions) {
		s.Unit = unit
	}
}
func WithMetricValue(value float64) MetricTestCaseOption {
	return func(s *MetricTestCaseOptions) {
		s.Value = value
	}
}
func WithMetricAttributes(attributes ...*otlpcommon.KeyValue) MetricTestCaseOption {
	return func(s *MetricTestCaseOptions) {
		s.Attributes = attributes
	}
}

func NewMetricTestCase(options ...MetricTestCaseOption) *MetricTestCase {
	opts := &MetricTestCaseOptions{}
	for _, option := range options {
		option(opts)
	}

	return &MetricTestCase{
		metricName:  opts.MetricName,
		description: opts.Description,
		unit:        opts.Unit,
		value:       opts.Value,
		attributes:  opts.Attributes,
	}
}
