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
trace.set_tracer_provider(TracerProvider(resource=resource))

# Adds span processor with the OTLP exporter to the tracer provider
trace.get_tracer_provider().add_span_processor(
    BatchSpanProcessor(OTLPSpanExporter())
)
tracer = trace.get_tracer(__name__)

# Starts and sets an attribute to a span test
with tracer.start_as_current_span("HelloWorldSpan") as span:
    span.set_attribute("foo",  "bar")
    print("Hello world")
