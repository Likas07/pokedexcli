package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "trim spaces",
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "multiple spaces",
			input:    "   multiple   spaces  between ",
			expected: []string{"multiple", "spaces", "between"},
		},
		{
			name:     "tabs and uppercase",
			input:    "HELLO\tWORLD",
			expected: []string{"hello", "world"},
		},
		{
			name:     "newlines",
			input:    "Line\nBreak",
			expected: []string{"line", "break"},
		},
		{
			name:     "empty",
			input:    "",
			expected: []string{},
		},
		{
			name:     "punctuation preserved",
			input:    "Punctuations, can't",
			expected: []string{"punctuations,", "can't"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := cleanInput(c.input)
			if !reflect.DeepEqual(got, c.expected) {
				t.Fatalf("cleanInput(%q) = %#v, want %#v", c.input, got, c.expected)
			}
		})
	}
}