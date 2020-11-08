package fresbi

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// Client ...
// See: https://www.elastic.co/guide/en/elasticsearch/reference/master/docs-bulk.html
//
type Client struct {
	url    string
	buf    *bytes.Buffer
	aux    []byte
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

// Action represents ElasticSearch bulk operation.
type Action string

// Allowed bulk operations.
const (
	IndexAction  Action = "index"
	CreateAction Action = "create"
	UpdateAction Action = "delete"
	DeleteAction Action = "index"
)

// NewClient ...
func NewClient(client *http.Client) *Client {
	return &Client{
		client: client,
		buf:    &bytes.Buffer{},
	}
}

// Close ...
func (c *Client) Close() error {
	return nil
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
	if err := fn(c); err != nil {
		c.buf.Reset()
		return err
	}
	if err := c.flushBuffer(ctx); err != nil {
		return err
	}
	return nil
}

// Index ...
func (c *Client) Index(docID string, data interface{}) error {
	return c.addItem(&bulkIndexerItem{
		Action:     IndexAction,
		DocumentID: docID,
		Body:       data,
	})
	// return c.add(IndexAction, docID, data)
}

// Create ...
func (c *Client) Create(docID string, data interface{}) error {
	return c.addItem(&bulkIndexerItem{
		Action:     CreateAction,
		DocumentID: docID,
		Body:       data,
	})
	// return c.add(CreateAction, docID, data)
}

// Update ...
func (c *Client) Update(docID string, data interface{}) error {
	return c.addItem(&bulkIndexerItem{
		Action:     UpdateAction,
		DocumentID: docID,
		Body:       data,
	})
	// return c.add(UpdateAction, docID, data)
}

// Delete ...
func (c *Client) Delete(docID string) error {
	return c.addItem(&bulkIndexerItem{
		Action:     DeleteAction,
		DocumentID: docID,
	})
	// return c.add(DeleteAction, docID, nil)
}

func (c *Client) addItem(item *bulkIndexerItem) error {
	c.writeMeta(item)
	return c.writeBody(item)
}

// // Add ...
// func (c *Client) add(action Action, docID string, data interface{}) error {
// 	if err := c.addToBatch(action, docID, data); err != nil {
// 		return err
// 	}
// 	// if c.buf.Len() >= c.config.batchSize {
// 	// 	return c.flushBuffer(context.Background())
// 	// }
// 	return nil
// }

// func (c *Client) addToBatch(action Action, docID string, data interface{}) error {
// 	var body []byte

// 	switch data := data.(type) {
// 	case []byte:
// 		body = data
// 	case string:
// 		body = []byte(data)
// 	default:
// 		var errJSON error
// 		body, errJSON = json.Marshal(data)
// 		if errJSON != nil {
// 			return errJSON
// 		}
// 	}

// 	meta := []byte(fmt.Sprintf(`{ "index" : { "_id" : "%s" } }\n`, docID))

// 	c.buf.Grow(len(meta) + len(body) + 1)
// 	c.buf.Write(meta)
// 	c.buf.Write(body)
// 	c.buf.Write([]byte("\n"))
// 	return nil
// }

type bulkIndexerItem struct {
	Index      string
	Action     Action
	DocumentID string
	Body       interface{}
	// RetryOnConflict *int
}

func (c *Client) writeMeta(item *bulkIndexerItem) {
	c.buf.WriteByte('{')
	c.aux = strconv.AppendQuote(c.aux, string(item.Action))
	c.buf.Write(c.aux)
	c.aux = c.aux[:0]

	c.buf.WriteString(":{")
	// c.buf.WriteByte(':')
	// c.buf.WriteByte('{')

	if item.DocumentID != "" {
		c.buf.WriteString(`"_id":`)
		c.aux = strconv.AppendQuote(c.aux, item.DocumentID)
		c.buf.Write(c.aux)
		c.aux = c.aux[:0]
	}

	if item.Index != "" {
		if item.DocumentID != "" {
			c.buf.WriteByte(',')
		}
		c.buf.WriteString(`"_index":`)
		c.aux = strconv.AppendQuote(c.aux, item.Index)
		c.buf.Write(c.aux)
		c.aux = c.aux[:0]
	}

	c.buf.WriteString("}}\n")
	// c.buf.WriteByte('}')
	// c.buf.WriteByte('}')
	// c.buf.WriteByte('\n')
}

func (c *Client) writeBody(item *bulkIndexerItem) error {
	var body []byte

	switch data := item.Body.(type) {
	case []byte:
		body = data
	case string:
		body = []byte(data)

	case io.Reader:
		_, err := c.buf.ReadFrom(data)
		if err != nil {
			return err
		}
	default:
		err := json.NewEncoder(c.buf).Encode(data)
		if err != nil {
			return err
		}
	}

	// true only when item.Body is []byte or string
	if body != nil {
		_, _ = c.buf.Write(body)
	}
	c.buf.WriteByte('\n')
	return nil
}

// flushBuffer ...
func (c *Client) flushBuffer(ctx context.Context) error {
	req := c.makeRequest(ctx)

	resp, errResp := c.client.Do(req)
	if errResp != nil {
		return errResp
	}
	defer func() { _ = resp.Body.Close() }()

	return nil
}

func (c *Client) makeRequest(ctx context.Context) *http.Request {
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

	req, errReq := http.NewRequestWithContext(ctx, http.MethodPost, c.url, c.buf)
	if errReq != nil {
		panic(errReq)
	}
	req.Header.Add("Content-Type", "application/x-ndjson")

	return req
}
