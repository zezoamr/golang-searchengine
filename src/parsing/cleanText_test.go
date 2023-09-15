package parsing

import "testing"

func TestCleanText(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "String with useless words and punctuation",
			input:    "Hello, world! This is a test.",
			expected: "Hello world test",
		},
		{
			name:     "String with only punctuation",
			input:    "!@#%&*()",
			expected: "",
		},
		{
			name:     "String with only useless words",
			input:    "the a an",
			expected: "",
		},
		{
			name:     "String with no useless words or punctuation",
			input:    "Hello world",
			expected: "Hello world",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := cleanText(testCase.input)
			if result != testCase.expected {
				t.Errorf("Expected %q, but got %q", testCase.expected, result)
			}
		})
	}
}

func TestRemoveUselessWords(t *testing.T) {
	text := "This is a test sentence about to remove any uselesswords."
	expected := "test sentence remove uselesswords."
	cleanedText := removeUselessWords(text)
	if cleanedText != expected {
		t.Errorf("Expected: %s, but got: %s", expected, cleanedText)
	}
}

func TestRemovePunctuation(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "No punctuation",
			input:    "Hello",
			expected: "Hello",
		},
		{
			name:     "Punctuation at the beginning",
			input:    "!Hello",
			expected: "Hello",
		},
		{
			name:     "Punctuation at the end",
			input:    "Hello!",
			expected: "Hello",
		},
		{
			name:     "Punctuation in the middle",
			input:    "He!llo",
			expected: "Hello",
		},
		{
			name:     "Multiple punctuations",
			input:    "H!e!l!l!o",
			expected: "Hello",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := removePunctuation(test.input)
			if result != test.expected {
				t.Errorf("Expected %s, but got %s", test.expected, result)
			}
		})
	}
}
