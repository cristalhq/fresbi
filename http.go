package fresbi

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

func send(ctx context.Context, httpClient *http.Client, url string, buf *bytes.Buffer) (*BulkResponse, error) {
	req, errReq := makeRequest(ctx, url, buf)
	if errReq != nil {
		return nil, errReq
	}

	resp, errResp := httpClient.Do(req)
	if errResp != nil {
		return nil, errResp
	}
	defer func() { _ = resp.Body.Close() }()

	bulkResp := &BulkResponse{}
	if err := json.NewDecoder(resp.Body).Decode(bulkResp); err != nil {
		return nil, err
	}

	bulkResp.HTTPResponse = resp

	return bulkResp, nil
}

func makeRequest(ctx context.Context, url string, buf *bytes.Buffer) (*http.Request, error) {
	req, errReq := http.NewRequestWithContext(ctx, http.MethodPost, url, buf)
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
