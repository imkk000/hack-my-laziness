package conv

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/urfave/cli/v3"
)

func BuildCommands(baseCmds []*cli.Command) []*cli.Command {
	return append(baseCmds, commands...)
}

var commands = []*cli.Command{
	{
		Name: "conv",
		Commands: []*cli.Command{
			{
				Name: "num",
				Action: func(_ context.Context, c *cli.Command) error {
					input := c.Args().First()
					n, err := strconv.ParseInt(input, 10, 64)
					if err != nil {
						return fmt.Errorf("parse integer: %w", err)
					}
					fmt.Printf("b: %b\n", n)
					fmt.Printf("o: %o\n", n)
					fmt.Printf("d: %d\n", n)
					fmt.Printf("h: %x\n", n)

					return nil
				},
			},
			{
				Name:  "json2struct",
				Usage: "JSON to Go Struct",
				Action: func(_ context.Context, c *cli.Command) error {
					input := strings.Join(c.Args().Slice(), "")
					var data map[string]any
					if err := json.Unmarshal([]byte(input), &data); err != nil {
						panic(err)
					}

					fmt.Println(generateStruct("Root", data))

					return nil
				},
			},
		},
	},
}
