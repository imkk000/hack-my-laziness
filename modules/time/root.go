package time

import (
	"context"
	"fmt"
	"strconv"
	"time"

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
	{
		Name: "conv",
		Commands: []*cli.Command{
			{
				Name:  "time",
				Usage: "Get Time",
				Action: func(_ context.Context, c *cli.Command) error {
					input := c.Args().Get(0)
					timestamp, err := strconv.ParseInt(input, 10, 64)
					if err != nil {
						return fmt.Errorf("parse uint64: %w", err)
					}
					date := time.Unix(timestamp, 0)

					fmt.Println(date.Format(time.RFC3339Nano))

					return nil
				},
			},
		},
	},
}
