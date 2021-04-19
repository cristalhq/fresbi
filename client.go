package fresbi

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type bulkClient struct {
	client Doer
	config Config
}

func newBulkClient(client Doer, config Config) *bulkClient {
	return &bulkClient{
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
		if errResp != io.EOF {
			return nil, errResp
		}
	}
	return resp, nil
}

func (bc *bulkClient) makeRequest(ctx context.Context, buf *bytes.Buffer) (*http.Request, error) {
	req, errReq := http.NewRequestWithContext(ctx, http.MethodPost, bc.config.URL, buf)
	if errReq != nil {
		return nil, errReq
	}
	req.Header.Add("Content-Type", "application/x-ndjson")

	params := url.Values{}
	if bc.config.Pipeline != "" {
		params.Set("pipeline", bc.config.Pipeline)
	}
	if bc.config.Refresh != "" {
		params.Set("refresh", bc.config.Refresh)
	}
	if bc.config.Routing != "" {
		params.Set("routing", bc.config.Routing)
	}
	// if v := bc.config.Pretty; v != nil {
	// 	params.Set("pretty", fmt.Sprint(*v))
	// }
	// if v := bc.config.Human; v != nil {
	// 	params.Set("human", fmt.Sprint(*v))
	// }
	if v := bc.config.ErrorTrace; v != nil {
		params.Set("error_trace", fmt.Sprint(*v))
	}
	if len(bc.config.FilterPath) > 0 {
		params.Set("filter_path", strings.Join(bc.config.FilterPath, ","))
	}
	if bc.config.Timeout != "" {
		params.Set("timeout", bc.config.Timeout)
	}
	if bc.config.WaitForActiveShards != "" {
		params.Set("wait_for_active_shards", bc.config.WaitForActiveShards)
	}
	return req, nil
}
