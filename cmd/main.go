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
	}
	binaryPath := filepath.Join(homeDir, ".bin")
	slog.Info("get binary path", "path", binaryPath)

	rootCmd := cmdbuilder.BuildRootCommand(binaryPath)
	if err := rootCmd.Run(context.Background(), os.Args); err != nil {
		slog.Error("command run", "err", err)
		os.Exit(1)
	}
}
