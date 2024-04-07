using System.Diagnostics;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;

namespace AspNetCoreApi.Controllers;

[ApiController]
[Route("helloworld")]
public class HelloWorldController : Controller
{
    private readonly ILogger<HelloWorldController> _logger;

    public HelloWorldController(ILogger<HelloWorldController> logger)
    {
        _logger = logger;
    }

    [HttpGet]
    public IActionResult Get()
    {
        // The log will be correlated with the current TraceId/SpanId
        _logger.LogInformation("This is a info message {foo}", "bar");
        return Ok("Hello world!");
    }
}
