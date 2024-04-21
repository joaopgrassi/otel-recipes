package otel.recipes;

import io.opentelemetry.api.common.Attributes;
import io.opentelemetry.exporter.otlp.logs.OtlpGrpcLogRecordExporter;
import io.opentelemetry.instrumentation.log4j.appender.v2_17.OpenTelemetryAppender;
import io.opentelemetry.sdk.OpenTelemetrySdk;
import io.opentelemetry.sdk.logs.SdkLoggerProvider;
import io.opentelemetry.sdk.logs.export.BatchLogRecordProcessor;
import io.opentelemetry.sdk.resources.Resource;
import io.opentelemetry.semconv.ResourceAttributes;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.apache.logging.log4j.message.StringMapMessage;

public class App {

  public static void main(String[] args) throws InterruptedException {
    SdkLoggerProvider loggerProvider =
        SdkLoggerProvider.builder()
            .setResource(
                Resource.create(Attributes.of(ResourceAttributes.SERVICE_NAME, "java.console.logs")))
            .addLogRecordProcessor(
                BatchLogRecordProcessor.builder(
                        OtlpGrpcLogRecordExporter.builder()
                            .setEndpoint("http://collector-otel-recipes:4317")
                            .build())
                    .build())
            .build();

    // Sets and registers the Log Provider as Global
    OpenTelemetrySdk sdk =
        OpenTelemetrySdk.builder().setLoggerProvider(loggerProvider).buildAndRegisterGlobal();

    // Install OpenTelemetry in log4j appender
    OpenTelemetryAppender.install(sdk);

    // Creates the Logger
    Logger log4jLogger = LogManager.getLogger("mylogger");

    // Log an info message with an attribute
    StringMapMessage message = new StringMapMessage();
    message.put("foo", "bar");
    message.put("message", "This is a info message");

    log4jLogger.info(message);

    Thread.sleep(2000);
  }
}
