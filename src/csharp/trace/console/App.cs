using System;
using System.Diagnostics;
using OpenTelemetry;
using OpenTelemetry.Resources;
using OpenTelemetry.Trace;

namespace Console
{
    public class App
    {
        // Creates the tracer
        private static readonly ActivitySource tracer = new ActivitySource("csharp.console.app");

        public static void Main(string[] args)
        {

            string endpoint = Environment.GetEnvironmentVariable("OTEL_EXPORTER_OTLP_ENDPOINT");
            System.Console.WriteLine(endpoint);
            System.Console.WriteLine("test console log");

            // Configures the SDK, exporting to a local running Collector
            using var tracerProvider = Sdk.CreateTracerProviderBuilder()
                // Need to register the tracer. Wildcard is supported: e.g. MyApp.*
                .AddSource(tracer.Name)
                .SetSampler(new AlwaysOnSampler())
                .SetResourceBuilder(ResourceBuilder.CreateDefault().AddService("csharp.console.app"))
                .AddOtlpExporter()
                .Build();

            // Start a span with a tag
            using var activity = tracer.StartActivity("HelloWorldSpan");
            activity?.SetTag("foo", "bar");
        }
    }
}
