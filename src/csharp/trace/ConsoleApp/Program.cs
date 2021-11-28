using System.Diagnostics;
using OpenTelemetry;
using OpenTelemetry.Resources;
using OpenTelemetry.Trace;

namespace ConsoleApp
{
    public class Program
    {
        private static readonly ActivitySource MyActivitySource = new ActivitySource(
            "ConsoleApp");
        
        public static void Main(string[] args)
        {
            using var tracerProvider = Sdk.CreateTracerProviderBuilder()
                .AddSource("ConsoleApp")
                .SetSampler(new AlwaysOnSampler())
                .SetResourceBuilder(ResourceBuilder.CreateDefault().AddService("console-app"))
                .AddOtlpExporter()
                .Build();

            using var activity = MyActivitySource.StartActivity("HelloWorldSpan");
            activity?.SetTag("foo", "bar");
        }
    }
}
