package cmdbuilder

import (
	"testing"

	"hack/model"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v3"
)

func TestMergeCommandGroup(t *testing.T) {
	want := &cli.Command{
		Name: "encode",
		Commands: []*cli.Command{
			{
				Name:  "json",
				Usage: "Get JSON",
				Flags: []cli.Flag{&cli.BoolFlag{Name: "pretty"}},
			},
			{Name: "xml", Usage: "Get XML", Flags: []cli.Flag{}},
		},
	}

	groups := []CommandGroup{
		{
			BinaryFile: BinaryFile{
				FullPath: ".bin/hack-core",
				Name:     "hack-core",
			},
			CmdInfos: []model.CmdInfo{{
				Name:  "json",
				Usage: "Get JSON",
				Flags: []model.CmdFlag{{Type: "bool", Name: "pretty"}},
			}},
		},
		{
			BinaryFile: BinaryFile{
				FullPath: ".bin/hack-core",
				Name:     "hack-core",
			},
			CmdInfos: []model.CmdInfo{{
				Name:  "xml",
				Usage: "Get XML",
			}},
		},
	}

	cmd := mergeCommandGroup("encode", groups)

	// clean action before compare
	for i := range cmd.Commands {
		cmd.Commands[i].Action = nil
	}
	assert.NotNil(t, cmd)
	assert.Equal(t, want, cmd)
}
