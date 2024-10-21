package integration

import (
	"testing"
)

func TestStatusCode(t *testing.T) {
	expectEq(t, getStatusCode(1), 0)
	expectEq(t, getStatusCode(2), 1)
	expectEq(t, getStatusCode(3), 0)
	expectEq(t, getStatusCode(4), 1)
	expectEq(t, getStatusCode(5), 0)
	expectEq(t, getStatusCode(6), 1)
}
