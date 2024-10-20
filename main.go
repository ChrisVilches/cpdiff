package main

import (
	"cpdiff/cli"
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed .metadata/version
var version string

//go:embed .metadata/package-name
var packageName string

//go:embed .metadata/description
var description string

//go:embed .metadata/description-long
var descriptionLong string

func main() {
	version = strings.TrimSpace(version)
	packageName = strings.TrimSpace(packageName)
	description = strings.TrimSpace(description)
	descriptionLong = strings.TrimSpace(descriptionLong)

	if err := cli.App(packageName, description, descriptionLong, version); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

// TODO: This project needs a massive code review.
// TODO: Here's a good benchmark to optimize it.
// time c++ urionlinejudge/3021-jumping_grasshopper.cpp
//    < ~/dev/data/J/input/J_29
// c++ urionlinejudge/3021-jumping_grasshopper.cpp
//    < ~/dev/data/J/input/J_29 | cpdiff -d ~/dev/data/J/output/J_29
// 698.490613ms with cpdiff, 0.17s without
// Or just print to files, without C++ program involved
// cpdiff -d ~/dev/data/C/output/C_25 ~/dev/data/C/output/C_25
// coloring makes it more than 2x slower
// I should be able to get a better time (but it won't be exactly the same).
// However, before micro-optimizing, I need to test everything
// so that it doesn't break while changing it.
