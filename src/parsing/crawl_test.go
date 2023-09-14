package parsing

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// TestCrawl tests the crawl function.
//
// It tests crawling multiple valid URLs and an empty URL list.
// It also tests crawling URLs with errors.
// The function takes no parameters and returns the links, words,
// and an error.
func TestCrawl(t *testing.T) {

	// Test case 1: Crawling multiple valid URLs
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		html := "<html><body><a href=\"link1\">Link 1</a><a href=\"link2\">Link 2</a><p>Paragraph 1</p><p>Paragraph 2</p></body></html>"
		w.Write([]byte(html))
	}))
	defer server.Close()
	urls := []string{server.URL}
	links, words, err := crawl(urls)
	expectedLinks := []string{server.URL + "/link1", server.URL + "/link2"}
	expectedWords := []string{"Paragraph 1", "Paragraph 2"}
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(links, expectedLinks) {
		t.Errorf("expected links: %v, got: %v", expectedLinks, links)
	}
	if !reflect.DeepEqual(words, expectedWords) {
		t.Errorf("expected words: %v, got: %v", expectedWords, words)
	}

	// Test case 2: Crawling an empty URL list
	urls = []string{}
	links, words, err = crawl(urls)
	expectedLinks = []string{}
	expectedWords = []string{}
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(links) != len(expectedLinks) {
		t.Errorf("expected links: %v, got: %v", expectedLinks, links)
	}
	if len(links) != len(expectedLinks) {
		t.Errorf("expected words: %v, got: %v", expectedWords, words)
	}

	// Test case 3: Crawling URLs with errors
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()
	urls = []string{server.URL}
	_, _, err = crawl(urls)
	if err == nil {
		t.Errorf("expected error instead got: %v", err)
	}
}
