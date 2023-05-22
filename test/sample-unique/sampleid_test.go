package main

import (
	"encoding/json"
	"flag"
	"io"
	"os"
	"path/filepath"
	"testing"
)

type Recipe struct {
	LanguageId string
	Signals    []Signal
}

type Signal struct {
	Id      string
	Samples []Sample
}

type Sample struct {
	Id string
}

var lang = flag.String("lang", "none", "The name of the lang folder")

var telemetrySignals = []string{"trace", "metrics", "logs"}

func TestSampleIdUnique(t *testing.T) {
	cwd, _ := os.Getwd()
	root := filepath.Clean(filepath.Join(cwd, "..", ".."))

	for _, ts := range telemetrySignals {
		sp := filepath.Join(root, "src", *lang, ts)
		if _, err := os.Stat(sp); os.IsNotExist(err) {
			t.Logf("No recipes yet for telemetry signal: %s", ts)
			continue
		}

		fp := filepath.Join(sp, "recipefile.json")
		t.Logf("The recipefile path to load is: %s", fp)

		jsonFile, err := os.Open(fp)
		if err != nil {
			t.Fatalf("Failed opening recipe json file: %v", err)
		}
		defer jsonFile.Close()

		byteValue, _ := io.ReadAll(jsonFile)
		var recipe Recipe
		json.Unmarshal(byteValue, &recipe)

		for _, s := range recipe.Signals {
			if hasDuplicatedSamples(s.Samples, t) {
				t.Errorf("Signal %s has duplicated sample ids. Check the logs to find the culprit", s.Id)
			}
		}
	}
}

func hasDuplicatedSamples(samples []Sample, t *testing.T) bool {
	ids := make(map[string]bool, 0)
	for _, s := range samples {
		if ids[s.Id] == true {
			t.Logf("Found duplicated sample with id: %s", s.Id)
			return true
		}
		ids[s.Id] = true
	}
	return false
}
