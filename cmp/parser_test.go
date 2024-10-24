package cmp

import (
	"testing"
)

func TestParseRawString(t *testing.T) {
	data := []string{
		"hello",
		"   ",
		" ",
		" a ",
		" . ",
		"   hello    ",
		"0011101010101010101010101011000010100010100101010100000000010000000000001111111111111111111",
		"  15654.a ",
		"  15654a ",
		"    00000 ",
		"    1e--9 ",
		"    00001 ",
		"  1111  01 ",
		"    00 ",
		"    0.. ",
		"    ..0 ",
		"    0..3111 ",
		"    0..3111  456",
		"    00.3111  456 ",
	}

	for _, testCase := range data {
		s := newComparable(testCase)

		if _, ok := s.(RawString); !ok {
			t.Fatalf("expected %s to be parsed as RawString", testCase)
		}
	}
}

func TestParseNumArray(t *testing.T) {
	data := []string{
		"10011101010101010101010101011000010100010100101010100000000010000000000001111111111111111111",
		"   10011101010101010101010101011000010100010100101010100000000010000000000001111111111111111111.00025 ",
		"    0.0000 ",
		"  1111  10 ",
		"    0.0001 ",
		"    -00001 ",
		"    -1 ",
		"-0011101010101010101010101011",
		"    -.1 ",
		"    1e-9 ",
		"    0. ",
		"    .0 ",
		"    .0 .1 .0 ",
		"    .0 .1 .0 -00001",
		"    0 ",
		" 0 1 2 3 4 5 6 7 8 9 ",
		"    0.3111  456",
		"  0.3111  456456456454654654654654654654654654.456456454979876546861615645649879846546546578979876 ",
	}

	for _, testCase := range data {
		s := newComparable(testCase)

		if _, ok := s.(NumArray); !ok {
			t.Fatalf("expected %s to be parsed as NumArray", testCase)
		}
	}
}

func TestParseEmpty(t *testing.T) {
	data := []string{""}

	for _, testCase := range data {
		s := newComparable(testCase)

		if _, ok := s.(Empty); !ok {
			t.Fatalf("expected %s to be parsed as Empty", testCase)
		}
	}
}
