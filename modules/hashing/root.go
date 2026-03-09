package hashing

import (
	"context"
	"fmt"
	"strings"

	"github.com/urfave/cli/v3"
)

func BuildCommands(baseCmds []*cli.Command) []*cli.Command {
	return append(baseCmds, commands...)
}

var commands = []*cli.Command{
	{
		Name: "hash", Usage: "Hash",
		Commands: []*cli.Command{
			{Name: "md5", Action: runAction(hashMD5)},
			{Name: "sha1", Action: runAction(hashSHA1)},
			{Name: "sha256", Action: runAction(hashSHA256)},
			{Name: "sha512", Action: runAction(hashSHA512)},
			{Name: "sha3", Action: runAction(hashSHA3_256)},
			{Name: "blake2b", Action: runAction(hashBLAKE2b)},
		},
	},
}

func runAction(fn func(string) string) cli.ActionFunc {
	return func(_ context.Context, c *cli.Command) error {
		data := strings.Join(c.Args().Slice(), " ")
		fmt.Println(fn(data))

		return nil
	}
}

// TODO: support hash from file
