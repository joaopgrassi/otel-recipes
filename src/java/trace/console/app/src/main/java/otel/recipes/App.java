package otel.recipes;

import io.opentelemetry.api.GlobalOpenTelemetry;
import io.opentelemetry.api.common.Attributes;
import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.Tracer;
import io.opentelemetry.exporter.otlp.trace.OtlpGrpcSpanExporter;
import io.opentelemetry.sdk.OpenTelemetrySdk;
import io.opentelemetry.sdk.resources.Resource;
import io.opentelemetry.sdk.trace.SdkTracerProvider;
import io.opentelemetry.sdk.trace.export.SimpleSpanProcessor;
import io.opentelemetry.semconv.resource.attributes.ResourceAttributes;

public class App {
  public static void main(String[] args) throws InterruptedException {
    OtlpGrpcSpanExporter spanExporter = OtlpGrpcSpanExporter.builder().build();

    // Builds the Tracer Provider with span processor and resource attribute
    SdkTracerProvider tracerProvider =
        SdkTracerProvider.builder()
            .addSpanProcessor(SimpleSpanProcessor.create(spanExporter))
            .setResource(
                Resource.create(Attributes.of(ResourceAttributes.SERVICE_NAME, "java.console.app")))
            .build();

    // Sets and registers the Tracer Provider as Global
    OpenTelemetrySdk.builder().setTracerProvider(tracerProvider).buildAndRegisterGlobal();

    // Creates the tracer
    Tracer tracer = GlobalOpenTelemetry.getTracer("java.console.app");

    // Creates a span, set its attributes and finishes it
    Span span = tracer.spanBuilder("HelloWorldSpan").startSpan();
    span.setAttribute("foo", "bar");
    span.end();

    // Waits for things to settle and shuts the Trace Provider down right before finishing the app
    Thread.sleep(2000);
    tracerProvider.shutdown();
  }
}
