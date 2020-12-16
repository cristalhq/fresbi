package fresbi_test

import (
	"context"
	"net/http"
	"os"
	"strconv"

	"github.com/cristalhq/fresbi"
)

func ExampleReliableClient() {
	url := os.Getenv("ELASTICSEARCH_URL") // by example "http://localhost:9200/_bulk"

	client := fresbi.NewReliableClient(http.DefaultClient, fresbi.Config{URL: url})

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

	// Output:
	//
}

func ExampleRawClient() {
	url := os.Getenv("ELASTICSEARCH_URL") // by example "http://localhost:9200/_bulk"

	client := fresbi.NewRawClient(http.DefaultClient, fresbi.Config{URL: url})

	err := client.Create(&fresbi.Item{
		Index: "best_index_ever",
		ID:    strconv.Itoa(1),
		Body:  `can be almost anything `,
	})
	if err != nil {
		panic(err)
	}

	resp, err := client.Send(context.Background())
	if err != nil {
		panic(err)
	}
	_ = resp
}
