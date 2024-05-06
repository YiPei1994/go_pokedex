package main

import (
	"testing"
)

func TestClearInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{
				"hello", "world",
			},
		},
		{
			input: "I am new To Go world",
			expected: []string{
				"i", "am", "new", "to", "go", "world",
			},
		},
	}

	for _, test := range cases {
		got := cleanInput(test.input)

		if len(got) != len(test.expected) {
			t.Errorf("clearInput(%q) = %q, want %q", test.input, got, test.expected)
			continue
		}

		for i := range got {
			actualWord := got[i]
			expectedWord := test.expected[i]

			if actualWord != expectedWord {
				t.Errorf("%v is not same as %v", actualWord, expectedWord)
			}
		}
	}
}
