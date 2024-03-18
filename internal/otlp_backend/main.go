package main

import (
	"io"
	"log/slog"
	"net/http"

	"github.com/golang/protobuf/proto"
	collogspb "go.opentelemetry.io/proto/otlp/collector/logs/v1"
	colmetricspb "go.opentelemetry.io/proto/otlp/collector/metrics/v1"
	coltracepb "go.opentelemetry.io/proto/otlp/collector/trace/v1"
	otlplogs "go.opentelemetry.io/proto/otlp/logs/v1"
	otlpmetrics "go.opentelemetry.io/proto/otlp/metrics/v1"
	otlptrace "go.opentelemetry.io/proto/otlp/trace/v1"
	"google.golang.org/protobuf/runtime/protoiface"
)

var resourceSpans map[string]*otlptrace.ResourceSpans
var resourceMetrics map[string]*otlpmetrics.ResourceMetrics
var resourceLogs map[string]*otlplogs.ResourceLogs

// OTLP HTTP receiver that stores all ResourceSpans in-memory
// used by the tests to assert the traces produced by the sample apps
func postTrace(w http.ResponseWriter, r *http.Request) {
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

	pbRequest := &coltracepb.ExportTraceServiceRequest{}
	err = proto.Unmarshal(body, pbRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error reading payload"))
	}

	for _, rs := range pbRequest.GetResourceSpans() {
		var sn string = ""
		for _, attr := range rs.Resource.Attributes {
			if attr.Key == "service.name" {
				sn = attr.GetValue().GetStringValue()
			}
		}

		if sn != "" {
			resourceSpans[sn] = rs
		} else {
			slog.Warn("Could not find service name attribute in OTLP resource spans")
		}
	}
}

// OTLP HTTP receiver that stores all ResourceMetrics in-memory
// used by the tests to assert the metrics produced by the sample apps
func postMetrics(w http.ResponseWriter, r *http.Request) {
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
		var sn string = ""
		for _, attr := range rm.Resource.Attributes {
			if attr.Key == "service.name" {
				sn = attr.GetValue().GetStringValue()
			}
		}

		if sn != "" {
			resourceMetrics[sn] = rm
		} else {
			slog.Warn("Could not find service name attribute in OTLP resource metrics")
		}
	}
}

// OTLP HTTP receiver that stores all ResourceLogs in-memory
// used by the tests to assert the logs produced by the sample apps
func postLogs(w http.ResponseWriter, r *http.Request) {
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

	pbRequest := &collogspb.ExportLogsServiceRequest{}
	err = proto.Unmarshal(body, pbRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error reading payload"))
	}

	for _, rl := range pbRequest.GetResourceLogs() {
		var sn string = ""
		for _, attr := range rl.Resource.Attributes {
			if attr.Key == "service.name" {
				sn = attr.GetValue().GetStringValue()
			}
		}

		if sn != "" {
			resourceLogs[sn] = rl
		} else {
			slog.Warn("Could not find service name attribute in OTLP resource logs")
		}
	}
}

// Gets the in-memory OTLP data, filtered by signal and service.name
func getOtlpData(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	signal := r.URL.Query().Get("signal")
	if signal == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing signal parameter"))
		return
	}

	serviceName := r.URL.Query().Get("servicename")
	if serviceName == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing serviceName parameter"))
		return
	}

	var res []byte
	switch signal {
	case "trace":
		if rm, found := resourceSpans[serviceName]; found {
			res = makeOtlpResponse(rm)
		}
	case "metrics":
		if rm, found := resourceMetrics[serviceName]; found {
			res = makeOtlpResponse(rm)
		}
	case "logs":
		if rm, found := resourceLogs[serviceName]; found {
			res = makeOtlpResponse(rm)
		}
	}

	if res == nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.Write(res)
	}
}

func makeOtlpResponse(m protoiface.MessageV1) []byte {
	if m == nil {
		slog.Info("m is nil")
		return nil
	}
	slog.Info("data: ", m)
	data, err := proto.Marshal(m)
	if err != nil {
		slog.Error("marshaling error: ", err)
		return nil
	}
	return data
}

func main() {
	resourceSpans = make(map[string]*otlptrace.ResourceSpans)
	resourceMetrics = make(map[string]*otlpmetrics.ResourceMetrics)
	resourceLogs = make(map[string]*otlplogs.ResourceLogs)

	http.HandleFunc("/v1/traces", postTrace)
	http.HandleFunc("/v1/metrics", postMetrics)
	http.HandleFunc("/v1/logs", postLogs)

	// GET endpoint called by the tests to assert the exported OTLP metrics filtered by signal and service.name
	http.HandleFunc("/getotlp", getOtlpData)

	http.ListenAndServe(":4319", nil)
}
