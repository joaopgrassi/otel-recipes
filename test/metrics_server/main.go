package main

import (
	"github.com/golang/protobuf/proto"
	colmetricspb "go.opentelemetry.io/proto/otlp/collector/metrics/v1"
	v1 "go.opentelemetry.io/proto/otlp/metrics/v1"
	"io"
	"log"
	"net/http"
)

var scopeMetrics map[string]*v1.ScopeMetrics

// Dummy OTLP HTTP receiver that stores all scopeMetrics in-memory
// used by the tests to assert the scopeMetrics produced by the sample apps
func postMetric(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error reading payload"))
	}

	pbRequest := &colmetricspb.ExportMetricsServiceRequest{}
	err = proto.Unmarshal(body, pbRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error reading payload"))
	}

	for _, rm := range pbRequest.GetResourceMetrics() {
		for _, sm := range rm.GetScopeMetrics() {
			scopeMetrics[sm.Scope.GetName()] = sm
		}
	}
}

// searches for the provided metric in the in-memory collection
func getScopeMetrics(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	scopeName := r.URL.Query().Get("scopeName")
	if scopeName == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing scopeName parameter"))
		return
	}

	if sm, found := scopeMetrics[scopeName]; found {
		data, err := proto.Marshal(sm)
		if err != nil {
			log.Fatal("marshaling error: ", err)
		}
		w.Write(data)
	}
}

func main() {
	scopeMetrics = make(map[string]*v1.ScopeMetrics)

	// POST endpoint called by the OTLP exporter in the collector
	http.HandleFunc("/v1/metrics", postMetric)

	// GET endpoint called by the tests to assert the exported scopeMetrics
	http.HandleFunc("/getMetric", getScopeMetrics)

	http.ListenAndServe(":4319", nil)
}
