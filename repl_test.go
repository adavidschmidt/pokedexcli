package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "	hello	world	",
			expected: []string{"hello", "world"},
		},
		{
			input:    " batman	is	best	",
			expected: []string{"batman", "is", "best"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length of output is not as expected")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Test failed %s does not equal %s", word, expectedWord)
			}
		}
	}
}
