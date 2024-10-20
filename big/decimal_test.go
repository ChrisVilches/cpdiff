package big

import (
	"testing"
)

func TestInsideRange(t *testing.T) {
	dataInside := [][]float64{
		{1, 0, 2},
		{0, 0, 2},
		{2, 0, 2},
		{1.5, 0, 2},
	}

	dataOutside := [][]float64{
		{-0.000001, 0, 2},
		{2.000001, 0, 2},
		{5, 0, 2},
		{-3.5, 0, 2},
	}

	for _, testCase := range dataInside {
		num := NewFromFloat64(testCase[0])
		u := testCase[1]
		v := testCase[2]

		if !num.InsideRange(u, v) {
			t.Fatalf("expected %s to be inside range [%f, %f]", num.String(), u, v)
		}
	}

	for _, testCase := range dataOutside {
		num := NewFromFloat64(testCase[0])
		u := testCase[1]
		v := testCase[2]

		if num.InsideRange(u, v) {
			t.Fatalf("expected %s to be outside range [%f, %f]", num.String(), u, v)
		}
	}
}

func TestMax(t *testing.T) {
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

		if !Max(a, b).ExactEq(c) {
			t.Fatalf("max function (%s, %s) expected to be %s", a.String(), b.String(), c.String())
		}
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
