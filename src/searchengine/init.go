package searchengine

import (
	"context"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/joho/godotenv"
)

// xpack.security.http.ssl:
//   enabled: false
//   keystore.path: certs/http.p12

func getClient() (*elasticsearch.TypedClient, *elasticsearch.Client, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, nil, err
	}

	cfg := elasticsearch.Config{
		Addresses: []string{
			os.Getenv("ELASTIC_URL"),
		},
		Username: os.Getenv("ELASTIC_USERNAME"),
		Password: os.Getenv("ELASTIC_PASSWORD"),
	}

	typedClient, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		log.Printf("Error creating elasticsearch client: %s", err)
		return nil, nil, err
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Printf("Error creating elasticsearch client: %s", err)
		return nil, nil, err
	}

	log.Printf("successfully created elasticsearch client")

	return typedClient, client, err
}

func createIndex(typedClient *elasticsearch.TypedClient, index string) error {

	_, err := typedClient.Indices.Create(index).Do(context.TODO())
	if err != nil {
		log.Printf("Failure indexing batch: %s", err)
		return err
	}
	log.Printf("successfully created index: %s", index)
	return nil
}

func deleteIndex(typedClient *elasticsearch.TypedClient, index string) {
	typedClient.Indices.Delete(index).Do(context.TODO())
	log.Printf("successfully deleted index: %s", index)
}
