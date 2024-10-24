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
	Comparable,
	bool,
	big.Decimal,
) ([]verdictRange, big.Decimal) {
	return []verdictRange{{Value: Verdicts.Correct}}, big.Decimal{}
}
