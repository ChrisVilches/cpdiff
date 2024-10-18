package comparison

// TODO: change name?? it's not a comparable interface anymore
// this is because the different comparing functions return different data... numbers return biggest difference,
// while strings don't.
type Comparable interface {
	Type() string
	Display() string
	ShortDisplay() string
}
