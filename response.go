package fresbi

import "net/http"

// BulkResponse represents the Elasticsearch bulk response.
//
type BulkResponse struct {
	HTTPResponse http.Response

	Took      int                            `json:"took,omitempty"`
	HasErrors bool                           `json:"errors,omitempty"`
	Items     []map[string]*BulkResponseItem `json:"items,omitempty"`
}

// BulkResponseItem represents the Elasticsearch bulk response item.
//
type BulkResponseItem struct {
	Index      string `json:"_index"`
	DocumentID string `json:"_id"`
	Version    int64  `json:"_version"`
	Result     string `json:"result"`

	Shards struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Failed     int `json:"failed"`
	} `json:"_shards"`

	SeqNo       int64 `json:"_seq_no"`
	PrimaryTerm int64 `json:"_primary_term"`
	Status      int   `json:"status"`

	Error struct {
		Type      string `json:"type"`
		Reason    string `json:"reason"`
		IndexUUID string `json:"index_uuid"`
		Shard     string `json:"shard"`
		Index     string `json:"index"`
		// Cause     struct {
		// 	Type   string `json:"type"`
		// 	Reason string `json:"reason"`
		// } `json:"caused_by"`
	} `json:"error,omitempty"`

	// Type          string `json:"_type,omitempty"`
	// ID            string `json:"_id,omitempty"`
	// ForcedRefresh bool   `json:"forced_refresh,omitempty"`
}
