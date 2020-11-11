package fresbi

import (
	"bytes"
	"context"
	"net/http"
)

// Client ...
// See: https://www.elastic.co/guide/en/elasticsearch/reference/master/docs-bulk.html
//
type Client struct {
	url string
	// buf    *bytes.Buffer
	// aux    []byte
	client *http.Client
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

// BulkRequest ...
type BulkRequest interface {
	Index(docID string, data interface{}) error
	Create(docID string, data interface{}) error
	Update(docID string, data interface{}) error
	Delete(docID string) error
}

// AsBatch ...
func (c *Client) AsBatch(ctx context.Context, fn func(BulkRequest) error) error {
	req := newBulkRequest(c.url, c.config)

	if err := fn(req); err != nil {
		req.Reset()
		return err
	}
	if err := c.flushBuffer(ctx, req); err != nil {
		return err
	}
	return nil
}

// flushBuffer ...
func (c *Client) flushBuffer(ctx context.Context, br *bulkRequest) error {
	req, errReq := makeRequest(ctx, c.url, br.Buffer())
	if errReq != nil {
		return errReq
	}

	resp, errResp := c.client.Do(req)
	if errResp != nil {
		return errResp
	}
	defer func() { _ = resp.Body.Close() }()

	return nil
}

func makeRequest(ctx context.Context, url string, buf *bytes.Buffer) (*http.Request, error) {
	req, errReq := http.NewRequestWithContext(ctx, http.MethodPost, url, buf)
	if errReq != nil {
		return nil, errReq
	}
	req.Header.Add("Content-Type", "application/x-ndjson")

	return req, nil
}

// params := url.Values{}
// if c.config.pipeline != "" {
// 	params.Set("pipeline", c.config.pipeline)
// }
// if c.config.refresh != "" {
// 	params.Set("refresh", c.config.refresh)
// }
// if c.config.routing != "" {
// 	params.Set("routing", c.config.routing)
// }
// if v := s.pretty; v != nil {
// 	params.Set("pretty", fmt.Sprint(*v))
// }
// if v := s.human; v != nil {
// 	params.Set("human", fmt.Sprint(*v))
// }
// if v := c.config.errorTrace; v != nil {
// 	params.Set("error_trace", fmt.Sprint(*v))
// }
// if len(c.config.filterPath) > 0 {
// 	params.Set("filter_path", strings.Join(c.config.filterPath, ","))
// }
// if c.config.timeout != "" {
// 	params.Set("timeout", c.config.timeout)
// }
// if c.config.waitForActiveShards != "" {
// 	params.Set("wait_for_active_shards", c.config.waitForActiveShards)
// }
