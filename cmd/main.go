package main

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"

	"hack/pkg/cmdbuilder"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		slog.Error("get user home", "err", err)
		os.Exit(1)
	}
	binaryPath := filepath.Join(homeDir, ".bin")

	rootCmd := cmdbuilder.BuildRootCommand(binaryPath)
	if err := rootCmd.Run(context.Background(), os.Args); err != nil {
		slog.Error("command run", "err", err)
		os.Exit(1)
	}
}
