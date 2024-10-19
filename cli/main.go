package cli

import (
	"github.com/urfave/cli/v2"
	"os"
)

func App() error {
	app := &cli.App{
		Name:                   "cpdiff",
		Usage:                  "File difference tool for competitive programming",
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
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
				Name:    "abort",
				Aliases: []string{"a"},
				Value:   false,
				Usage:   "Abort when it finds the first incorrect answer",
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
