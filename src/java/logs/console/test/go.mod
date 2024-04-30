module github.com/joaopgrassi/otel-recipes/java/logs/console

go 1.22.1

require github.com/joaopgrassi/otel-recipes/internal/common v0.0.0

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	go.opentelemetry.io/proto/otlp v1.2.0 // indirect
	google.golang.org/protobuf v1.34.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/joaopgrassi/otel-recipes/internal/common v0.0.0 => ../../../../../internal/common
