package parsing

import (
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty URL",
			input:    "",
			expected: "",
		},
		{
			name:     "URL with lowercase scheme and host",
			input:    "https://example.com",
			expected: "https://example.com",
		},
		{
			name:     "URL with uppercase scheme and host",
			input:    "HTTP://EXAMPLE.COM",
			expected: "http://example.com",
		},
		{
			name:     "URL with fragment",
			input:    "https://example.com#fragment",
			expected: "https://example.com",
		},
		{
			name:     "URL with query parameters",
			input:    "https://example.com/?b=2&a=1",
			expected: "https://example.com/?a=1&b=2",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if actual != tc.expected {
				t.Errorf("Expected normalized URL %s, but got %s", tc.expected, actual)
			}
		})
	}
}
