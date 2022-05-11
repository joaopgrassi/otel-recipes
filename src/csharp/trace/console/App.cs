using System.Diagnostics;
using OpenTelemetry;
using OpenTelemetry.Resources;
using OpenTelemetry.Trace;

namespace Console
{
    public class App
    {
        // Creates the Tracer
        private static readonly ActivitySource Tracer = new ActivitySource("csharp.console.app");

        public static void Main(string[] args)
        {
            // test
            // Configures the SDK, exporting to a local running Collector
            using var tracerProvider = Sdk.CreateTracerProviderBuilder()
                .AddSource(Tracer.Name)
                .SetSampler(new AlwaysOnSampler())
                .SetResourceBuilder(ResourceBuilder.CreateDefault().AddService("csharp.console.app"))
                .AddOtlpExporter()
                .Build();

            // Start a span with a tag
            using var activity = Tracer.StartActivity("HelloWorldSpan");
            activity?.SetTag("foo", "bar");
        }
    }
}
