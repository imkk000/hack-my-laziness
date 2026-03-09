package network

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/urfave/cli/v3"
)

func BuildCommands(baseCmds []*cli.Command) []*cli.Command {
	return append(baseCmds, commands...)
}

var commands = []*cli.Command{{
	Name:  "show",
	Usage: "Show",
	Commands: []*cli.Command{{
		Name:  "port",
		Usage: "Port",
		Action: func(_ context.Context, _ *cli.Command) error {
			cmd := exec.Command("lsof", "-iTCP", "-sTCP:LISTEN", "-P", "-n")
			out, err := cmd.Output()
			if err != nil {
				return fmt.Errorf("exec command: %w", err)
			}
			fmt.Println(string(out))

			return nil
		},
	}},
}}
