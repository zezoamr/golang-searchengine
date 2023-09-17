package parsing

import (
	"reflect"
	"testing"
)

// TestParsePage is a unit test for the parsePage function.
//
// It tests the functionality of the parsePage function by providing a sample HTML body
// and comparing the extracted links and words with the expected values.
// The function takes a testing.T object as a parameter, which allows it to report any
// unexpected errors during the test execution.
// The function does not return any values.
func TestParsePage(t *testing.T) {
	body := "<html><body><a href=\"link1\">Link 1</a><a href=\"link2\">Link 2</a><p>Paragraph 1!</p><p>this is Paragraph 2</p></body></html>"
	expectedLinks := []string{"http://example.com/link1", "http://example.com/link2"}
	expectedWords := "Paragraph 1 Paragraph 2"

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
