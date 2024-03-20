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

const serviceName = "go.console.app"

// Creates the tracer to be shared across the application
var tracer = otel.Tracer(serviceName)

func main() {
	ctx := context.Background()

	// Configures the SDK
	// Exports to a locally running collector on port 4317
	tp := initTracer(ctx)
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	// Starts a span with an attribute
	_, span := tracer.Start(
		ctx,
		"HelloWorldSpan",
		trace.WithAttributes(attribute.String("foo", "bar")))
	defer span.End()

	fmt.Println("Hello world")
}

func initTracer(ctx context.Context) *sdktrace.TracerProvider {
	// Creates a resource with the service.name attribute
	res, err := resource.New(ctx,
		resource.WithAttributes(semconv.ServiceNameKey.String(serviceName)),
	)
	handleErr(err, "failed to create the resource")

	traceExporter, err := otlptracegrpc.New(
		ctx, otlptracegrpc.WithInsecure(), otlptracegrpc.WithEndpoint("http://collector-otel-recipes:4317"))

	handleErr(err, "failed to create the trace exporter")

	// Configures the SDK, exporting to a local running Collector
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(sdktrace.NewSimpleSpanProcessor(traceExporter)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp
}

func handleErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}
