package fresbi

import (
	"bytes"
	"encoding/json"
	"io"
	"strconv"
)

type bulkRequest struct {
	buf *bytes.Buffer
}

func newBulkRequest(url string, cfg Config) *bulkRequest {
	return &bulkRequest{}
}

func (br *bulkRequest) Buffer() *bytes.Buffer {
	return br.buf
}

func (br *bulkRequest) Reset() {
	br.buf.Reset()
}

// Index ...
func (br *bulkRequest) Index(docID string, data interface{}) error {
	return br.addItem(&bulkIndexerItem{
		Action:     "index",
		DocumentID: docID,
		Body:       data,
	})
}

// Create ...
func (br *bulkRequest) Create(docID string, data interface{}) error {
	return br.addItem(&bulkIndexerItem{
		Action:     "create",
		DocumentID: docID,
		Body:       data,
	})
}

// Update ...
func (br *bulkRequest) Update(docID string, data interface{}) error {
	return br.addItem(&bulkIndexerItem{
		Action:     "update",
		DocumentID: docID,
		Body:       data,
	})
}

// Delete ...
func (br *bulkRequest) Delete(docID string) error {
	return br.addItem(&bulkIndexerItem{
		Action:     "delete",
		DocumentID: docID,
	})
}

func (br *bulkRequest) addItem(item *bulkIndexerItem) error {
	br.writeMeta(item)
	return br.writeBody(item)
}

func (br *bulkRequest) writeMeta(item *bulkIndexerItem) {
	var aux []byte

	br.buf.WriteByte('{')
	aux = strconv.AppendQuote(aux, item.Action)
	br.buf.Write(aux)
	aux = aux[:0]

	br.buf.WriteString(":{")
	// br.buf.WriteByte(':')
	// br.buf.WriteByte('{')

	if item.DocumentID != "" {
		br.buf.WriteString(`"_id":`)
		aux = strconv.AppendQuote(aux, item.DocumentID)
		br.buf.Write(aux)
		aux = aux[:0]
	}

	if item.Index != "" {
		if item.DocumentID != "" {
			br.buf.WriteByte(',')
		}
		br.buf.WriteString(`"_index":`)
		aux = strconv.AppendQuote(aux, item.Index)
		br.buf.Write(aux)
		aux = aux[:0]
	}

	br.buf.WriteString("}}\n")
	// br.buf.WriteByte('}')
	// br.buf.WriteByte('}')
	// br.buf.WriteByte('\n')
}

func (br *bulkRequest) writeBody(item *bulkIndexerItem) error {
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
