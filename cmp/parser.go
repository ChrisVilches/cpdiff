package cmp

import (
	"cpdiff/big"
	"strings"
)

func toNumArray(s string) ([]big.Decimal, bool) {
	parts := strings.Fields(s)
	res := []big.Decimal{}

	if len(parts) == 0 {
		return nil, false
	}

	for _, part := range parts {
		if len(part) > 1 && part[0] == '0' && part[1] != '.' {
			return nil, false
		}

		val, ok := big.NewFromString(part)

		if !ok {
			return nil, false
		}

		res = append(res, val)
	}

	return res, true
}

func newComparable(line string) Comparable {
	if len(line) == 0 {
		return Empty{}
	}

	nums, ok := toNumArray(line)

	if ok {
		return NumArray{nums: nums, rawData: line}
	}

	return RawString{value: line}
}

func newComparableNonNumeric(line string) Comparable {
	if len(line) == 0 {
		return Empty{}
	}

	return RawString{value: line}
}
