package fresbi

import (
	"context"
	"net/http"
)

// RawClient ...
type RawClient struct {
	client *http.Client
	url    string
	config Config
	req    *bulkRequest
}

// NewRawClient ...
func NewRawClient(client *http.Client) *RawClient {
	return &RawClient{}
}

// Reset ...
func (rc *RawClient) Reset() {
	rc.req.Reset()
}

// Flush ...
func (rc *RawClient) Flush(ctx context.Context) (*Response, error) {
	req, errReq := makeRequest(ctx, rc.url, rc.req.Buffer())
	if errReq != nil {
		return nil, errReq
	}

	resp, errResp := send(ctx, rc.client, rc.url, rc.req.Buffer())
	if errResp != nil {
		return nil, errResp
	}
	return resp, nil
}

// Index ...
func (rc *RawClient) Index(item *IndexItem) error {
	item.op = "index"
	return rc.req.Index(docID, data)
}

// Create ...
func (rc *RawClient) Create(index, docID string, data interface{}) error {
	return rc.req.Create(docID, data)
}

// Update ...
func (rc *RawClient) Update(index, docID string, data interface{}) error {
	return rc.req.Update(docID, data)
}

// Delete ...
func (rc *RawClient) Delete(index, docID string) error {
	return rc.req.Delete(docID)
}
