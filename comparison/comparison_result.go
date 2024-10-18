package comparison

type ComparisonResult int

var ComparisonResults = struct {
	Correct   ComparisonResult
	Incorrect ComparisonResult
	Approx    ComparisonResult
}{
	Correct:   0,
	Incorrect: 1,
	Approx:    2,
}
