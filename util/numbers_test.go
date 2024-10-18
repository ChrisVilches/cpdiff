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
