package com.otel.recipes.springbootapi;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import io.opentelemetry.api.OpenTelemetry;
import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.Tracer;

@RestController()
@RequestMapping("helloworld")
public class HelloWorldController {

	private final Tracer tracer;

	public HelloWorldController(OpenTelemetry openTelemetry) {
		// Obtain the OpenTelemetry trace provider from DI
		// and use it to create the tracer
		this.tracer = openTelemetry.getTracer("java.springboot.api");
	}

	@GetMapping()
	public String Get() {
		// Start a span with an attribute
		Span span = tracer.spanBuilder("HelloWorldSpan").startSpan();
		span.setAttribute("foo", "bar");
		span.end();
		return "Hello world!";
	}
}
