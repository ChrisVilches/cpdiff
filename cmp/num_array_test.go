package cmp

import (
	"math/big"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestHasRealNumbers(t *testing.T) {
	data := []string{"1 2 3 4 5 6", "1 2 3 4.005 5 6"}
	ans := []bool{false, true}

	for i, testCase := range data {
		input := testCase
		expected := ans[i]

		s := newComparable(input).(NumArray)
		res := s.HasRealNumbers()

		if res != expected {
			t.Fatalf("expected %s to have real numbers (%t)", input, expected)
		}
	}
}

func TestNumArrayShortDisplay(t *testing.T) {
	data := []string{"1 2 3 4 5 6", "1 2 3 4.005 5 6  0 ", " 5  ", " 1 5 ", " 5.5551"}
	ans := []string{"(6 numbers...)", "(7 numbers...)", " 5  ", "(2 numbers...)", " 5.5551"}

	for i, testCase := range data {
		input := testCase
		expected := ans[i]

		s := newComparable(input).(NumArray)
		res := s.ShortDisplay()

		if res != expected {
			t.Fatalf("expected '%s' to be '%s'", input, expected)
		}
	}
}

const defaultError = 0.0001

func TestCompareNums(t *testing.T) {
	data := [][]string{
		{"1 2 3", "1 2 3"},
		{"1 5 5 5 5", "1 5 5 5 5 9"},
		{"1 2 3 5 9 8", "0 1 2 5 9 8"},
		{"1 ", "4 5 6 7 8 9"},
		{"1", "1"},
		{"1", "1.000001"},
		{"1", "1.00009"},
		{"1", "1.001"},
	}

	err := big.NewFloat(defaultError)

	ans := [][]cmpRange{
		{{From: 0, To: 3, Result: CmpRes.Correct}},
		{{From: 0, To: 5, Result: CmpRes.Correct}, {From: 5, To: 6, Result: CmpRes.Incorrect}},
		{{From: 0, To: 3, Result: CmpRes.Incorrect}, {From: 3, To: 6, Result: CmpRes.Correct}},
		{{From: 0, To: 6, Result: CmpRes.Incorrect}},
		{{From: 0, To: 1, Result: CmpRes.Correct}},
		{{From: 0, To: 1, Result: CmpRes.Approx}},
		{{From: 0, To: 1, Result: CmpRes.Approx}},
		{{From: 0, To: 1, Result: CmpRes.Incorrect}},
	}

	for i, testCase := range data {
		s1 := newComparable(testCase[0]).(NumArray)
		s2 := newComparable(testCase[1]).(NumArray)
		expected := ans[i]

		res1, _ := compareNums(s1, s2, err, false)
		res2, _ := compareNums(s2, s1, err, false)

		if !cmp.Equal(res1, expected) || !cmp.Equal(res2, expected) {
			t.Fatalf("expected ranges to be equal when comparing %s and %s", testCase[0], testCase[1])
		}
	}
}

func strToBigFloat(s string) *big.Float {
	val, _ := new(big.Float).SetString(s)

	return val
}

func TestCompareNumsMaxErr(t *testing.T) {
	data := [][]string{
		{"1 2 3", "1 2 3"},
		{"1.5", "1.50006"},
		{"1.50005", "1.50006"},
		{"0", "0"},
		{"0", "0.00000001"},
	}

	err := big.NewFloat(defaultError)

	ans := []*big.Float{
		strToBigFloat("0"),
		strToBigFloat("0.00006"),
		strToBigFloat("0.00001"),
		strToBigFloat("0"),
		strToBigFloat("0.00000001"),
	}

	for i, testCase := range data {
		s1 := newComparable(testCase[0]).(NumArray)
		s2 := newComparable(testCase[1]).(NumArray)
		expected := ans[i]

		_, res := compareNums(s1, s2, err, false)

		if res.String() != expected.String() {
			t.Fatalf("expected error to be %s but got %s", expected.String(), res.String())
		}
	}
}
