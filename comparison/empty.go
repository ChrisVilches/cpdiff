package comparison

type Empty struct {
}

func (*Empty) Display() string {
	return "-"
}

func (*Empty) ShortDisplay() string {
	return "-"
}

func (*Empty) Type() string {
	return "empty"
}
