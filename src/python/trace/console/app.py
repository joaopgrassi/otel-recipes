from opentelemetry import trace
from opentelemetry.sdk.resources import Resource
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import (
    BatchSpanProcessor
)
from opentelemetry.exporter.otlp.proto.grpc.trace_exporter import (
    OTLPSpanExporter
)

# Creates a resource and adds it to the tracer provider
resource = Resource.create({"service.name": "python.console.app"})
provider = TracerProvider(resource=resource)
trace.set_tracer_provider(provider)

# Adds span processor with the OTLP exporter to the tracer provider
provider.add_span_processor(
    BatchSpanProcessor(OTLPSpanExporter())
)
tracer = trace.get_tracer(__name__)

# Starts and sets an attribute to a span
with tracer.start_as_current_span("HelloWorldSpan") as span:
    span.set_attribute("foo",  "bar")
    print("Hello world")
