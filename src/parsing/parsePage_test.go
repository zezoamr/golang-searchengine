package parsing

import (
	"reflect"
	"testing"
)

// TestParsePage is a unit test for the parsePage function.
//
// It tests the functionality of the parsePage function by providing a sample HTML body and expected results.
// The function checks if the parsed links and words match the expected ones, and if any error is returned.
// It uses the reflect.DeepEqual function to compare the slices of links and words.
// The test fails if the links or words do not match the expected ones, or if an error is returned.
func TestParsePage(t *testing.T) {
	body := "<html><body><a href=\"link1\">Link 1</a><a href=\"link2\">Link 2</a><p>Paragraph 1</p><p>Paragraph 2</p></body></html>"
	expectedLinks := []string{"http://example.com/link1", "http://example.com/link2"}
	expectedWords := []string{"Paragraph 1", "Paragraph 2"}

	links, words, err := parsePage("http://example.com", body)
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
