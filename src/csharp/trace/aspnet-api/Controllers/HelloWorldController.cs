using Microsoft.AspNetCore.Mvc;

namespace AspNetCoreApi.Controllers;

[ApiController]
[Route("helloworld")]
public class HelloWorldController : Controller
{
    [HttpGet]
    public IActionResult Get()
    {
        using var activity = Startup.Tracer.StartActivity("HelloWorldSpan");
        activity?.SetTag("foo", "bar");
        return Ok("Hello world!");
    }
}
