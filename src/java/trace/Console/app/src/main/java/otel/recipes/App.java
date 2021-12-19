package otel.recipes;

import io.opentelemetry.api.OpenTelemetry;
import io.opentelemetry.api.common.Attributes;
import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.Tracer;
import io.opentelemetry.sdk.OpenTelemetrySdk;
import io.opentelemetry.sdk.resources.Resource;

import io.opentelemetry.exporter.otlp.trace.OtlpGrpcSpanExporter;
import io.opentelemetry.sdk.trace.SdkTracerProvider;
import io.opentelemetry.sdk.trace.export.SimpleSpanProcessor;
import io.opentelemetry.semconv.resource.attributes.ResourceAttributes;

public class App {
    public static void main(String[] args) {
        OtlpGrpcSpanExporter spanExporter = OtlpGrpcSpanExporter.builder().build();

        OpenTelemetry openTelemetry = OpenTelemetrySdk.builder()
            .setTracerProvider(
                SdkTracerProvider.builder()
                    .addSpanProcessor(SimpleSpanProcessor.create(spanExporter))
                    .setResource(
                        Resource.create(
                            Attributes.of(ResourceAttributes.SERVICE_NAME, "java.console.app")))
                    .build())
            .build();

        Tracer tracer = openTelemetry.getTracer("java.console.app");

        Span span = tracer.spanBuilder("HelloWorldSpan").startSpan();
        span.setAttribute("foo", "bar");
        span.end();
        System.out.println("Hello World");
    }
}