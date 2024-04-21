const api = require('@opentelemetry/api');
const { BasicTracerProvider, SimpleSpanProcessor } = require('@opentelemetry/sdk-trace-base');
const { OTLPTraceExporter } = require('@opentelemetry/exporter-trace-otlp-grpc');
const { Resource } = require('@opentelemetry/resources');
const { SEMRESATTRS_SERVICE_NAME } = require('@opentelemetry/semantic-conventions');

// Creates the tracer provider and configures OTLP collector
const provider = new BasicTracerProvider({
  resource: new Resource({
    [SEMRESATTRS_SERVICE_NAME]: 'js.console.traces',
  }),
});

provider.addSpanProcessor(new SimpleSpanProcessor(new OTLPTraceExporter({
  url: "http://collector-otel-recipes:4317"
})));

provider.register();

// Creates the tracer
const tracer = api.trace.getTracer("js.console.traces");

// Start a span with an attribute
const span = tracer.startSpan("HelloWorldSpan", {
  attributes: {
    foo: 'bar'
  }
});

span.end();
