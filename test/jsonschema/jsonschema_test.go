package main

import (
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/xeipuuv/gojsonschema"
)

var lang = flag.String("lang", "none", "The name of the lang folder containing the recipefile.json to be validated")

var telemetrySignals = []string{"trace", "metrics", "logs"}

func TestJsonSchema(t *testing.T) {
	cwd, _ := os.Getwd()
	root := filepath.Clean(filepath.Join(cwd, "..", ".."))

	for _, ts := range telemetrySignals {
		sp := filepath.Join(root, "src", *lang, ts)
		if _, err := os.Stat(sp); os.IsNotExist(err) {
			t.Logf("No recipes yet for telemetry signal: %s", ts)
			continue
		}

		schemaPath := "file://" + filepath.Join(root, "otel-recipes-schema.json")
		recipeFile := "file://" + filepath.Join(sp, "recipefile.json")

		t.Logf("The schema path to load is: %s", schemaPath)
		t.Logf("The recipefile path to load is: %s", recipeFile)

		schemaLoader := gojsonschema.NewReferenceLoader(schemaPath)
		documentLoader := gojsonschema.NewReferenceLoader(recipeFile)

		result, err := gojsonschema.Validate(schemaLoader, documentLoader)
		if err != nil {
			t.Fatalf("Failed validating the JSON schema: %v", err)
		}

		if !result.Valid() {
			t.Errorf("The JSON schema is not valid. See errors:")
			for _, desc := range result.Errors() {
				t.Logf("- %s", desc)
			}
		}
	}
}
