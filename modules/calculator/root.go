package calculator

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/chzyer/readline"
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
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "interactive",
						Aliases: []string{"i"},
					},
				},
				Action: func(_ context.Context, c *cli.Command) error {
					if !c.Bool("interactive") {
						s := strings.Join(c.Args().Slice(), "")
						result, err := evaluateExpr(s)
						if err != nil {
							return err
						}
						fmt.Println(result)

						return nil
					}

					rl, err := readline.New("> ")
					if err != nil {
						return fmt.Errorf("new interactive: %w", err)
					}
					defer rl.Close()

					for {
						line, err := rl.Readline()
						if err != nil {
							break
						}

						result, err := evaluateExpr(line)
						if err != nil {
							slog.Error("eval",
								"expr", line,
								"err", err,
							)
							continue
						}

						fmt.Println(result)
					}

					return nil
				},
			},
		},
	},
}

func evaluateExpr(s string) (any, error) {
	result, err := expr.Eval(s, nil)
	if err != nil {
		return 0, fmt.Errorf("evaluate expr: %w", err)
	}

	return result, nil
}
