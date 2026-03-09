package uid

import (
	"context"
	"crypto/rand"
	"fmt"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"github.com/urfave/cli/v3"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func BuildCommands(baseCmds []*cli.Command) []*cli.Command {
	return append(baseCmds, commands...)
}

var commands = []*cli.Command{{
	Name: "gen",
	Commands: []*cli.Command{
		{
			Name: "uuid4",
			Action: func(_ context.Context, _ *cli.Command) error {
				fmt.Println(uuid.New().String())
				return nil
			},
		},
		{
			Name: "uuid7",
			Action: func(_ context.Context, _ *cli.Command) error {
				fmt.Println(uuid.Must(uuid.NewV7()).String())
				return nil
			},
		},
		{
			Name: "ulid",
			Action: func(_ context.Context, _ *cli.Command) error {
				entropy := ulid.Monotonic(rand.Reader, 0)
				fmt.Println(ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String())
				return nil
			},
		},
		{
			Name: "snowflake",
			Action: func(_ context.Context, _ *cli.Command) error {
				node, err := snowflake.NewNode(1)
				if err != nil {
					return fmt.Errorf("new node: %w", err)
				}
				fmt.Println(node.Generate().Int64())

				return nil
			},
		},
		{
			Name: "oid",
			Action: func(_ context.Context, _ *cli.Command) error {
				oid := bson.NewObjectID()
				fmt.Println(oid.Hex())

				return nil
			},
		},
	},
}}
