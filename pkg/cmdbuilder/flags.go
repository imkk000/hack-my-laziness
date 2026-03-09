package cmdbuilder

import (
	"strings"

	"hack/model"

	"github.com/urfave/cli/v3"
)

func buildFlags(flags []model.CmdFlag) []cli.Flag {
	cFlags := make([]cli.Flag, len(flags))
	for i, flag := range flags {
		// TODO: start with 2 types first, add more later (default is string)
		switch flag.Type {
		case "bool":
			cFlags[i] = &cli.BoolFlag{Name: flag.Name}
		default:
			if strings.HasPrefix(flag.Type, "[]") {
				cFlags[i] = &cli.StringSliceFlag{Name: flag.Name}
				continue
			}

			cFlags[i] = &cli.StringFlag{Name: flag.Name}
		}
	}

	return cFlags
}
