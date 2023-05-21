using System.Collections.Generic;
using System.Diagnostics.Metrics;
using OpenTelemetry;
using OpenTelemetry.Metrics;

namespace Console
{
    public class App
    {
        // Creates the Meter
        private static readonly Meter Meter = new("csharp.console.app", "1.0");
        
        // Creates the Counter instrument
        private static readonly Counter<int> MyCounter = Meter.CreateCounter<int>("myCounter", "1", "I count things");
        
        // Creates the Gauge instrument passing the callback that will produce the metric values
        static App()
        {
            Meter.CreateObservableGauge("myGauge", GetThreadCpuTime, "1", "I gauge things");
        }

        public static void Main(string[] args)
        {
            // Configures the SDK, exporting to a local running Collector
            using var meterProvider = Sdk.CreateMeterProviderBuilder()
                .AddMeter("csharp.console.app")
                .AddOtlpExporter()
                .Build();

            // Add to our counter with an attribute
            MyCounter.Add(3, new KeyValuePair<string, object>("foo", "bar"));
        }

        internal static Measurement<double> GetThreadCpuTime()
        {
            // Simulating getting a value for ThreadCpuTime with an attribute
            return new Measurement<double>(3.5, new KeyValuePair<string, object>("foo", "bar"));
        }
    }
}
