package cmp

type Empty struct{}

func (Empty) Display() string {
	return "-"
}

func (Empty) ShortDisplay() string {
	return "-"
}

func (Empty) Type() ComparableType {
	return ComparableTypes.Empty
}
