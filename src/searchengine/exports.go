package searchengine

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/zezoamr/golang-searchengine/parsing"
)

func GetClient() (*elasticsearch.TypedClient, *elasticsearch.Client, error) {
	return getClient()
}

func CreateIndex(typedClient *elasticsearch.TypedClient, index string) error {
	return createIndex(typedClient, index)
}

func DeleteIndex(typedClient *elasticsearch.TypedClient, index string) error {
	return deleteIndex(typedClient, index)
}

func Search(typedClient *elasticsearch.Client, index string, searchword string, count int) (*esapi.Response, error) {
	return search(typedClient, index, searchword, count)
}

func CreateDocuments(typedClient *elasticsearch.TypedClient, index string, pages []parsing.Page) (*esapi.Response, error) {
	return createDocuments(typedClient, index, pages)
}
