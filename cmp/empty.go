package cmp

import "github.com/ChrisVilches/cpdiff/big"

type Empty struct{}

func (Empty) Display() string {
	return "-"
}

func (Empty) ShortDisplay(int) string {
	return "-"
}

func (Empty) compare(
	c Comparable,
	_ bool,
	_ big.Decimal,
) ([]verdictRange, big.Decimal) {
	if _, ok := c.(Empty); ok {
		return []verdictRange{{Value: Verdicts.Correct}}, big.Decimal{}
	}

	return []verdictRange{{Value: Verdicts.Incorrect}}, big.Decimal{}
}
