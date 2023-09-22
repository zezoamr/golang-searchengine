package searchengine

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func search(client *elasticsearch.Client, index string, searchword string, count int) (*esapi.Response, error) {
	var buf strings.Builder
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"words": map[string]interface{}{
					"query":     strings.ToLower(searchword),
					"fuzziness": "AUTO",
				},
			},
		},
		"size": count,
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Printf("Error encoding query: %s", err)
		return nil, err
	}

	res, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex(index),
		client.Search.WithBody(bytes.NewReader([]byte(buf.String()))),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithPretty(),
	)
	if err != nil {
		log.Printf("Error getting response: %s", err)
		return nil, err
	}
	return res, err
}
