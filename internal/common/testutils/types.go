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

	st := &SpanTest{
		serviceName: opts.ServiceName,
		spanName:    opts.SpanName,
		attributes:  opts.Attributes,
	}

	return st
}
