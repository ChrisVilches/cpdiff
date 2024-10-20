package cli

import (
	"fmt"
	"math/big"
	"os"

	"cpdiff/util"

	"github.com/urfave/cli/v2"
)

type options struct {
	short            bool
	showColor        bool
	showDuration     bool
	showLineNum      bool
	skipEmptyLines   bool
	useRelativeError bool
	abortEarly       bool
	showOnlyWrong    bool
	trim             bool
	numbers          bool
	error            *big.Float
}

func getConfigError(errorString string) (*big.Float, *string) {
	res := new(big.Float)

	res.SetString(defaultError)

	if len(errorString) == 0 {
		return res, nil
	}

	parsedVal := new(big.Float)
	_, ok := parsedVal.SetString(errorString)

	if !ok || util.BigFloatOutsideRange(parsedVal, 0, 1) {
		warn := fmt.Sprintf(
			"Error value is incorrect. Using default value %s\n",
			defaultError,
		)
		return res, &warn
	}

	res.Set(parsedVal)

	return res, nil
}

func newOptions(ctx *cli.Context) options {
	err, warn := getConfigError(ctx.String("error"))

	if warn != nil {
		fmt.Fprint(os.Stderr, warn)
	}

	res := options{
		trim:             ctx.Bool("trim"),
		short:            ctx.Bool("short"),
		showColor:        !ctx.Bool("no-color"),
		showDuration:     ctx.Bool("duration"),
		showLineNum:      ctx.Bool("linenum"),
		useRelativeError: ctx.Bool("relative"),
		abortEarly:       ctx.Bool("exit"),
		showOnlyWrong:    ctx.Bool("wrong"),
		skipEmptyLines:   ctx.Bool("ignore-empty"),
		numbers:          ctx.Bool("numbers"),
		error:            err,
	}

	return res
}
