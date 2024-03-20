package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

const serviceName = "go.gin.api"

// Tracer the tracer to be shared across the application
var Tracer trace.Tracer

func main() {
	ctx := context.Background()

	// Configures the SDK
	// Exports to a locally running collector on port 4317
	tp := initTracer()
	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	r := gin.New()
	// Enable the Gin auto-instrumentation
	r.Use(otelgin.Middleware(serviceName))
	r.GET("/helloworld", GetHelloWorld)
	_ = r.Run(":8080")
}

func initTracer() *sdktrace.TracerProvider {
	ctx := context.Background()

	// Creates a resource with the service.name attribute
	res, err := resource.New(ctx,
		resource.WithAttributes(semconv.ServiceNameKey.String(serviceName)),
	)
	handleErr(err, "failed to create the resource")

	// Exports to a locally running collector on port 4317
	traceExporter, err := otlptracegrpc.New(
		ctx, otlptracegrpc.WithInsecure(), otlptracegrpc.WithEndpoint("collector-otel-recipes:4317"))

	handleErr(err, "failed to create the trace exporter")

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(sdktrace.NewSimpleSpanProcessor(traceExporter)),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	// Initializes the tracer to be used across the application
	Tracer = otel.Tracer(serviceName)
	return tp
}

func handleErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}
