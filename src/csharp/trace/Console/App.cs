﻿using System.Diagnostics;
using OpenTelemetry;
using OpenTelemetry.Resources;
using OpenTelemetry.Trace;

namespace Console
{
    public class App
    {
        private static readonly ActivitySource tracer = new ActivitySource("csharp.console.app");
        
        public static void Main(string[] args)
        {
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