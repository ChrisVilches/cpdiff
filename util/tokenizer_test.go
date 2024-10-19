package util

import (
	"fmt"
	"testing"
)

func TestStringFieldsKeepWhitespace(t *testing.T) {
	data := [][]string{
		{"  hello  world", "[  hello  ][world]"},
		{"  hello  world    ", "[  hello  ][world    ]"},
		{" a b c d e f ", "[ a ][b ][c ][d ][e ][f ]"},
		{"", ""},
		{"   ", ""},
		{"      ", ""},
		{"   a   ", "[   a   ]"},
		{"   a   x", "[   a   ][x]"},
	}

	for _, testCase := range data {
		in := testCase[0]
		expected := testCase[1]
		res := ""
		prev := 0
		prevIdx := -1

		for idx, i := range StringFieldsKeepWhitespace(in) {
			res += fmt.Sprintf("[%s]", in[prev:i])
			prev = i

			if prevIdx+1 != idx {
				t.Fatalf("expected index to be increasing")
			}

			prevIdx = idx
		}

		if res != expected {
			t.Fatalf("%s expected to be tokenized as %s", in, expected)
		}

		count := 0

		for range StringFieldsKeepWhitespace("a b c d e f g h i j k") {
			count++
			if count == 2 {
				break
			}
		}

		if count != 2 {
			t.Fatalf("expected to quit iteration")
		}
	}
}
