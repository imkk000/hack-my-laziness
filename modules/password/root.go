package password

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"

	"github.com/urfave/cli/v3"
)

func BuildCommands(baseCmds []*cli.Command) []*cli.Command {
	return append(baseCmds, commands...)
}

var commands = []*cli.Command{{
	Name:  "gen",
	Usage: "generate",
	Commands: []*cli.Command{{
		Name: "pass",
		Flags: []cli.Flag{
			&cli.Uint64Flag{
				Name:    "length",
				Aliases: []string{"l", "len"},
				Value:   32,
			},
			&cli.Uint64Flag{
				Name:    "count",
				Aliases: []string{"c"},
				Value:   1,
			},
			&cli.BoolFlag{
				Name:    "alpha",
				Aliases: []string{"a"},
				Value:   true,
			},
			&cli.BoolFlag{
				Name:    "number",
				Aliases: []string{"n"},
				Value:   true,
			},
			&cli.BoolFlag{
				Name:    "symbol",
				Aliases: []string{"s"},
			},
		},
		Action: func(_ context.Context, c *cli.Command) error {
			length := c.Uint64("length")
			count := c.Uint64("count")

			var charset string
			if c.Bool("alpha") {
				charset += "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
			}
			if c.Bool("number") {
				charset += "0123456789"
			}
			if c.Bool("symbol") {
				charset += "!@#$%^&*()-_=+[]{}|;:,.<>?"
			}
			if charset == "" {
				return errors.New("at least one character type required")
			}

			for i := uint64(0); i < count; i++ {
				pass, err := generatePassword(length, charset)
				if err != nil {
					return err
				}
				fmt.Println(pass)
			}

			return nil
		},
	}},
}}

func generatePassword(length uint64, charset string) (string, error) {
	password := make([]byte, length)
	password[0] = charset[0]

	// using crypto random instead of math random
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
	if err != nil {
		return "", err
	}
	password[0] = charset[n.Int64()]

	for i := uint64(1); i < length; i++ {
		// might take long time to generate,
		// but ensure no adjacent consecutive letters
		for {
			n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
			if err != nil {
				return "", err
			}

			next := charset[n.Int64()]
			prev := password[i-1]

			diff := int(next) - int(prev)
			if diff != 1 && diff != -1 {
				password[i] = next
				break
			}
		}
	}

	return string(password), nil
}
