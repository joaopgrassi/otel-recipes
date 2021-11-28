/**  
 * Install the following packages:
 *   - @opentelemetry/api
 *   - @opentelemetry/auto-instrumentations-node
 *   - @opentelemetry/exporter-collector-grpc
 *   - @opentelemetry/resources
 *   - @opentelemetry/semantic-conventions
 *   - @opentelemetry/sdk-trace-base
 *   - @opentelemetry/sdk-trace-node
*/

/* tracer.js */
const opentelemetry = require('@opentelemetry/api');
const { getNodeAutoInstrumentations } = require("@opentelemetry/auto-instrumentations-node");
const { CollectorTraceExporter } = require("@opentelemetry/exporter-collector-grpc");
const { registerInstrumentations } = require('@opentelemetry/instrumentation');
const { Resource } = require('@opentelemetry/resources');
const { SimpleSpanProcessor } = require('@opentelemetry/sdk-trace-base');
const { NodeTracerProvider } = require('@opentelemetry/sdk-trace-node');
const { SemanticResourceAttributes } = require('@opentelemetry/semantic-conventions');

module.exports = (serviceName) => {
  const provider = new NodeTracerProvider({
    resource: new Resource({
        [SemanticResourceAttributes.SERVICE_NAME]: serviceName,
        [SemanticResourceAttributes.SERVICE_VERSION]: "1.0.0",
    }),
  });

  const OTLPoptions = {
    url: "http://otel-collector:4317/v1/trace",
  };

  const collectorExporter = new CollectorTraceExporter(OTLPoptions);

  provider.addSpanProcessor(new SimpleSpanProcessor(collectorExporter));

  provider.register();

  registerInstrumentations({
    instrumentations: [getNodeAutoInstrumentations()],
  });

  return opentelemetry.trace.getTracer(serviceName);
};


/**  
 * How to create a manual span
*/

const api = require('@opentelemetry/api');
const tracer = require('./tracer')("nodejs-recipe");

const parentSpan = api.trace.getSpan(api.context.active());
const ctx = api.trace.setSpan(api.context.active(), parentSpan);
const span = tracer.startSpan('manual-span', undefined, ctx);

span.setAttribute('attribute_key', 'attribute_value');

span.addEvent('invoking createManualSpan');

span.end();