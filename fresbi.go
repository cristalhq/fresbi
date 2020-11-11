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
	URL                 string
	BatchSize           int
	Pipeline            string
	Refresh             string
	Routing             string
	ErrorTrace          *interface{}
	FilterPath          []string
	Timeout             string
	WaitForActiveShards string
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
func NewClient(client *http.Client, config Config) *Client {
	return &Client{
		client: client,
		config: config,
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

	resp, errResp := send(ctx, c.client, c.config.URL, req.Buffer())
	if errResp != nil {
		return nil, errResp
	}
	return resp, nil
}
