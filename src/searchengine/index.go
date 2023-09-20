package searchengine

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/zezoamr/golang-searchengine/parsing"
)

func createDocuments(typedClient *elasticsearch.TypedClient, index string, pages []parsing.Page) (*esapi.Response, error) {

	var buf bytes.Buffer
	for _, page := range pages {
		meta := []byte(fmt.Sprintf(`{ "index" : { "_id" : "%s", "_index" : "%s" } }%s`, page.Url, index, "\n"))

		data, err := json.Marshal(page)
		if err != nil {
			log.Fatalf("Cannot encode article %s: %s", page.Url, err)
		}

		data = append(data, "\n"...)

		buf.Grow(len(meta) + len(data))
		buf.Write(meta)
		buf.Write(data)
	}

	req := esapi.BulkRequest{
		Body: &buf,
	}

	res, err := req.Do(context.Background(), typedClient)
	if err != nil {
		log.Printf("Failure indexing batch: %s", err)
		return res, err
	} else if res.IsError() {
		log.Printf("Error indexing document")
		return res, err
	}
	if err == nil {
		log.Printf("Successfully indexed %d documents", len(pages))
	}
	return res, nil
}
