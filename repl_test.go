package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, testCase := range testCases {
		actual := cleanInput(testCase.input)
		for i, actualWord := range actual {
			expectedWord := testCase.expected[i]
			if actualWord != expectedWord {
				t.Errorf("expected %q, got %q", expectedWord, actual)
			}
		}
	}
}
