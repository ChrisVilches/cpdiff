package cli

// TODO: I think I need to tell the user whether I'm using absolute or relative error.
// i dont even know what's the diff between the two.

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func App() {
	app := &cli.App{
		Name:  "cpdiff",
		Usage: "Compare two files (or use stdin as input)",
		Flags: []cli.Flag{
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
				ctx.String("error"),
				args,
			)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// TODO: Add flag to skip correct lines.
// TODO: Flag to stop when both channels have different length
// TODO: Add flag to show all lines including empty ones (this is hard because I need to format properly and know which ones to compare to which ones)
