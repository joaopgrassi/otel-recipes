using System.Diagnostics;
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
    public static readonly ActivitySource Tracer = new("csharp.aspnet.api");

    public void ConfigureServices(IServiceCollection services)
    {
        services.AddControllers();

        // Configures the SDK, enabling Http client and ASP.NET Core instrumentation
        // Exports to a locally running collector on port 4317
        services.AddOpenTelemetryTracing(options =>
        {
            options
                .AddSource("csharp.aspnet.api")
                .SetResourceBuilder(ResourceBuilder.CreateDefault().AddService("csharp.aspnet.api"))
                .SetSampler(new AlwaysOnSampler())
                .AddHttpClientInstrumentation()
                .AddAspNetCoreInstrumentation()
                .AddOtlpExporter();
        });
    }

    public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
    {
        app.UseHttpsRedirection();
        app.UseRouting();
        app.UseEndpoints(endpoints => { endpoints.MapControllers(); });
    }
}
