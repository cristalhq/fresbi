package fresbi

import (
	"context"
	"net/http"
)

// Client ...
// See: https://www.elastic.co/guide/en/elasticsearch/reference/master/docs-bulk.html
//
type Client struct {
	client *bulkClient
}

// NewClient ...
func NewClient(url string, client Doer, config *Config) *Client {
	return &Client{
		client: newBulkClient(url, client, config),
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
func (c *Client) AsBatch(ctx context.Context, fn func(Batch) error) (*http.Response, error) {
	req := newBulkRequest()

	if err := fn(req); err != nil {
		req.Reset()
		return nil, err
	}

	resp, errResp := c.client.send(ctx, req.Buffer())
	if errResp != nil {
		return nil, errResp
	}
	return resp, nil
}
