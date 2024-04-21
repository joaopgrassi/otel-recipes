using System;
using System.Diagnostics;
using System.Reflection;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.DependencyInjection;
using OpenTelemetry.Resources;
using OpenTelemetry.Trace;

namespace AspNetCoreApi;

public class Startup
{
    // Creates the tracer to be shared across the application
    // It could also be in its own class, or registered in the DI framework
    public static readonly ActivitySource Tracer = new("aspnet.api.traces");

    public void ConfigureServices(IServiceCollection services)
    {
        services.AddControllers();

        // Build a resource configuration action to set service information.
        Action<ResourceBuilder> configureResource = r => r.AddService(
            serviceName: "aspnet.api.traces",
            serviceVersion: Assembly.GetExecutingAssembly().GetName().Version?.ToString() ?? "unknown");

        // Configures the SDK, enabling Http client and ASP.NET Core instrumentation
        // Exports to a locally running collector on port 4317
        services.AddOpenTelemetry()
            .ConfigureResource(configureResource)
            .WithTracing(builder =>
        {
            builder
                .AddSource(Tracer.Name)
                .SetSampler(new AlwaysOnSampler())
                .AddHttpClientInstrumentation()
                .AddAspNetCoreInstrumentation()
                .AddOtlpExporter(opts => {
                    opts.Endpoint = new Uri("http://collector-otel-recipes:4317");
                });
        });
    }

    public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
    {
        app.UseRouting();
        app.UseEndpoints(endpoints => { endpoints.MapControllers(); });
    }
}
