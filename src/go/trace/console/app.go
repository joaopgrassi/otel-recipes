package main

import (
	"context"
	"fmt"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

func main() {
	ctx := context.Background()
	res, err := resource.New(ctx,
		resource.WithAttributes(semconv.ServiceNameKey.String("go.console.app")),
	)
	handleErr(err, "failed to create the resource")

	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure())
	handleErr(err, "failed to create the trace exporter")

	// Configures the SDK, exporting to a local running Collector
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(sdktrace.NewSimpleSpanProcessor(traceExporter)),
	)
	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.TraceContext{})
	defer tracerProvider.Shutdown(ctx)

	// Creates the tracer test
	tracer := otel.Tracer("go.console.app")

	// Start a span with an attribute
	ctx, span := tracer.Start(
		ctx,
		"HelloWorldSpan",
		trace.WithAttributes(attribute.String("foo", "bar")))
	defer span.End()

	// Important: pass around the ctx to operations!
	fmt.Println("Hello world")
}

func handleErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}
