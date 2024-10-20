package cli

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func warn(w error) {
	fmt.Fprint(os.Stderr, "Warning: ")
	fmt.Fprintln(os.Stderr, w)
}

func printfLnColor(c color.Attribute, useColor bool, s string, printArgs ...any) {
	if !useColor {
		fmt.Println(fmt.Sprintf(s, printArgs...))

		return
	}

	fmt.Println(color.New(c).Sprintf(s, printArgs...))
}
