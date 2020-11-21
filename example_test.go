package fresbi_test

import (
	"context"
	"net/http"
	"strconv"

	"github.com/cristalhq/fresbi"
)

func Example() {
	client := fresbi.NewClient("http://localhost:9200", http.DefaultClient, fresbi.Config{})

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
