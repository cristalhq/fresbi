package fresbi_test

import (
	"context"
	"net/http"
	"os"
	"strconv"

	"github.com/cristalhq/fresbi"
)

func Example() {
	url := os.Getenv("ELASTICSEARCH_URL")

	client := fresbi.NewClient(url, http.DefaultClient, fresbi.Config{})

	msgs := []string{"hi", "there", "everyone"}

	resp, err := client.AsBatch(context.Background(), func(req fresbi.Batch) error {
		for i, msg := range msgs {

			req.Create(&fresbi.Item{
				Index: "best_index_ever",
				ID:    strconv.Itoa(i),
				Body:  `can be almost anything ` + msg,
			})
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	_ = resp
}
