using System;
using System.Reflection;
using Microsoft.AspNetCore.Builder;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Logging;
using OpenTelemetry.Logs;
using OpenTelemetry.Resources;
using OpenTelemetry.Trace;

// ReSharper disable once EmptyNamespace
var appBuilder = WebApplication.CreateBuilder(args);

// Build a resource configuration action to set service information.
Action<ResourceBuilder> configureResource = r => r.AddService(
    serviceName: "aspnet.api.logs",
    serviceVersion: Assembly.GetExecutingAssembly().GetName().Version?.ToString() ?? "unknown");

// Also configure tracing so logs are correlated with traces
appBuilder.Services.AddOpenTelemetry()
    .ConfigureResource(configureResource)
    .WithTracing(builder =>
    {
        builder
            .SetSampler(new AlwaysOnSampler())
            .AddHttpClientInstrumentation()
            .AddAspNetCoreInstrumentation()
            .AddOtlpExporter(opts => {
                opts.Endpoint = new Uri("http://collector-otel-recipes:4317");
            });
    });

// Clear default logging providers used by WebApplication host.
appBuilder.Logging.ClearProviders();

// Configure OpenTelemetry Logging.
appBuilder.Logging.AddOpenTelemetry(options =>
{
    var resourceBuilder = ResourceBuilder.CreateDefault();
    configureResource(resourceBuilder);
    options.SetResourceBuilder(resourceBuilder);
    options.AddOtlpExporter(otlpOptions =>
    {
        otlpOptions.Endpoint = new Uri("http://collector-otel-recipes:4317");
    });
});

appBuilder.Services.AddControllers();
var app = appBuilder.Build();
app.UseHttpsRedirection();
app.MapControllers();
app.Run();
