"use strict";

const opentelemetry = require('@opentelemetry/api');
const { getNodeAutoInstrumentations } = require("@opentelemetry/auto-instrumentations-node");
const { CollectorTraceExporter } = require("@opentelemetry/exporter-collector-grpc");
const { registerInstrumentations } = require('@opentelemetry/instrumentation');
const { Resource } = require('@opentelemetry/resources');
const { SemanticResourceAttributes } = require('@opentelemetry/semantic-conventions');
const { SimpleSpanProcessor } = require('@opentelemetry/sdk-trace-base');
const { NodeTracerProvider } = require('@opentelemetry/sdk-trace-node');

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