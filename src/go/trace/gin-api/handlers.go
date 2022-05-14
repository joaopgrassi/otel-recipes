package main

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// GetHelloWorld Handles calls to /helloworld
func GetHelloWorld(c *gin.Context) {

	// Starts a span with an attribute
	_, span := Tracer.Start(
		c.Request.Context(),
		"HelloWorldSpan",
		trace.WithAttributes(attribute.String("foo", "bar")))
	defer span.End()

	c.String(200, "Hello world!")
}
