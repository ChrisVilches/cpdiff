package cli

import (
	"fmt"
	"os"
)

func warn(w error) {
	fmt.Fprint(os.Stderr, "Warning: ")
	fmt.Fprintln(os.Stderr, w)
}
