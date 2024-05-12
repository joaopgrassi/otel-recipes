from opentelemetry import trace
from opentelemetry.sdk.resources import Resource
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import BatchSpanProcessor
from opentelemetry.exporter.otlp.proto.grpc.trace_exporter import OTLPSpanExporter
from opentelemetry.sdk.trace.sampling import TraceIdRatioBased

# Configure the TraceIdRatioBased sampler to sample 1 in every 1000 traces
sampler = TraceIdRatioBased(1/1000)

# Creates a resource and adds it to the tracer provider passing the sampler
resource = Resource.create({"service.name": "python.traces.traceidratio.sampler"})
provider = TracerProvider(resource=resource, sampler=sampler)

# Set the sampler onto the global tracer provider
trace.set_tracer_provider(provider)

# Adds batch span processor with the OTLP exporter to the tracer provider
provider.add_span_processor(
    BatchSpanProcessor(
        OTLPSpanExporter(endpoint="http://collector-otel-recipes:4317", insecure=True)
    )
)

tracer = trace.get_tracer(__name__)
# Generates 10k spans now sampled by the TraceIdRatioBased sampler
for _ in range(10000):
    with tracer.start_as_current_span("Sampling") as span:
        # Trace flag sampled (0x01) will be set for the sampled ones
        span.set_attribute("sampler", "traceidratiobased")
        # Check if span is active and recording (not NonRecordingSpan)
        if span.is_recording():
            print("Doing something with sampled spans")
