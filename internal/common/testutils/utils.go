package testutils // import "github.com/joaopgrassi/otel-recipes/internal/common/testutils"

import (
	"io"
	"net/http"
	"testing"
)

func InvokeSampleApi(t *testing.T, url string) string {
	t.Logf("Going to call the sample API: %s", url)
	r, err := http.Get(url)
	if err != nil {
		t.Fatalf("Failed calling the sample API: %v", err)
	}

	t.Log("Received 200 response from the sample API")

	defer r.Body.Close()

	//We Read the response body on the line below.
	body, err := io.ReadAll(io.Reader(r.Body))
	if err != nil {
		t.Fatalf("Failed reading response body from the sample API: %v", err)
	}

	return string(body)
}
