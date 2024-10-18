package comparison

import (
	"cpdiff/util"
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
	if util.IsEmptyLine(line) {
		return Empty{}
	}

	nums, ok := toNumArray(line)

	if ok {
		return NumArray{nums: nums, rawData: line}
	} else {
		return RawString{value: line, rawData: line}
	}
}
