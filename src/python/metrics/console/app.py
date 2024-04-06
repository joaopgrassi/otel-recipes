from typing import Iterable
from opentelemetry import metrics
from opentelemetry.metrics import Observation, CallbackOptions
from opentelemetry.sdk.resources import Resource
from opentelemetry.sdk.metrics import MeterProvider
from opentelemetry.sdk.metrics.export import (
    PeriodicExportingMetricReader,
)
from opentelemetry.exporter.otlp.proto.grpc.metric_exporter import OTLPMetricExporter

# Creates a resource and adds it to the tracer provider
resource = Resource.create({"service.name": "python.console.app"})

reader = PeriodicExportingMetricReader(OTLPMetricExporter(endpoint="http://collector-otel-recipes:4317"))
provider = MeterProvider(metric_readers=[reader], resource=resource)
metrics.set_meter_provider(provider)

# Creates the meter
meter = provider.get_meter("python.console.app")

# Creates the Counter instrument
counter = meter.create_counter("mycounter", "1", "I count things")

# Add to our counter with an attribute
counter.add(3, {"foo": "bar"})


# Creates the Gauge instrument, registering the callback that will produce the metric values
def observable_gauge_func(options: CallbackOptions) -> Iterable[Observation]:
    yield Observation(3.5, {"foo": "bar"})


meter.create_observable_gauge(
    "mygauge",
    callbacks=[observable_gauge_func],
    unit="1",
    description="I gauge things",
)
