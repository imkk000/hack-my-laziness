package http

import (
	"context"

	"github.com/urfave/cli/v3"
)

func BuildCommands(baseCmds []*cli.Command) []*cli.Command {
	return append(baseCmds, commands...)
}

var headerFlag = &cli.StringSliceFlag{
	Name:    "header",
	Aliases: []string{"H"},
	Usage:   `Custom header, e.g. -H "Authorization: Bearer token"`,
}

var bodyFlag = &cli.StringFlag{
	Name:    "body",
	Aliases: []string{"b"},
	Usage:   "Request body as string",
}

var fileFlag = &cli.StringFlag{
	Name:    "file",
	Aliases: []string{"f"},
	Usage:   "Read request body from file",
}

var contentTypeFlag = &cli.StringFlag{
	Name:    "content-type",
	Aliases: []string{"c"},
	Usage:   "Content type: json, form, xml, text (default: json when body is set)",
	Value:   "json",
}

func bodyFlags() []cli.Flag {
	return []cli.Flag{headerFlag, bodyFlag, fileFlag, contentTypeFlag}
}

var commands = []*cli.Command{
	{
		Name:  "get",
		Usage: "Send GET request",
		Commands: []*cli.Command{{
			Name:  "url",
			Flags: []cli.Flag{headerFlag},
			Action: func(_ context.Context, c *cli.Command) error {
				return doRequest("GET", RequestOptions{
					URL:     c.Args().First(),
					Headers: c.StringSlice("header"),
				})
			},
		}},
	},
	{
		Name:  "post",
		Usage: "Send POST request",
		Commands: []*cli.Command{{
			Name:  "url",
			Flags: bodyFlags(),
			Action: func(_ context.Context, c *cli.Command) error {
				return doRequest("POST", RequestOptions{
					URL:         c.Args().First(),
					Headers:     c.StringSlice("header"),
					Body:        c.String("body"),
					BodyFile:    c.String("file"),
					ContentType: c.String("content-type"),
				})
			},
		}},
	},
	{
		Name:  "put",
		Usage: "Send PUT request",
		Commands: []*cli.Command{{
			Name:  "url",
			Flags: bodyFlags(),
			Action: func(_ context.Context, c *cli.Command) error {
				return doRequest("PUT", RequestOptions{
					URL:         c.Args().First(),
					Headers:     c.StringSlice("header"),
					Body:        c.String("body"),
					BodyFile:    c.String("file"),
					ContentType: c.String("content-type"),
				})
			},
		}},
	},
	{
		Name:  "patch",
		Usage: "Send PATCH request",
		Commands: []*cli.Command{{
			Name:  "url",
			Flags: bodyFlags(),
			Action: func(_ context.Context, c *cli.Command) error {
				return doRequest("PATCH", RequestOptions{
					URL:         c.Args().First(),
					Headers:     c.StringSlice("header"),
					Body:        c.String("body"),
					BodyFile:    c.String("file"),
					ContentType: c.String("content-type"),
				})
			},
		}},
	},
	{
		Name:  "delete",
		Usage: "Send DELETE request",
		Commands: []*cli.Command{{
			Name:  "url",
			Flags: []cli.Flag{headerFlag},
			Action: func(_ context.Context, c *cli.Command) error {
				return doRequest("DELETE", RequestOptions{
					URL:     c.Args().First(),
					Headers: c.StringSlice("header"),
				})
			},
		}},
	},
}
