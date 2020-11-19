package fresbi

import (
	"bytes"
	"context"
	"net/http"
)

type bulkClient struct {
	url    string
	client *http.Client
	config *Config
}

func newBulkClient(url string, client *http.Client, config *Config) *bulkClient {
	return &bulkClient{
		url:    url,
		client: client,
		config: config,
	}
}

func (bc *bulkClient) send(ctx context.Context, buf *bytes.Buffer) (*http.Response, error) {
	req, errReq := bc.makeRequest(ctx, buf)
	if errReq != nil {
		return nil, errReq
	}

	resp, errResp := bc.client.Do(req)
	if errResp != nil {
		return nil, errResp
	}
	return resp, nil
}

func (bc *bulkClient) makeRequest(ctx context.Context, buf *bytes.Buffer) (*http.Request, error) {
	req, errReq := http.NewRequestWithContext(ctx, http.MethodPost, bc.url, buf)
	if errReq != nil {
		return nil, errReq
	}
	req.Header.Add("Content-Type", "application/x-ndjson")

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

	return req, nil
}
