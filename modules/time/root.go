package time

import (
	"github.com/urfave/cli/v3"
)

func BuildCommands(baseCmds []*cli.Command) []*cli.Command {
	return append(baseCmds, commands...)
}

var commands = []*cli.Command{
	{
		Name: "get",
		Commands: []*cli.Command{
			{
				Name:  "time",
				Usage: "Get Time",
				Commands: []*cli.Command{
					getTimeNow,
					getTimeZone,
				},
			},
		},
	},
}
