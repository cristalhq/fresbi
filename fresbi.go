package fresbi

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
