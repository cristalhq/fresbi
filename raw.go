package fresbi

import (
	"context"
	"net/http"
)

// RawClient ...
type RawClient struct {
	client *http.Client
	url    string
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
func (rc *RawClient) Flush(ctx context.Context) error {
	req, errReq := makeRequest(ctx, rc.url, rc.req.Buffer())
	if errReq != nil {
		return errReq
	}

	resp, errResp := rc.client.Do(req)
	if errResp != nil {
		return errResp
	}
	defer func() { _ = resp.Body.Close() }()

	return nil
}

// Index ...
func (rc *RawClient) Index(index, docID string, data interface{}) error {
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
