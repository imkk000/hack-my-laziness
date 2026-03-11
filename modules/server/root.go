package server

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/urfave/cli/v3"
)

func BuildCommands(baseCmds []*cli.Command) []*cli.Command {
	return append(baseCmds, commands...)
}

var commands = []*cli.Command{
	{
		Name:  "run",
		Usage: "Run a service",
		Commands: []*cli.Command{
			{
				Name:  "http",
				Usage: "Run an API stub HTTP server from a config file",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "config",
						Aliases: []string{"c"},
						Usage:   "Path to YAML config file",
						Value:   "stub.yaml",
					},
					&cli.StringFlag{
						Name:    "port",
						Aliases: []string{"p"},
						Usage:   "Port to listen on",
						Value:   "8080",
					},
				},
				Action: func(_ context.Context, c *cli.Command) error {
					cfgPath := c.String("config")
					port := c.String("port")

					cfg, err := loadConfig(cfgPath)
					if err != nil {
						return fmt.Errorf("load config: %w", err)
					}

					e := echo.New()
					e.Use(middleware.Recover())

					fmt.Println("Registered routes:")
					registerRoutes(e, cfg.Routes)

					return e.Start("127.0.0.1:" + port)
				},
			},
		},
	},
}
