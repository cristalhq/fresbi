package fresbi

import (
	"bytes"
	"context"
	"testing"
)

func Test_makeRequest(t *testing.T) {
	cfg := Config{
		URL:                 "test-url",
		BatchSize:           123,
		Pipeline:            "test-pipeline",
		Refresh:             "test-refresh",
		Routing:             "test-routing",
		ErrorTrace:          nil,
		FilterPath:          []string{},
		Timeout:             "test-timeout",
		WaitForActiveShards: "test-waitForActiveShards",
	}
	c := newBulkClient(nil, cfg)

	req, err := c.makeRequest(context.Background(), &bytes.Buffer{})
	if err != nil {
		t.Fatal(err)
	}

	if v := req.Header.Get("Content-Type"); v != "application/x-ndjson" {
		t.Errorf("want %s, got %v", "application/x-ndjson", v)
	}
}
