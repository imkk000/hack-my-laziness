package time

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v3"
)

const zoneDir = "/usr/share/zoneinfo/"

var getTimeZone = &cli.Command{
	Name:  "zone",
	Usage: "Zone",
	Action: func(_ context.Context, _ *cli.Command) error {
		var zones []string
		err := filepath.Walk(zoneDir, func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() {
				return err
			}
			if isZoneNameSkipper(path) {
				return nil
			}

			zone := strings.TrimPrefix(path, zoneDir)
			zones = append(zones, zone)
			return nil
		})
		if err != nil {
			return fmt.Errorf("walk timezone directory: %w", err)
		}

		for _, zone := range zones {
			fmt.Println(zone)
		}

		return nil
	},
}

func isZoneNameSkipper(path string) bool {
	return strings.HasSuffix(path, ".tab") ||
		strings.HasSuffix(path, ".list") ||
		strings.Contains(path, "posix") ||
		strings.Contains(path, "right") ||
		strings.Contains(path, ".zi")
}
