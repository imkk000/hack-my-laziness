package calculator

import (
	"context"
	"fmt"
	"math"
	"strconv"

	"github.com/urfave/cli/v3"
)

func BuildCommands(baseCmds []*cli.Command) []*cli.Command {
	return append(baseCmds, commands...)
}

var commands = []*cli.Command{
	{
		Name:     "calc",
		Commands: generateMethod([]*cli.Command{calcExpr}),
	},
}

var fns = map[string]func(float64) float64{
	"log2":  math.Log2,
	"log10": math.Log10,
	"sqrt":  math.Sqrt,
	"sin":   math.Sin,
	"cos":   math.Cos,
	"tan":   math.Tan,
}

func generateMethod(cmd []*cli.Command) []*cli.Command {
	for name, fn := range fns {
		cmd = append(cmd, &cli.Command{
			Name: name,
			Action: func(_ context.Context, c *cli.Command) error {
				input := c.Args().Get(0)
				dec, err := strconv.ParseFloat(input, 64)
				if err != nil {
					return fmt.Errorf("parse float: %w", err)
				}
				fmt.Println(fn(dec))

				return nil
			},
		})
	}

	return cmd
}
