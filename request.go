package fresbi

import (
	"bytes"
	"encoding/json"
	"io"
)

type bulkRequest struct {
	buf *bytes.Buffer
}

func newBulkRequest() *bulkRequest {
	return &bulkRequest{}
}

func (br *bulkRequest) Buffer() *bytes.Buffer {
	return br.buf
}

func (br *bulkRequest) Reset() {
	br.buf.Reset()
}

// Index ...
func (br *bulkRequest) Index(item *BulkItem) error {
	item.action = "index"
	return br.addItem(item)
}

// Create ...
func (br *bulkRequest) Create(item *BulkItem) error {
	item.action = "create"
	return br.addItem(item)
}

// Update ...
func (br *bulkRequest) Update(item *BulkItem) error {
	item.action = "update"
	return br.addItem(item)
}

// Delete ...
func (br *bulkRequest) Delete(docID string) error {
	item.action = "delete"
	return br.addItem(item)
}

func (br *bulkRequest) addItem(item *BulkItem) error {
	if err := br.writeMeta(item); err != nil {
		return err
	}
	if err := br.writeBody(item); err != nil {
		return err
	}
	return nil
}

func (br *bulkRequest) writeMeta(item *BulkItem) error {
	b, err := json.Marshal(item)
	if err != nil {
		return err
	}

	br.buf.WriteString(`{"`)
	br.buf.WriteString(item.action)
	br.buf.WriteString(`"}:`)
	_, _ = br.buf.Write(b)
	br.buf.WriteString(`}\n`)
	return nil
}

func (br *bulkRequest) writeBody(item *BulkItem) error {
	var body []byte

	switch data := item.Body.(type) {
	case []byte:
		body = data
	case string:
		body = []byte(data)

	case io.Reader:
		_, err := br.buf.ReadFrom(data)
		if err != nil {
			return err
		}
	default:
		err := json.NewEncoder(br.buf).Encode(data)
		if err != nil {
			return err
		}
	}

	// true only when item.Body is []byte or string
	if body != nil {
		_, _ = br.buf.Write(body)
	}
	br.buf.WriteByte('\n')
	return nil
}
