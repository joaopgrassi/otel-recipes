package testutils // import "github.com/joaopgrassi/otel-recipes/internal/common/testutils"

import (
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	otlplogs "go.opentelemetry.io/proto/otlp/logs/v1"
	"google.golang.org/protobuf/proto"
)

func AssertLogWithAttributeExists(t *testing.T, tc *LogTestCase) {
	backoffSchedule := []time.Duration{
		1 * time.Second,
		3 * time.Second,
		10 * time.Second,
		15 * time.Second,
		20 * time.Second,
		30 * time.Second,
	}

	var actual *otlplogs.LogRecord
	for _, backoff := range backoffSchedule {
		logs := GetLogsWithRetry(t, tc.serviceName)
		log := findLog(logs, tc.body)

		if log != nil {
			actual = log
			break
		}

		t.Logf("Log not found yet, retrying in %v\n", backoff)
		time.Sleep(backoff)
	}

	// assert
	assert.Equal(t, tc.severity, actual.GetSeverityText())

	if tc.withTrace {
		assert.NotEmpty(t, actual.GetTraceId())
		assert.NotEmpty(t, actual.GetSpanId())
	}

	for _, exp := range tc.attributes {
		assert.Contains(t, actual.Attributes, exp)
	}
}

func findLog(logs *otlplogs.ResourceLogs, body string) *otlplogs.LogRecord {
	for _, sl := range logs.ScopeLogs {
		for _, l := range sl.LogRecords {
			if l.Body.GetStringValue() == body {
				return l
			}
		}
	}
	return nil
}

func GetLogsWithRetry(t *testing.T, serviceName string) *otlplogs.ResourceLogs {
	backoffSchedule := []time.Duration{
		1 * time.Second,
		3 * time.Second,
		10 * time.Second,
		15 * time.Second,
		20 * time.Second,
		30 * time.Second,
	}

	var rl *otlplogs.ResourceLogs

	// do some retries until we backend has it
	for _, backoff := range backoffSchedule {
		rl = GetLog(t, serviceName)

		if rl != nil {
			break
		}

		t.Logf("Log not found yet, retrying in %v\n", backoff)
		time.Sleep(backoff)
	}

	if len(rl.ScopeLogs) == 0 {
		t.Fatalf("Could not find logs for sample: %s", serviceName)
	}

	return rl
}

func GetLog(t *testing.T, serviceName string) *otlplogs.ResourceLogs {
	t.Logf("Going to call OTLP backend to fetch logs for sample: %s", serviceName)
	r, err := http.Get(fmt.Sprintf("%s/getotlp?signal=logs&servicename=%s", OtlpBackendUri, serviceName))
	if err != nil {
		t.Fatalf("Failed getting logs from OTLP backend: %v", err)
	}

	t.Log("Received 200 response from OTLP backend")

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		t.Fatalf("Error reading payload from OTLP backend: %v", err)
	}

	if len(body) == 0 {
		return nil
	}

	rl := &otlplogs.ResourceLogs{}
	err = proto.Unmarshal(body, rl)
	if err != nil {
		t.Fatalf("Error reading payload from OTLP backend: %v", err)
	}
	return rl
}
