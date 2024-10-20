package big

import (
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
		a := NewFromFloat64(testCase[0])
		b := NewFromFloat64(testCase[1])
		c := NewFromFloat64(testCase[2])

		if !BigDecimalMax(a, b).ExactEq(c) {
			t.Fatalf("BigMax(%s, %s) expected to be %s", a.String(), b.String(), c.String())
		}
	}

	if BigDecimalMax(NewFromFloat64(5), nil).String() != "5" {
		t.Fatalf("expected to return the non-nil value")
	}
	if BigDecimalMax(nil, NewFromFloat64(7)).String() != "7" {
		t.Fatalf("expected to return the non-nil value")
	}
	if BigDecimalMax(nil, nil) != nil {
		t.Fatalf("expected to return nil")
	}
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
		a := NewFromStringUnsafe(testCase[0])
		b := NewFromStringUnsafe(testCase[1])
		c := NewFromStringUnsafe(testCase[2])
		res := absError(a, b)

		// TODO: I think this one is wrong, because it won't be == exactly. Or would it????
		if !res.ExactEq(c) {
			t.Fatalf("expected error to be %s but got %s", c.String(), res.String())
		}
	}
}

func TestRelError(t *testing.T) {
	data := [][]string{
		{"1.41421356237", "1.41", "0.002988342106382979"},
		{"5", "5.55", "0.09909909909909910"},
		{"5", "9.55", "0.4764397905759162"},
		{"-1", "1.5", "1.666666666666667"},
		{"1.5", "1.5", "0"},
		{"0", "1", "1"},
		{"1", "0", "+Inf"},
		{"0", "0", "+Inf"},
	}

	for _, testCase := range data {
		a := NewFromStringUnsafe(testCase[0])
		b := NewFromStringUnsafe(testCase[1])
		c := NewFromStringUnsafe(testCase[2])
		res := relError(a, b)

		if !res.ExactEq(c) {
			t.Fatalf("expected error to be %s but got %s", c.String(), res.String())
		}
	}
}
