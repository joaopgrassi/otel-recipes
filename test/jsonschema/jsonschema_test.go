package main

import (
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/xeipuuv/gojsonschema"
)

var recipePath = flag.String("path", "none", "The path of the folder containing the recipefile.json to be validated")

func TestJsonSchema(t *testing.T) {
	cwd, _ := os.Getwd()
	root := filepath.Clean(filepath.Join(cwd, "..", ".."))

	schemaPath := "file://" + filepath.Join(root, "otel-recipes-schema.json")
	recipeFile := "file://" + *recipePath

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
