package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCrawl(t *testing.T) {

	// Test case 1: Successful decoding of JSON data
	t.Run("DecodeJSONSuccess", func(t *testing.T) {
		// Create a request with valid JSON data
		reqBody := bytes.NewBufferString(`{"links": ["link1", "link2"]}`)
		req, err := http.NewRequest("POST", "/crawl?maxTotal=1&maxPerPage=1&rate=2", reqBody)
		if err != nil {
			t.Fatal(err)
		}

		// Create a response recorder to capture the response
		rr := httptest.NewRecorder()

		// Call the crawl function with the request and response recorder
		handleCrawl(rr, req)

		// Check the response status code
		if rr.Code != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
		}

		// Check the response body
		expectedBody := "crawling"
		if rr.Body.String() != expectedBody {
			t.Errorf("Expected body %q, but got %q", expectedBody, rr.Body.String())
		}
	})

	// Test case 2: Invalid maxTotal parameter
	t.Run("InvalidMaxTotalParam", func(t *testing.T) {
		// Create a request with invalid maxTotal parameter
		reqBody := bytes.NewBufferString(`{"links": ["link1", "link2"]}`)
		req, err := http.NewRequest("POST", "/crawl?maxTotal=invalid&maxPerPage=1&rate=2", reqBody)
		if err != nil {
			t.Fatal(err)
		}

		// Create a response recorder to capture the response
		rr := httptest.NewRecorder()

		// Call the crawl function with the request and response recorder
		handleCrawl(rr, req)

		// Check the response status code
		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, rr.Code)
		}

		// Check the response body
		expectedBody := "strconv.Atoi: parsing \"invalid\": invalid syntax\n"
		if rr.Body.String() != expectedBody {
			t.Errorf("Expected body %q, but got %q", expectedBody, rr.Body.String())
		}
	})

	// Test case 3: Invalid maxPerPage parameter
	t.Run("InvalidMaxPerPageParam", func(t *testing.T) {
		// Create a request with invalid maxPerPage parameter
		reqBody := bytes.NewBufferString(`{"links": ["link1", "link2"]}`)
		req, err := http.NewRequest("POST", "/crawl?maxPerPage=invalid&maxTotal=1&rate=2", reqBody)
		if err != nil {
			t.Fatal(err)
		}

		// Create a response recorder to capture the response
		rr := httptest.NewRecorder()

		// Call the crawl function with the request and response recorder
		handleCrawl(rr, req)

		// Check the response status code
		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, rr.Code)
		}

		// Check the response body
		expectedBody := "strconv.Atoi: parsing \"invalid\": invalid syntax\n"
		if rr.Body.String() != expectedBody {
			t.Errorf("Expected body %q, but got %q", expectedBody, rr.Body.String())
		}
	})

	// Test case 4: Invalid rate parameter
	t.Run("InvalidRateParam", func(t *testing.T) {
		// Create a request with invalid rate parameter
		reqBody := bytes.NewBufferString(`{"links": ["link1", "link2"]}`)
		req, err := http.NewRequest("POST", "/crawl?rate=invalid&maxTotal=1&maxPerPage=1", reqBody)
		if err != nil {
			t.Fatal(err)
		}
		// Create a response recorder to capture the response
		rr := httptest.NewRecorder()

		// Call the crawl function with the request and response recorder
		handleCrawl(rr, req)

		// Check the response status code
		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, rr.Code)
		}

		// Check the response body
		expectedBody := "strconv.Atoi: parsing \"invalid\": invalid syntax\n"
		if rr.Body.String() != expectedBody {
			t.Errorf("Expected body %q, but got %q", expectedBody, rr.Body.String())
		}
	})

}
