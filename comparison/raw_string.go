package comparison

type RawString struct {
	value   string
	rawData string
}

const ellipses = "..."
const maxLength = 10

// TODO: Should this be reference?
func (r *RawString) Display() string {
	return r.value
}

func (r *RawString) ShortDisplay() string {
	if len(r.value) > maxLength {
		return r.value[0:maxLength-len(ellipses)] + ellipses
	}

	return r.value
}

func (*RawString) Type() ComparableType {
	return ComparableTypes.RawString
}

func compareStrings(a, b *RawString) ComparisonResult {
	if a.value == b.value {
		return ComparisonResults.Correct
	} else {
		return ComparisonResults.Incorrect
	}
}
