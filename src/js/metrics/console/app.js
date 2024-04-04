const {
  MeterProvider,
  PeriodicExportingMetricReader,
} = require('@opentelemetry/sdk-metrics');
const { OTLPMetricExporter } = require('@opentelemetry/exporter-metrics-otlp-grpc');
const { Resource } = require('@opentelemetry/resources');
const { SEMRESATTRS_SERVICE_NAME } = require('@opentelemetry/semantic-conventions');

const metricReader = new PeriodicExportingMetricReader({
  exporter: new OTLPMetricExporter({
    url: "http://collector-otel-recipes:4317"
  }),
});

// Creates the meter provider
const meterProvider = new MeterProvider({
  resource: new Resource({
    [SEMRESATTRS_SERVICE_NAME]: 'js.console.app',
  }),
  readers: [metricReader]
});

// Creates the meter
const meter = meterProvider.getMeter('js.console.app');

// Creates the Counter instrument
const counter = meter.createCounter('myCounter', {
  description: 'I count things',
  unit: '1'
});

// Add to our counter with an attribute
counter.add(3, { foo: 'bar' });

// Creates the Gauge instrument, registering the callback that will produce the metric values
const gauge = meter.createObservableGauge('myGauge', {
  description: 'I gauge things',
  unit: '1'
});

gauge.addCallback((result) => {
  result.observe(3.5, { foo: 'bar' });
});

meterProvider.forceFlush().then(
  () => console.log("SDK shut down successfully"),
  (err) => console.log("Error shutting down SDK", err)
);
