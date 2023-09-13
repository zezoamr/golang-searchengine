package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetPage(t *testing.T) {
	// Create a test server to mock the HTTP response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<html><body><h1>Hello, World!</h1></body></html>"))
	}))

	// Test case 1: Test a valid URL
	url := server.URL
	expected := "<html><body><h1>Hello, World!</h1></body></html>"
	result, err := getPage(url)
	if err != nil {
		t.Errorf("Expected: %s, but got: %s", expected, err)
	}
	if result != expected {
		t.Errorf("Expected: %s, but got: %s", expected, result)
	}

	// Test case 2: Test an invalid URL
	url = "invalid-url"
	expected = ""
	_, err = getPage(url)
	if err == nil {
		t.Errorf("Expected: %s, but got: %s", "an error", err)
	}

	server.Close()

	// Test case 3: Test a URL that returns an error
	// Note: This test case assumes that the server is not running or unreachable
	url = "http://localhost:1234"
	expected = ""
	_, err = getPage(url)
	if err == nil {
		t.Errorf("Expected: %s, but got: %s", "an error", err)
	}

	// Test case 4: Test a URL that returns empty body
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	url = server.URL
	expected = ""
	result, err = getPage(url)
	if err != nil {
		t.Errorf("Expected: %s, but got: %s", expected, err)
	}
	if result != expected {
		t.Errorf("Expected: %s, but got: %s", expected, result)
	}
}

func TestParsePage(t *testing.T) {
	body := "<html><body><a href=\"link1\">Link 1</a><a href=\"link2\">Link 2</a><p>Paragraph 1</p><p>Paragraph 2</p></body></html>"
	expectedLinks := []string{"link1", "link2"}
	expectedWords := []string{"Paragraph 1", "Paragraph 2"}

	links, words, err := parsePage(body)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(links, expectedLinks) {
		t.Errorf("expected links: %v, got: %v", expectedLinks, links)
	}

	if !reflect.DeepEqual(words, expectedWords) {
		t.Errorf("expected words: %v, got: %v", expectedWords, words)
	}
}

// func TestCrawl(t *testing.T) {
// 	// Test case 1: Crawling multiple valid URLs
// 	urls := []string{"https://google.com"}
// 	crawl(urls)

// 	// Test case 2: Crawling an empty URL list
// 	urls = []string{}
// 	crawl(urls)

// 	// Test case 3: Crawling URLs with errors
// 	urls = []string{"https://invalidurl"}
// 	crawl(urls)
// }
// generate an endpoint for those tests
