package cli

import (
	"cpdiff/big"
	"fmt"
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
	error            big.Decimal
}

func getConfigError(errorString string) (big.Decimal, error) {
	res := big.NewFromStringUnsafe(defaultError)

	if len(errorString) == 0 {
		return res, nil
	}

	parsedVal, ok := big.NewFromString(errorString)

	if !ok || !parsedVal.InsideRange(0, 1) {
		warn := fmt.Errorf(
			"error value is incorrect (using default value %s)",
			defaultError,
		)
		return res, warn
	}

	return parsedVal, nil
}

func newOptions(ctx *cli.Context) options {
	err, warnMsg := getConfigError(ctx.String("error"))

	if warnMsg != nil {
		warn(warnMsg)
	}

	res := options{
		trim:             ctx.Bool("trim"),
		short:            ctx.Bool("short"),
		showColor:        ctx.Bool("color"),
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
