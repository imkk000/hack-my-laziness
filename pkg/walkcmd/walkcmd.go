package walkcmd

import (
	"reflect"

	"hack/model"

	"github.com/urfave/cli/v3"
)

func Walk(c *cli.Command) model.CmdInfo {
	info := model.CmdInfo{
		Name:  c.Name,
		Usage: c.Usage,
	}

	for _, cmd := range c.Commands {
		if cmdSkipper.Is(cmd.Name) {
			continue
		}
		info.Commands = append(info.Commands, Walk(cmd))
	}

	for _, flag := range c.Flags {
		var flags []model.CmdFlag
		for _, name := range flag.Names() {
			if cmdSkipper.Is(name) {
				continue
			}
			t := reflect.TypeOf(flag.Get()).String()
			docFlag := flag.(cli.DocGenerationFlag)
			flags = append(flags, model.CmdFlag{
				Type:  t,
				Name:  name,
				Usage: docFlag.GetUsage(),
			})
		}

		info.Flags = append(info.Flags, flags...)
	}

	return info
}
