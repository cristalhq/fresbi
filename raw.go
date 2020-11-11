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
	resp, errResp := send(ctx, rc.client, rc.url, rc.req.Buffer())
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
