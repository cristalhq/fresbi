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

	var body = struct {
		Foo string `json:"foo"`
		Bar int64  `json:"bar"`
	}{
		Foo: "test-value",
		Bar: 42,
	}
	err = req.Create(&Item{
		ID:          "create-id-with-field",
		VersionType: "test-version",
		Body:        body,
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

	want := `{"index":{"_id":"index-id","parent":"test-parent"}}` + "\n" +
		`index-raw-string-body` + "\n" +
		`{"create":{"_id":"create-id","version_type":"test-version"}}` + "\n" +
		`create-byte-slice-body` + "\n" +
		`{"create":{"_id":"create-id-with-field","version_type":"test-version"}}` + "\n" +
		`{"foo":"test-value","bar":42}` + "\n" +
		`{"update":{"_index":"test-index","_id":"update-id","pipeline":"test-pipeline"}}` + "\n" +
		`body from io.Reader` + "\n" +
		`{"delete":{"_id":"delete-id"}}` + "\n"

	if want != got {
		t.Errorf("want %v %q, got %v %q", len(want), want, len(got), got)
	}
}
