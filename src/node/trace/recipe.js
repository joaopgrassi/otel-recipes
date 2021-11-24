// install this package @opentelemetry/sdk-node @opentelemetry/api

const opentelemetry = require("@opentelemetry/sdk-node");

const sdk = new opentelemetry.NodeSDK({
  traceExporter: new opentelemetry.tracing.ConsoleSpanExporter(),
});

sdk.start()

// how to start a span
//..
