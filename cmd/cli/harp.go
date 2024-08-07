package cli

import (
	"github.com/urfave/cli/v2"
)

var App = cli.App{
	Name:  "harp",
	Usage: "instrument your Go code with ease",
	Commands: []*cli.Command{
        &instrumentCmd,
    },
}
