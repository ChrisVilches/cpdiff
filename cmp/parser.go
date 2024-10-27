package cmp

import (
	"github.com/ChrisVilches/cpdiff/big"
	"strings"
)

func toNumArray(s string) []big.Decimal {
	parts := strings.Fields(s)
	res := []big.Decimal{}

	if len(parts) == 0 {
		return nil
	}

	for _, part := range parts {
		if len(part) > 1 && part[0] == '0' && part[1] != '.' {
			return nil
		}

		val, ok := big.NewFromString(part)

		if !ok {
			return nil
		}

		res = append(res, val)
	}

	return res
}

func newComparable(line string) Comparable {
	if len(line) == 0 {
		return Empty{}
	}

	nums := toNumArray(line)

	if nums != nil {
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
