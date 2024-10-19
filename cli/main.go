package cli

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func App() {
	app := &cli.App{
		Name:                   "cpdiff",
		Usage:                  "Competitive Programming Difference tool. Compare two files (or stdin).",
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "trim",
				Aliases: []string{"t"},
				Value:   true,
				// TODO: This one affects mostly how empty lines and raw strings are handled, but not number sequences.
				// Maybe rename it "strict presentation" --strict-presentation -p
				// THough that still doesn't explain why numbers are still kept
				// Maybe just explain that it only affects how empty and raw strings are handled, but not numbers or number sequences.
				// Yes, just put that phrase in the explanation, explain it in few words.
				Usage: "Trim and remove trailing/leading whitespace and ignore empty lines (lines that contain numbers or number sequences are still compared item-wise, ignoring extra whitespace)",
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
			return mainCommand(newOptions(ctx, true), args)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
