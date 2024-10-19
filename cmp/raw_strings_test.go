package cmp

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRawStringShortDisplay(t *testing.T) {
	data := [][]string{
		{"helloworldhahaha", "hellowo..."},
		{"helloworld", "helloworld"},
		{"hello", "hello"},
		{"    hello   ", "    hel..."},
	}

	for _, testCase := range data {
		input := testCase[0]
		expected := testCase[1]

		s := RawString{value: input}
		res := s.ShortDisplay()

		if res != expected {
			t.Fatalf("expected %s to be %s", res, expected)
		}
	}
}

func TestCompareStrings(t *testing.T) {
	data := [][]string{
		{"AABBCC", "AAXYCC"},
		{"AAABBBCCC", "AAABBBCCC"},
		{"CCCYYYAAA", "AAABBBCCC"},
		{"CCCBBBAAA", "AAABBBCCC"},
		{"AAAA", "AAAAAAAA"},
		{"B", "AAAAAAAAA"},
	}

	ans := [][]cmpRange{
		{{From: 0, To: 2, Result: CmpRes.Correct}, {From: 2, To: 4, Result: CmpRes.Incorrect}, {From: 4, To: 6, Result: CmpRes.Correct}},
		{{From: 0, To: 9, Result: CmpRes.Correct}},
		{{From: 0, To: 9, Result: CmpRes.Incorrect}},
		{{From: 0, To: 3, Result: CmpRes.Incorrect}, {From: 3, To: 6, Result: CmpRes.Correct}, {From: 6, To: 9, Result: CmpRes.Incorrect}},
		{{From: 0, To: 4, Result: CmpRes.Correct}, {From: 4, To: 8, Result: CmpRes.Incorrect}},
		{{From: 0, To: 9, Result: CmpRes.Incorrect}},
	}

	for i, testCase := range data {
		expected := ans[i]

		s1 := RawString{value: testCase[0]}
		s2 := RawString{value: testCase[1]}

		res1 := compareStrings(s1, s2)
		res2 := compareStrings(s2, s1)

		if !cmp.Equal(res1, expected) || !cmp.Equal(res2, expected) {
			t.Fatalf("expected ranges to be equal when comparing %s and %s", testCase[0], testCase[1])
		}
	}
}
