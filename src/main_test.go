package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHtml(t *testing.T) {
	// Create a test server to mock the HTTP response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<html><body><h1>Hello, World!</h1></body></html>"))
	}))

	// Test case 1: Test a valid URL
	url := server.URL
	expected := "<html><body><h1>Hello, World!</h1></body></html>"
	result, err := getHtml(url)
	if err != nil {
		t.Errorf("Expected: %s, but got: %s", expected, err)
	}
	if result != expected {
		t.Errorf("Expected: %s, but got: %s", expected, result)
	}

	// Test case 2: Test an invalid URL
	url = "invalid-url"
	expected = ""
	_, err = getHtml(url)
	if err == nil {
		t.Errorf("Expected: %s, but got: %s", "an error", err)
	}

	server.Close()

	// Test case 3: Test a URL that returns an error
	// Note: This test case assumes that the server is not running or unreachable
	url = "http://localhost:1234"
	expected = ""
	_, err = getHtml(url)
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
	result, err = getHtml(url)
	if err != nil {
		t.Errorf("Expected: %s, but got: %s", expected, err)
	}
	if result != expected {
		t.Errorf("Expected: %s, but got: %s", expected, result)
	}
}
