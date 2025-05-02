package list

import (
	"os"
	"strings"

	"github.com/apex/log"
	"github.com/urfave/cli/v2"

	"github.com/glamorousis/distillery/pkg/common"
	"github.com/glamorousis/distillery/pkg/config"
	"github.com/glamorousis/distillery/pkg/inventory"
)

func Execute(c *cli.Context) error {
	cfg, err := config.New(c.String("config"))
	if err != nil {
		return err
	}

	inv := inventory.New(os.DirFS(cfg.GetPath()), cfg.GetPath(), cfg.GetOptPath(), cfg)

	for _, key := range inv.GetBinsSortedKeys() {
		bin := inv.Bins[key]
		log.Infof("%s (versions: %s)", key, strings.Join(bin.ListVersions(), ", "))
	}

	return nil
}

func init() {
	cmd := &cli.Command{
		Name:        "list",
		Usage:       "list installed binaries and versions",
		Description: `list installed binaries and versions`,
		Before:      common.Before,
		Flags:       common.Flags(),
		Action:      Execute,
	}

	common.RegisterCommand(cmd)
}
