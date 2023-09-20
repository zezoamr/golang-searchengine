package searchengine

import (
	"errors"
	"io"
	"testing"
	"time"

	"github.com/zezoamr/golang-searchengine/parsing"
)

func TestCrawlPage(t *testing.T) {
	typedClient, client, err := getClient()
	if err != nil {
		t.Fatal(err)
	}
	deleteIndex(typedClient, "test")
	// if test index does not exist it won't return an error
	// if it exists it cleans it for the test and thus we are indeed testing it
	// (if this fails and test index exisits below will return an err)
	// by cleaning up when starting and not defering it we are avoiding crashes preventing the cleaning of the index for the subsquent tests
	err = createIndex(typedClient, "test")
	if err != nil {
		t.Fatal(err)
	}

	pages := []parsing.Page{
		{
			Url:   "https://www.google.com",
			Words: "google google google",
			Links: []string{
				"https://www.google.com",
				"https://www.google.com/?q=hello",
			},
			Err: errors.New(""),
		},
		{
			Url:   "https://www.facebook.com",
			Words: "facebook google",
			Links: []string{
				"https://www.facebook.com",
				"https://www.facebook.com/signup",
			},
			Err: errors.New(""),
		},
	}

	resp, err := createDocuments(typedClient, "test", pages)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(string(body))

	time.Sleep(5 * time.Second) //give time for documents to be available to be searched

	resp, err = search(client, "test", "google", 1)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(string(body))

	resp, err = search(client, "test", "google", 3)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(string(body))

	resp, err = search(client, "test", "gogle", 2) //testing fuziness
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(string(body))

}
