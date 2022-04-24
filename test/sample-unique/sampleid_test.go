package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
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

var language = flag.String("language", "none", "The name of the language folder containing the recipefile.json to be validated")

func TestSampleIdUnique(t *testing.T) {
	cwd, _ := os.Getwd()
	root := filepath.Clean(filepath.Join(cwd, "..", ".."))

	fp := filepath.Join(root, "src", *language, "recipefile.json")

	t.Logf("The recipefile path to load is: %s", fp)

	jsonFile, err := os.Open(fp)
	if err != nil {
		t.Fatalf("Failed opening recipe json file: %v", err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var recipe Recipe
	json.Unmarshal(byteValue, &recipe)

	for _, s := range recipe.Signals {
		if hasDuplicatedSamples(s.Samples, t) {
			t.Errorf("Signal %s has duplicated sample ids. Check the logs to find the culprit", s.Id)
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
