package testutils // import "github.com/joaopgrassi/otel-recipes/internal/common/testutils"

// Address of the OTLP back-end running inside compose
const OtlpBackendUri string = "http://otlp-backend:4319"

// Constants for signals
const TraceSignal string = "trace"
const MetricsSignal string = "metrics"
const LogsSignal string = "logs"
