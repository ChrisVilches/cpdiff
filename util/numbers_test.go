package util

import (
	"math/big"
	"testing"
)

func TestBigMax(t *testing.T) {
	data := [][]float64{
		{1, 2, 2},
		{7.55, 4.1212, 7.55},
		{5, 0.0000, 5},
		{-5, 0.0000, 0},
	}

	for _, testCase := range data {
		a := big.NewFloat(testCase[0])
		b := big.NewFloat(testCase[1])
		c := big.NewFloat(testCase[2])

		if BigMax(a, b).Cmp(c) != 0 {
			t.Fatalf("BigMax(%s, %s) expected to be %s", a.String(), b.String(), c.String())
		}
	}
}

func strToBigFloat(s string) *big.Float {
	val, _ := new(big.Float).SetString(s)

	return val
}

func TestAbsError(t *testing.T) {
	data := [][]string{
		{"5", "5.55", "0.55"},
		{"5", "9.55", "4.55"},
		{"-1", "1.5", "2.5"},
		{"1.5", "1.5", "0"},
		{"1.41421356237", "1.41", "0.00421356237"},
	}

	for _, testCase := range data {
		a := strToBigFloat(testCase[0])
		b := strToBigFloat(testCase[1])
		c := strToBigFloat(testCase[2])
		res := AbsError(a, b)

		if res.String() != c.String() {
			t.Fatalf("expected error to be %s but got %s", c.String(), res.String())
		}
	}
}

func TestRelError(t *testing.T) {
	data := [][]string{
		{"1.41421356237", "1.41", "0.002988342106"},
		{"5", "5.55", "0.0990990991"},
		{"5", "9.55", "0.4764397906"},
		{"-1", "1.5", "1.666666667"},
		{"1.5", "1.5", "0"},
		{"0", "1", "1"},
		{"1", "0", "+Inf"},
		{"0", "0", "+Inf"},
	}

	for _, testCase := range data {
		a := strToBigFloat(testCase[0])
		b := strToBigFloat(testCase[1])
		c := strToBigFloat(testCase[2])
		res := RelError(a, b)

		if res.String() != c.String() {
			t.Fatalf("expected error to be %s but got %s", c.String(), res.String())
		}
	}
}
