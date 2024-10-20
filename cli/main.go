package cli

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func App(packageName, description, descriptionLong, version string) error {
	app := &cli.App{
		Name:                   packageName,
		Version:                version,
		Usage:                  description,
		UsageText:              fmt.Sprintf("%s [global options] files", packageName),
		Description:            descriptionLong,
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "numbers",
				Aliases: []string{"n"},
				Value:   true,
				Usage:   "Treat a line with one or multiple numbers as a number array, and use big float logic for comparison",
			},
			&cli.BoolFlag{
				Name:    "ignore-empty",
				Aliases: []string{"i"},
				Value:   true,
				Usage:   "Ignore empty lines",
			},
			&cli.BoolFlag{
				Name:    "trim",
				Aliases: []string{"t"},
				Value:   true,
				Usage:   "Trim leading and trailing spaces from lines",
			},
			&cli.BoolFlag{
				Name:    "wrong",
				Aliases: []string{"w"},
				Value:   false,
				Usage:   "Ignore correct lines and show only wrong ones",
			},
			&cli.BoolFlag{
				Name:    "exit",
				Aliases: []string{"x"},
				Value:   false,
				Usage:   "Exit early when it finds the first incorrect answer",
			},
			&cli.BoolFlag{
				Name:    "linenum",
				Aliases: []string{"l"},
				Value:   false,
				Usage:   "Show line numbers",
			},
			&cli.BoolFlag{
				Name:    "no-color",
				Aliases: []string{"c"},
				Value:   false,
				Usage:   "Show text without colors",
			},
			&cli.BoolFlag{
				Name:    "short",
				Aliases: []string{"s"},
				Value:   false,
				Usage:   "Shorten long outputs to avoid cluttering the screen",
			},
			&cli.BoolFlag{
				Name:    "relative",
				Aliases: []string{"r"},
				Value:   false,
				Usage:   "Use relative error instead of absolute",
			},
			&cli.BoolFlag{
				Name:    "duration",
				Aliases: []string{"d"},
				Value:   false,
				Usage:   "Display execution time",
			},
			&cli.StringFlag{
				Name:    "error",
				Aliases: []string{"e"},
				Value:   defaultError,
				Usage:   "Specify allowed error when comparing real numbers",
			},
		},
		Action: func(ctx *cli.Context) error {
			args := ctx.Args().Slice()

			return mainCommand(newOptions(ctx), args)
		},
	}

	return app.Run(os.Args)
}
