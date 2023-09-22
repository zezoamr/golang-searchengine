package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearch(t *testing.T) {

	// Test case: Invalid resultsCount parameter
	t.Run("Invalid resultsCount parameter", func(t *testing.T) {
		// Create request with invalid resultsCount parameter
		req, err := http.NewRequest("GET", "/search/?resultsCount=invalid", nil)
		if err != nil {
			t.Fatal(err)
		}
		rec := httptest.NewRecorder()

		// Call search function
		handleSearch(rec, req)

		// Check response status code
		if rec.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rec.Code)
		}
	})

}
