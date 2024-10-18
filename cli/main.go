package cli

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func App() {
	app := &cli.App{
		Name:                   "cpdiff",
		Usage:                  "Compare two files (or use stdin as input)",
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
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
			return mainCommand(
				ctx.Bool("short"),
				!ctx.Bool("no-color"),
				ctx.Bool("duration"),
				ctx.Bool("linenum"),
				ctx.Bool("relative"),
				ctx.Bool("abort"),
				ctx.Bool("wrong"),
				ctx.String("error"),
				args,
			)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// TODO: Add flag to show all lines including empty ones (this is hard because I need to format properly and know which ones to compare to which ones)
//       Also I need to find a good case example from my Algorithms repository (which I don't have right now.)
//       Or just create some dummy files for testing.
