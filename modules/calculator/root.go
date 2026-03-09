package calculator

import (
	"context"
	"fmt"
	"strings"

	"github.com/expr-lang/expr"
	"github.com/urfave/cli/v3"
)

func BuildCommands(baseCmds []*cli.Command) []*cli.Command {
	return append(baseCmds, commands...)
}

var commands = []*cli.Command{
	{
		Name: "calc",
		Commands: []*cli.Command{
			{
				Name:  "expr",
				Usage: "Evaluate Expression",
				Action: func(_ context.Context, c *cli.Command) error {
					exp := strings.Join(c.Args().Slice(), "")
					result, err := expr.Eval(exp, nil)
					if err != nil {
						return fmt.Errorf("evaluate expr: %w", err)
					}
					fmt.Println(result)

					return nil
				},
			},
		},
	},
}
