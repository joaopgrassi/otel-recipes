package otel.recipes;

import static io.opentelemetry.api.common.AttributeKey.stringKey;

import io.opentelemetry.api.GlobalOpenTelemetry;
import io.opentelemetry.api.common.Attributes;
import io.opentelemetry.api.metrics.LongCounter;
import io.opentelemetry.api.metrics.Meter;
import io.opentelemetry.exporter.otlp.metrics.OtlpGrpcMetricExporter;
import io.opentelemetry.sdk.OpenTelemetrySdk;
import io.opentelemetry.sdk.metrics.SdkMeterProvider;
import io.opentelemetry.sdk.metrics.export.PeriodicMetricReader;
import io.opentelemetry.sdk.resources.Resource;
import io.opentelemetry.semconv.ResourceAttributes;

public class App {

  public static void main(String[] args) {
    OtlpGrpcMetricExporter metricsExporter = OtlpGrpcMetricExporter.builder().build();

    // Builds the Meter Provider with the OTLP exporter
    SdkMeterProvider meterProvider =
        SdkMeterProvider.builder()
            .setResource(
                Resource.create(Attributes.of(ResourceAttributes.SERVICE_NAME, "java.console.app")))
            .registerMetricReader(PeriodicMetricReader.builder(metricsExporter).build())
            .build();

    // Sets and registers the Meter Provider as Global
    OpenTelemetrySdk.builder().setMeterProvider(meterProvider).buildAndRegisterGlobal();

    // Creates the meter
    Meter meter = GlobalOpenTelemetry.getMeter("java.console.app");

    // Creates the Counter instrument
    LongCounter counter =
        meter.counterBuilder("myCounter").setDescription("I count things!").setUnit("1").build();

    // Add to our counter with an attribute
    counter.add(3, Attributes.of(stringKey("foo"), "bar"));

    // Creates the Gauge instrument passing the callback that will produce the metric values
    meter
        .gaugeBuilder("myGauge")
        .setDescription("I gauge things")
        .setUnit("1")
        .buildWithCallback(gauge -> gauge.record(3.5, Attributes.of(stringKey("foo"), "bar")));

    meterProvider.shutdown();
  }
}
