package cmdbuilder

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"syscall"

	"hack/model"

	"github.com/urfave/cli/v3"
)

func BuildRootCommand(binaryPath string) *cli.Command {
	rootCmd := &cli.Command{
		Name:                  "hack",
		Usage:                 "Get Modules",
		EnableShellCompletion: true,
		Commands: []*cli.Command{{
			Name: "my",
			Commands: []*cli.Command{{
				Name: "life",
				Action: func(_ context.Context, _ *cli.Command) error {
					fmt.Println(`
    /\_____/\
   /  o   o  \
  ( ==  w  == )
   )         (
  (           )
 ( (  )   (  ) )
(__(__)___(__)__)`)

					return nil
				},
			}},
		}},
	}

	// get all binary names
	binaries := DiscoverBinaries(binaryPath, binaryNamePrefix)

	// fetch completion in JSON then group them with subcommand
	groups := make(map[string][]CommandGroup)
	for _, file := range binaries {
		cmdInfo, err := FetchCompletion(file.FullPath)
		if err != nil {
			slog.Error("get completion",
				"path", file.FullPath,
				"name", file.Name,
			)
			continue
		}
		cmdInfo.Name = file.Name

		// group with embed binary info
		// (keep full binary path and command info list)
		for _, cmd := range cmdInfo.Commands {
			groups[cmd.Name] = append(groups[cmd.Name], CommandGroup{
				BinaryFile: file,
				CmdInfos:   cmd.Commands,
			})
		}
	}

	for subCommand, group := range groups {
		cmd := mergeCommandGroup(subCommand, group)
		rootCmd.Commands = append(rootCmd.Commands, cmd)
	}

	return rootCmd
}

func mergeCommandGroup(subCommand string, groups []CommandGroup) *cli.Command {
	var cmds []*cli.Command
	for _, cmd := range groups {
		for _, info := range cmd.CmdInfos {
			cmd := buildCommand(info, cmd.FullPath)
			cmds = append(cmds, cmd)
		}
	}

	return &cli.Command{
		Name:     subCommand,
		Commands: cmds,
	}
}

func buildCommand(info model.CmdInfo, binaryName string) *cli.Command {
	cmd := &cli.Command{
		Name:  info.Name,
		Usage: info.Usage,
		Flags: buildFlags(info.Flags),
	}

	if len(info.Commands) > 0 {
		for _, subCommand := range info.Commands {
			cmd.Commands = append(cmd.Commands, buildCommand(subCommand, binaryName))
		}

		return cmd
	}

	cmd.Action = func(_ context.Context, c *cli.Command) error {
		argv := append([]string{binaryName}, c.Root().Args().Slice()...)

		return syscall.Exec(binaryName, argv, os.Environ())
	}

	return cmd
}

type CommandRunner interface {
	Run(name string, args ...string) ([]byte, error)
}

type CommandRunnerImpl int

func NewCommandRunner() *CommandRunnerImpl {
	return new(CommandRunnerImpl)
}

func (CommandRunnerImpl) Run(name string, args ...string) ([]byte, error) {
	return exec.Command(name, args...).Output()
}
