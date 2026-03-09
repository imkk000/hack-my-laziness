package encoding

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/urfave/cli/v3"
)

func BuildCommands(baseCmds []*cli.Command) []*cli.Command {
	return append(baseCmds, commands...)
}

var commands = []*cli.Command{
	{
		Name: "encode",
		Commands: []*cli.Command{
			{
				Name: "base64",
				Action: func(_ context.Context, c *cli.Command) error {
					data := strings.Join(c.Args().Slice(), " ")
					fmt.Println(base64.StdEncoding.EncodeToString([]byte(data)))
					return nil
				},
			},
		},
	},
	{
		Name: "decode",
		Commands: []*cli.Command{
			{
				Name: "base64",
				Action: func(_ context.Context, c *cli.Command) error {
					data := strings.Join(c.Args().Slice(), " ")
					raw, err := base64.StdEncoding.DecodeString(data)
					if err != nil {
						return fmt.Errorf("decode base64: %w", err)
					}
					fmt.Println(string(raw))

					return nil
				},
			},
		},
	},
}

// base64
