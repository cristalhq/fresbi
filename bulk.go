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
	return &bulkRequest{
		buf: &bytes.Buffer{},
	}
}

func (br *bulkRequest) Buffer() *bytes.Buffer {
	return br.buf
}

func (br *bulkRequest) Reset() {
	br.buf.Reset()
}

// Index ...
func (br *bulkRequest) Index(item *Item) error {
	item.action = "index"
	return br.addItem(item)
}

// Create ...
func (br *bulkRequest) Create(item *Item) error {
	item.action = "create"
	return br.addItem(item)
}

// Update ...
func (br *bulkRequest) Update(item *Item) error {
	item.action = "update"
	return br.addItem(item)
}

// Delete ...
func (br *bulkRequest) Delete(item *Item) error {
	item.action = "delete"
	return br.addItem(item)
}

func (br *bulkRequest) addItem(item *Item) error {
	if err := br.writeMeta(item); err != nil {
		return err
	}
	if err := br.writeBody(item); err != nil {
		return err
	}
	return nil
}

func (br *bulkRequest) writeMeta(item *Item) error {
	br.buf.WriteString(`{"`)
	br.buf.WriteString(item.action)
	br.buf.WriteString(`":`)

	b, err := json.Marshal(item)
	if err != nil {
		return err
	}
	_, _ = br.buf.Write(b)

	br.buf.WriteString("}\n")
	return nil
}

func (br *bulkRequest) writeBody(item *Item) error {
	if item.action == "delete" { // doesn't have body, only meta
		return nil
	}

	switch data := item.Body.(type) {
	case []byte:
		_, _ = br.buf.Write(data)
		br.buf.WriteByte('\n')
		return nil

	case string:
		_, _ = br.buf.WriteString(data)
		br.buf.WriteByte('\n')
		return nil

	case io.Reader:
		_, err := br.buf.ReadFrom(data)
		if err != nil {
			return err
		}
		br.buf.WriteByte('\n')
		return nil

	default:
		return json.NewEncoder(br.buf).Encode(data)
	}
}
