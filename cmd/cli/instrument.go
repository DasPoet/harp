package cli

import (
	"bytes"
	"os"

	"github.com/daspoet/harp"
	"github.com/daspoet/harp/pkg/ast"
	"github.com/daspoet/harp/pkg/instrument"
	"github.com/daspoet/harp/pkg/pattern"
	"github.com/urfave/cli/v2"
)

var instrumentCmd = cli.Command{
	Name:  "add-timings",
	Usage: "Add timings to specific methods",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "type",
			Usage: "the type(s) to instrument",
		},
		&cli.StringFlag{
			Name:  "method",
			Usage: "the method to instrument",
		},
		&cli.StringFlag{
			Name:  "visibility",
			Value: "*",
			Usage: "the visibility of the method(s) to instrument",
		},
	},
	Action: instrumentAction,
}

func instrumentAction(ctx *cli.Context) error {
	filename, typ, method, visibility, err := prepareInstrumentArgs(ctx)
	if err != nil {
		return err
	}
	return doInstrument(filename, typ, method, visibility)
}

func prepareInstrumentArgs(ctx *cli.Context) (string, pattern.Pattern, pattern.Pattern, ast.Visibility, error) {
	filename := ctx.Args().First()
	if filename == "" {
		return "", nil, nil, 0, cli.Exit("missing argument FILENAME", 1)
	}

	var (
		typ    pattern.Pattern
		method pattern.Pattern
	)

    if rawType := ctx.String("type"); rawType != "" {
        typ = pattern.WithWildcards(rawType)
    }
    if rawMethod := ctx.String("method"); rawMethod != "" {
        method = pattern.WithWildcards(rawMethod)
    }

	visibility, err := ast.ParseVisibility(ctx.String("visibility"))
	if err != nil {
		return "", nil, nil, 0, err
	}
	return filename, typ, method, visibility, nil
}

func doInstrument(filename string, typ, method pattern.Pattern, visibility ast.Visibility) error {
	timings := harp.Timings{
		Targets: []instrument.Target{
			{
				Type:       typ,
				Method:     method,
				Visibility: visibility,
			},
		},
	}

	var out bytes.Buffer
	if err := instrument.InstrumentFile(filename, nil, &out, timings); err != nil {
		return err
	}
	return os.WriteFile(filename, out.Bytes(), os.ModePerm)
}
