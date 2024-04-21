# Test utils

This folder contains a set of utilities to aid in building and running the integration/e2e tests
for the sample applications.

The utilities can be split into:

- `types.go`: The interfaces used for create the test cases
- `trace.go`, `metrics.go`, `logs.go`: The utilities to fetch the telemetry and perform assertion
  based on the provided `TestCase`

## Writing tests

Each recipe application must contain a go module `test`. The module then depends on the `testutils`.
The application test must declare the `TestCase` and invoke the methods available for each signal.

1. Create the `test` go module at the root of the recipe application folder
2. In `go.mod` set the dependencies to the `testutils` module. For example

```go
module github.com/joaopgrassi/otel-recipes/csharp/trace/console

go 1.22.1

require github.com/joaopgrassi/otel-recipes/internal/common v0.0.0

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	go.opentelemetry.io/proto/otlp v1.1.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/joaopgrassi/otel-recipes/internal/common v0.0.0 => ../../../../../internal/common
```
Once you have the test module ready, simple add a new file containing your test.
As file name convention for the tests, please use: `<signal>_test.go`. E.g., `traces_test.go`.

### Trigger telemetry generation

During CI, the recipe application is automatically started.

Depending on the recipe application, the "to be asserted" telemetry may be automatically
generated upon app startup. For example, for applications that are "console-like" there's no need
to invoke it to generate the telemetry.

But some other types of application may need to be invoke. For example REST APIs. For this, the
`testutils` module offers a convenient way to invoke the sample application.

You can use the `InvokeSampleApi` function from [utils.go](./testutils/utils.go) to call
the API endpoint that will generate the telemetry to be asserted. `InvokeSampleApi`.

For example, to invoke an recipe API that generates a log record in a `helloworld` GET endpoint:

```go
package test

import (
	"testing"

	tu "github.com/joaopgrassi/otel-recipes/internal/common/testutils"
)

func TestLogGeneratedFromSample(t *testing.T) {
	tu.InvokeSampleApi(t, "http://localhost:8080/helloworld")

	tc := tu.NewLogTestCase("csharp.aspnet.api", "Information", "This is a info message {foo}", tu.StringAttribute("foo", "bar"))

	tu.AssertLogWithAttributeExists(t, tc)
}
```

Once the recipe application is triggered, you can use the test utils to assert the data.
As telemetry export may take a while until it reaches the OTLP back-end, the assertion
methods are prepared with retries until the data is found.

### Trace tests

For recipe applications that uses traces, an example test that checks for a span
with name `HelloWorldSpan` and an attribute `foo=bar` you can write:

```go
package test

import (
	"testing"

	tu "github.com/joaopgrassi/otel-recipes/internal/common/testutils"
)

func TestTraceGeneratedFromSample(t *testing.T) {
	tc := tu.NewTraceTestCase("csharp.console.app", "HelloWorldSpan", tu.StringAttribute("foo", "bar"))

	tu.AssertSpanWithAttributeExists(t, tc)
}
```

### Metric tests

For recipe applications that uses metrics, an example test that checks for `Counter` and `Gauge` metrics
with an attribute `foo=bar` you can write:

```go
package test

import (
	"testing"

	tu "github.com/joaopgrassi/otel-recipes/internal/common/testutils"
)

func TestMetricsGeneratedFromSample(t *testing.T) {
	rm := tu.GetMetricsWithRetry(t, "csharp.console.app")
	m := rm.GetScopeMetrics()[0].Metrics

	// Counter metric
	ctc := tu.NewMetricTestCase("myCounter", "I count things", "1", int64(3), tu.StringAttribute("foo", "bar"))
	tu.AssertCounter(t, ctc, m)

	// Gauge metric
	ctg := tu.NewMetricTestCase("myGauge", "I gauge things", "1", float64(3.5), tu.StringAttribute("foo", "bar"))
	tu.AssertGauge(t, ctg, m)
}
```

### Log tests

For recipe applications that uses logs, an example test that checks for a log record
with an attribute `foo=bar` you can write:

```go
package test

import (
	"testing"

	tu "github.com/joaopgrassi/otel-recipes/internal/common/testutils"
)

func TestLogGeneratedFromSample(t *testing.T) {
	tu.InvokeSampleApi(t, "http://localhost:8080/helloworld")

	tc := tu.NewLogTestCase("csharp.aspnet.api", "Information", "This is a info message {foo}", tu.StringAttribute("foo", "bar"))

	tu.AssertLogWithAttributeExists(t, tc)
}
```
