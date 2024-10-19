package cli

import (
	"cpdiff/util"
	"fmt"
	"math/big"

	"github.com/urfave/cli/v2"
)

// TODO: Should this be capitalized simply because it's a struct? I don't want to export it.
// The same applies to other structs I may have somewhere else.
type options struct {
	short            bool
	showColor        bool
	showDuration     bool
	showLineNum      bool
	useRelativeError bool
	abortEarly       bool
	showOnlyWrong    bool
	error            *big.Float
}

func getConfigError(errorString string, outputWarnings bool) (res *big.Float) {
	res = new(big.Float)

	res.SetString(defaultError)

	if util.IsEmptyLine(errorString) {
		return
	}

	parsedVal := new(big.Float)

	if _, ok := parsedVal.SetString(errorString); !ok || util.BigFloatOutsideRange(parsedVal, 0, 1) {
		if outputWarnings {
			warn(fmt.Sprintf("Error value is incorrect. Using default value %s\n", defaultError))
		}
		return
	}

	res.Set(parsedVal)
	return
}

func newOptions(ctx *cli.Context, outputWarnings bool) options {
	res := options{
		short:            ctx.Bool("short"),
		showColor:        !ctx.Bool("no-color"),
		showDuration:     ctx.Bool("duration"),
		showLineNum:      ctx.Bool("linenum"),
		useRelativeError: ctx.Bool("relative"),
		abortEarly:       ctx.Bool("abort"),
		showOnlyWrong:    ctx.Bool("wrong"),
		error:            getConfigError(ctx.String("error"), outputWarnings),
	}

	return res
}
