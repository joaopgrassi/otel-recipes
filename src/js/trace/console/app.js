const api = require('@opentelemetry/api');
const { BasicTracerProvider, SimpleSpanProcessor } = require('@opentelemetry/sdk-trace-base');
const { OTLPTraceExporter } = require('@opentelemetry/exporter-trace-otlp-grpc');
const { Resource } = require('@opentelemetry/resources');
const { SemanticResourceAttributes } = require('@opentelemetry/semantic-conventions');

// Creates the tracer provider and configures OTLP collector
const provider = new BasicTracerProvider({
  resource: new Resource({
    [SemanticResourceAttributes.SERVICE_NAME]: 'js.console.app',
  }),
});

provider.addSpanProcessor(new SimpleSpanProcessor(new OTLPTraceExporter()));

provider.register();

// Creates the tracer test
const tracer = api.trace.getTracer("js.console.app");

// Start a span with an attribute
const span = tracer.startSpan("HelloWorldSpan", {
  attributes: {
    foo: 'bar'
  }
});

span.end();
