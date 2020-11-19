package fresbi

import (
	"context"
	"net/http"
)

// RawClient ...
type RawClient struct {
	client *bulkClient
	req    *bulkRequest
}

// NewRawClient ...
func NewRawClient(url string, client *http.Client, config *Config) *RawClient {
	return &RawClient{
		client: newBulkClient(url, client, config),
		req:    newBulkRequest(),
	}
}

// Reset ...
func (rc *RawClient) Reset() {
	rc.req.Reset()
}

// Send ...
func (rc *RawClient) Send(ctx context.Context) (*http.Response, error) {
	resp, errResp := rc.client.send(ctx, rc.req.Buffer())
	if errResp != nil {
		return nil, errResp
	}
	return resp, nil
}

// Index ...
func (rc *RawClient) Index(item *Item) error {
	return rc.req.Index(item)
}

// Create ...
func (rc *RawClient) Create(item *Item) error {
	return rc.req.Create(item)
}

// Update ...
func (rc *RawClient) Update(item *Item) error {
	return rc.req.Update(item)
}

// Delete ...
func (rc *RawClient) Delete(item *Item) error {
	return rc.req.Delete(item)
}
