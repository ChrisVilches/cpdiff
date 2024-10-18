package comparison

type Comparable interface {
	Type() ComparableType
	Display() string
	ShortDisplay() string
}

type ComparableType int

var ComparableTypes = struct {
	RawString ComparableType
	NumArray  ComparableType
	Empty     ComparableType
}{
	RawString: 0,
	NumArray:  1,
	Empty:     2,
}
