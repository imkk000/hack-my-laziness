package authentication

import "github.com/urfave/cli/v3"

func BuildCommands(baseCmds []*cli.Command) []*cli.Command {
	return append(baseCmds, commands...)
}

var commands = []*cli.Command{{
	Name:     "auth",
	Commands: []*cli.Command{},
}}
