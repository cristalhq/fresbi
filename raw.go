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

// NewRawClient instantiates a new RawClient.
func NewRawClient(client Doer, config Config) *RawClient {
	return &RawClient{
		client: newBulkClient(client, config),
		req:    newBulkRequest(),
	}
}

// Reset the current buffer.
func (rc *RawClient) Reset() {
	rc.req.Reset()
}

// Send the given buffer as a bulk request.
func (rc *RawClient) Send(ctx context.Context) (*http.Response, error) {
	resp, errResp := rc.client.send(ctx, rc.req.Buffer())
	if errResp != nil {
		return nil, errResp
	}
	return resp, nil
}

// Index a document in a bulk request.
func (rc *RawClient) Index(item *Item) error {
	return rc.req.Index(item)
}

// Create a document in a bulk request.
func (rc *RawClient) Create(item *Item) error {
	return rc.req.Create(item)
}

// Update a document in a bulk request.
func (rc *RawClient) Update(item *Item) error {
	return rc.req.Update(item)
}

// Delete a document in a bulk request.
func (rc *RawClient) Delete(item *Item) error {
	return rc.req.Delete(item)
}
