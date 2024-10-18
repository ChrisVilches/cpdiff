package comparison

type RawString struct {
	value   string
	rawData *string
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

// TODO: BEST IDEA. Add another method to do a "ShortDisplay" which everyone has to implement.
// then I call that from the parent. That way I remove dependencies to parameters that can vary
// and makes the polymorphism fucked up.

// TODO: Maybe using string for the type would be computationally expensive.
func (*RawString) Type() string {
	return "string"
}

func compareStrings(a, b *RawString) int {
	if a.value == b.value {
		return Correct
	} else {
		return Incorrect
	}
}
