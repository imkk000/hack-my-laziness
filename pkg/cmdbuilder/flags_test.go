package cmdbuilder

import (
	"testing"

	"hack/model"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v3"
)

func TestBuildFlags(t *testing.T) {
	want := []cli.Flag{
		&cli.StringFlag{Name: "output"},
		&cli.StringSliceFlag{Name: "outputs"},
		&cli.StringSliceFlag{Name: "numbers"},
		&cli.BoolFlag{Name: "pretty"},
	}

	cmdFlags := []model.CmdFlag{
		{Type: "string", Name: "output"},
		{Type: "[]string", Name: "outputs"},
		{Type: "[]int64", Name: "numbers"},
		{Type: "bool", Name: "pretty"},
	}
	flags := buildFlags(cmdFlags)

	assert.Equal(t, want, flags)
}
