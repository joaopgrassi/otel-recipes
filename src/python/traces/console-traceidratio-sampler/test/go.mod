module github.com/joaopgrassi/otel-recipes/python/traces/console-traceidratio-sampler

go 1.22.1

require (
	github.com/joaopgrassi/otel-recipes/internal/common v0.0.0
	github.com/stretchr/testify v1.9.0
	go.opentelemetry.io/proto/otlp v1.2.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	google.golang.org/protobuf v1.34.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/joaopgrassi/otel-recipes/internal/common v0.0.0 => ../../../../../internal/common
