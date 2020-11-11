package fresbi

import (
	"context"
	"net/http"
)

// Client ...
// See: https://www.elastic.co/guide/en/elasticsearch/reference/master/docs-bulk.html
//
type Client struct {
	client *http.Client
	url    string
	config Config
}

// Config ...
type Config struct {
	batchSize           int
	pipeline            string
	refresh             string
	routing             string
	errorTrace          *interface{}
	filterPath          []string
	timeout             string
	waitForActiveShards string
}

// Stats ...
type Stats struct {
	NumAdded    uint64
	NumFlushed  uint64
	NumFailed   uint64
	NumIndexed  uint64
	NumCreated  uint64
	NumUpdated  uint64
	NumDeleted  uint64
	NumRequests uint64
}

// NewClient ...
func NewClient(client *http.Client) *Client {
	return &Client{
		client: client,
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
func (c *Client) AsBatch(ctx context.Context, fn func(Batch) error) (*Response, error) {
	req := newBulkRequest()

	if err := fn(req); err != nil {
		req.Reset()
		return nil, err
	}

	resp, errResp := send(ctx, c.client, c.url, req.Buffer())
	if errResp != nil {
		return nil, errResp
	}
	return resp, nil
}
