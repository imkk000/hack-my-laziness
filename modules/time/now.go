package time

import (
	"context"
	"fmt"
	"time"

	"github.com/urfave/cli/v3"
)

var getTimeNow = &cli.Command{
	Name:  "now",
	Usage: "Now",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "format",
			Aliases: []string{"f"},
			Value:   time.RFC3339Nano,
		},
		&cli.StringFlag{
			Name:    "location",
			Aliases: []string{"l"},
			Value:   time.Local.String(),
		},
	},
	Action: func(_ context.Context, c *cli.Command) error {
		loc, err := time.LoadLocation(c.String("location"))
		if err != nil {
			return fmt.Errorf("load location: %w", err)
		}
		format := c.String("format")
		fmt.Println(time.Now().In(loc).Format(format))

		return nil
	},
}
