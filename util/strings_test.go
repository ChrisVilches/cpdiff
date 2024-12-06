package util

import (
	"testing"
)

func TestRemoveTrailingNewLine(t *testing.T) {
	data := [][]string{
		{"hello ", "hello "},
		{"hello", "hello"},
		{"hello\n", "hello"},
		{"hello\n\n", "hello\n"},
	}

	for _, testCase := range data {
		in := testCase[0]
		expected := testCase[1]

		if RemoveTrailingNewLine(in) != expected {
			t.Fatalf("%s expected to be %s", in, expected)
		}
	}
}
