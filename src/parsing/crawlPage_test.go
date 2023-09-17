package parsing

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestCrawlPage(t *testing.T) {
	parsedChannel := make(chan Page)

	// Test case 1: Crawling multiple valid URLs
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		html := "<html><body><a href=\"link1\">Link 1</a><a href=\"link2\">Link 2</a><p>Paragraph 1</p><p>Paragraph 2</p></body></html>"
		w.Write([]byte(html))
	}))
	defer server.Close()

	url := server.URL
	go crawlPage(url, parsedChannel)

	PAGE := <-parsedChannel
	expectedLinks := []string{server.URL + "/link1", server.URL + "/link2"}
	expectedWords := "Paragraph 1 Paragraph 2"

	if PAGE.Err != nil {
		t.Errorf("unexpected error: %v", PAGE.Err)
	}
	if !reflect.DeepEqual(PAGE.Links, expectedLinks) {
		t.Errorf("expected links: %v, got: %v", expectedLinks, PAGE.Links)
	}
	if !reflect.DeepEqual(PAGE.Words, expectedWords) {
		t.Errorf("expected words: %v, got: %v", expectedWords, PAGE.Words)
	}

	// Test case 2: Crawling URLs with errors
	server2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Not Found")
	}))
	defer server2.Close()

	url = server2.URL
	go crawlPage(url, parsedChannel)
	PAGE = <-parsedChannel
	if PAGE.Err == nil {
		t.Errorf("expected error 404 StatusNotFound instead got: %v", PAGE.Err)
		t.Errorf("page content was %s and its url was %s", PAGE.Words, PAGE.Url)
		t.Errorf("and server url is %s ", server2.URL)
	}

	// Test case 3: Crawling an empty URL
	url = ""
	go crawlPage(url, parsedChannel)
	PAGE = <-parsedChannel
	if PAGE.Err == nil {
		t.Errorf("unexpected error: %v", PAGE.Err)
	}
	fmt.Printf("Error: %v", PAGE.Err)

	close(parsedChannel)
}
