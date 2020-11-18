package fresbi

import (
	"strings"
	"testing"
)

func TestBulkRequest(t *testing.T) {
	req := newBulkRequest()

	var err error

	err = req.Index(&Item{
		ID:     "index-id",
		Parent: "test-parent",
		Body:   "index-raw-string-body",
	})
	if err != nil {
		t.Error(err)
	}

	err = req.Create(&Item{
		ID:          "create-id",
		VersionType: "test-version",
		Body:        []byte("create-byte-slice-body"),
	})
	if err != nil {
		t.Error(err)
	}

	err = req.Update(&Item{
		ID:       "update-id",
		Index:    "test-index",
		Pipeline: "test-pipeline",
		Body:     strings.NewReader("body from io.Reader"),
	})
	if err != nil {
		t.Error(err)
	}

	err = req.Delete(&Item{
		ID: "delete-id",
	})
	if err != nil {
		t.Error(err)
	}

	got := req.Buffer().String()

	want := `{"index":{"_id":"index-id","parent":"test-parent"}}
index-raw-string-body
{"create":{"_id":"create-id","version_type":"test-version"}}
create-byte-slice-body
{"update":{"_index":"test-index","_id":"update-id","pipeline":"test-pipeline"}}
body from io.Reader
{"delete":{"_id":"delete-id"}}
`

	if want != got {
		t.Fatalf("want %q, got %q", want, got)
	}
}
