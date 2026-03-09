package main

import (
	"context"
	"fmt"
	"html/template"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/urfave/cli/v3"
)

const (
	binaryNameFormat = "hack-%s"
	modulePath       = "modules"
	templateFilename = "tools/gen_template/main.go.tmpl"
)

func main() {
	rootCmd := &cli.Command{
		Usage: "Generate module",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "module", Required: true},
			&cli.StringFlag{Name: "output", Required: true},
		},
		Action: func(_ context.Context, c *cli.Command) error {
			moduleName := c.String("module")
			path := filepath.Join(modulePath, moduleName)
			if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
				return fmt.Errorf("check exists module: %w", err)
			}

			name := fmt.Sprintf(binaryNameFormat, moduleName)
			outputPath := c.String("output")
			outputBinaryFilename := filepath.Join(outputPath, name)

			path, err := os.MkdirTemp(".", "*")
			if err != nil {
				return fmt.Errorf("make temp directory: %w", err)
			}
			defer os.RemoveAll(path)

			tmpl, err := template.ParseFiles(templateFilename)
			if err != nil {
				return fmt.Errorf("parse file: %w", err)
			}

			outputFilename := filepath.Join(path, "main.go")
			fs, err := os.Create(outputFilename)
			if err != nil {
				return fmt.Errorf("create file: %w", err)
			}
			defer fs.Close()

			if err := tmpl.Execute(fs, map[string]any{
				"Module": moduleName,
			}); err != nil {
				return fmt.Errorf("execute template: %w", err)
			}

			cmd := exec.Command("go", "build", "-o", outputBinaryFilename, outputFilename)
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout

			if err := cmd.Run(); err != nil {
				return fmt.Errorf("run build command: %w", err)
			}

			return nil
		},
	}
	if err := rootCmd.Run(context.Background(), os.Args); err != nil {
		slog.Error("run command", "err", err)
		os.Exit(1)
	}
}
