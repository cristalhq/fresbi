package fresbi

import (
	"context"
	"net/http"
)

// ReliableClient ...
// See: https://www.elastic.co/guide/en/elasticsearch/reference/master/docs-bulk.html
//
type ReliableClient struct {
	client *bulkClient
}

// NewReliableClient ...
func NewReliableClient(client Doer, config Config) *ReliableClient {
	return &ReliableClient{
		client: newBulkClient(client, config),
	}
}

// Batch ...
type Batch interface {
	Index(item *Item) error
	Create(item *Item) error
	Update(item *Item) error
	Delete(item *Item) error
}

// AsBatch ...
func (c *ReliableClient) AsBatch(ctx context.Context, fn func(Batch) error) (*http.Response, error) {
	req := newBulkRequest()
	defer req.Reset()

	if err := fn(req); err != nil {
		return nil, err
	}

	resp, errResp := c.client.send(ctx, req.Buffer())
	if errResp != nil {
		return nil, errResp
	}
	return resp, nil
}
