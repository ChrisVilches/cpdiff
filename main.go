package main

import (
	_ "embed"
	"fmt"
	"github.com/ChrisVilches/cpdiff/cli"
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

	err := cli.App(packageName, description, descriptionLong, version)

	if _, notAccepted := err.(cli.NotAcceptedError); notAccepted {
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

// TODO: This project needs a massive code review.
