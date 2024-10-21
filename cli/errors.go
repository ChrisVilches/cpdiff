package cli

type NotAcceptedError struct {
}

func (NotAcceptedError) Error() string {
	return "Not Accepted"
}
