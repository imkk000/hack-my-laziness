package url

import (
	"context"
	"fmt"
	"strings"

	urlpkg "net/url"

	"github.com/pkg/browser"
	"github.com/urfave/cli/v3"
)

func BuildCommands(baseCmds []*cli.Command) []*cli.Command {
	return append(baseCmds, commands...)
}

var commands = []*cli.Command{
	{
		Name:  "view",
		Usage: "View",
		Commands: []*cli.Command{
			{
				Name:  "url",
				Usage: "Url",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "provider",
						Aliases:  []string{"p"},
						Required: true,
					},
					&cli.BoolFlag{
						Name:    "web",
						Aliases: []string{"w"},
					},
				},
				Action: func(_ context.Context, c *cli.Command) error {
					provider := c.String("provider")
					format, valid := providers.In(provider)
					if !valid {
						return fmt.Errorf("provider %s: not found", provider)
					}

					keywords := strings.Join(c.Args().Slice(), " ")
					q := urlpkg.QueryEscape(keywords)
					fullURL := fmt.Sprintf(format, q)

					if !c.Bool("web") {
						fmt.Println(fullURL)

						return nil
					}
					if err := browser.OpenURL(fullURL); err != nil {
						return fmt.Errorf("open browser: %w", err)
					}

					return nil
				},
			},
		},
	},
}
