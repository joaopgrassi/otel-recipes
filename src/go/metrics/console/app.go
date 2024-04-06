package main

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

const serviceName = "go.console.app"

func main() {
	ctx := context.Background()

	// Creates the meter provider
	mp := initMeter(ctx)
	defer func() {
		if err := mp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down meter provider: %v", err)
		}
	}()

	// Creates the meter
	meter := otel.Meter(serviceName)

	attributes := attribute.NewSet(attribute.String("foo", "bar"))

	// Creates the Counter instrument
	counter, _ := meter.Int64Counter(
		"myCounter",
		metric.WithDescription("I count things"),
		metric.WithUnit("1"),
	)
	// Add to our counter with an attribute
	counter.Add(ctx, 3, metric.WithAttributeSet(attributes))

	// Creates the Gauge instrument, registering the callback that will produce the metric values
	meter.Float64ObservableGauge(
		"myGauge",
		metric.WithDescription(
			"I gauge things",
		),
		metric.WithUnit("1"),
		metric.WithFloat64Callback(func(_ context.Context, o metric.Float64Observer) error {
			o.Observe(3.5, metric.WithAttributeSet(attributes))
			return nil
		}),
	)
}

func initMeter(ctx context.Context) *sdkmetric.MeterProvider {
	// Creates a resource with the service.name attribute
	res, err := resource.New(ctx,
		resource.WithAttributes(semconv.ServiceNameKey.String(serviceName)),
	)
	handleErr(err, "failed to create the resource")

	metricsExporter, err := otlpmetricgrpc.New(
		ctx, otlpmetricgrpc.WithInsecure(), otlpmetricgrpc.WithEndpoint("collector-otel-recipes:4317"))
	handleErr(err, "failed to create the metrics exporter")

	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(res),
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(metricsExporter,
			// Default is 1m. Set to 3s for demonstrative purposes.
			sdkmetric.WithInterval(1*time.Second))),
	)
	otel.SetMeterProvider(meterProvider)
	return meterProvider
}

func handleErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}
