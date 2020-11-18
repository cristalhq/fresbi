package fresbi

import (
	"net/http"
)

// Item ...
type Item struct {
	action string

	Index         string `json:"_index,omitempty"`
	ID            string `json:"_id,omitempty"`
	Type          string `json:"_type,omitempty"`
	Parent        string `json:"parent,omitempty"`
	Routing       string `json:"routing,omitempty"`
	Version       *int64 `json:"version,omitempty"`
	VersionType   string `json:"version_type,omitempty"`
	IfSeqNo       *int64 `json:"if_seq_no,omitempty"`
	IfPrimaryTerm *int64 `json:"if_primary_term,omitempty"`
	Pipeline      string `json:"pipeline,omitempty"` // 'index' only

	RetryOnConflict *int `json:"retry_on_conflict,omitempty"` // 'index' and 'update' only

	AsSource bool `json:"-"` // 'update' only

	Body interface{} `json:"-"`
}

// Response represents the Elasticsearch bulk response.
//
type Response struct {
	HTTPResponse *http.Response

	Took      int                        `json:"took,omitempty"`
	HasErrors bool                       `json:"errors,omitempty"`
	Items     []map[string]*ResponseItem `json:"items,omitempty"`
}

// ResponseItem represents the Elasticsearch bulk response item.
//
type ResponseItem struct {
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

// type bulkIndexerItem struct {
// 	Index      string
// 	Action     string
// 	DocumentID string
// 	Body       interface{}
// 	// RetryOnConflict *int
// }
