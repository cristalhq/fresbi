package fresbi_test

import (
	"context"
	"net/http"
	"strconv"

	"github.com/cristalhq/fresbi"
)

func Example() {
	client := fresbi.NewClient(http.DefaultClient)

	msgs := []string{"hi", "there", "everyone"}

	_, err := client.AsBatch(context.Background(), func(req fresbi.Batch) error {
		for i, msg := range msgs {
			docID := strconv.Itoa(i)

			req.Create(docID, msg)
		}
		return nil
	})
	_ = err

	// Output:
}
