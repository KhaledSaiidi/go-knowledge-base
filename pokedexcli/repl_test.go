package main

import "testing"

const (
	ErrorLengthMismatch = "length mismatch: expected %d, got %d"
	ErrorWordMismatch   = "word mismatch: expected '%s', got '%s'"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "Hello World",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "HELLO WORLD",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, cs := range cases {
		actual := CleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf(ErrorLengthMismatch,
				len(cs.expected),
				len(actual))
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]
			if actualWord != expectedWord {
				t.Errorf(ErrorWordMismatch, expectedWord, actualWord)
			}
		}

	}
}
