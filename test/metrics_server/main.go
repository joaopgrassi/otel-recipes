package main

import (
	"github.com/golang/protobuf/proto"
	colmetricspb "go.opentelemetry.io/proto/otlp/collector/metrics/v1"
	v1 "go.opentelemetry.io/proto/otlp/metrics/v1"
	"io"
	"log"
	"net/http"
)

var metrics map[string]*v1.Metric

// Dummy OTLP HTTP receiver that stores all metrics in-memory
// used by the tests to assert the metrics produced by the sample apps
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
			for _, m := range sm.GetMetrics() {
				metrics[m.GetName()] = m
			}
		}
	}
}

// searches for the provided metric in the in-memory collection
func getNumberDps(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	metricName := r.URL.Query().Get("metricName")
	if metricName == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing metricName parameter"))
		return
	}

	if m, found := metrics[metricName]; found {
		data, err := proto.Marshal(m)
		if err != nil {
			log.Fatal("marshaling error: ", err)
		}
		w.Write(data)
	}
}

func main() {
	metrics = make(map[string]*v1.Metric)

	// POST endpoint called by the OTLP exporter in the collector
	http.HandleFunc("/v1/metrics", postMetric)

	// GET endpoint called by the tests to assert the exported metrics
	http.HandleFunc("/getMetric", getNumberDps)

	http.ListenAndServe(":4319", nil)
}
