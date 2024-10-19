package cmp

import (
	"math/big"
	"strings"
)

func toNumArray(s string) ([]*big.Float, bool) {
	parts := strings.Fields(s)
	res := []*big.Float{}

	for _, part := range parts {
		num := new(big.Float)

		if _, ok := num.SetString(part); !ok {
			return nil, false
		}

		res = append(res, num)
	}

	return res, true
}

func LineToComparable(line string) Comparable {
	if len(line) == 0 {
		return Empty{}
	}

	nums, ok := toNumArray(line)

	if ok {
		return NumArray{nums: nums, rawData: line}
	}

	return RawString{value: line}
}
