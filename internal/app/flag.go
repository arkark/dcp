package app

import (
	"github.com/arkark/dcp/internal/comp"
	"github.com/urfave/cli/v2"
)

func getFlags() []cli.Flag {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Print the version",
	}

	appFlags := []cli.Flag{
		&cli.BoolFlag{
			Name:    "archive",
			Aliases: []string{"a"},
			Usage:   "Archive mode (copy all uid/gid information)",
		},
		&cli.BoolFlag{
			Name:    "follow-link",
			Aliases: []string{"L"},
			Usage:   "Always follow symbol link in SRC_PATH",
		},
		&cli.BoolFlag{
			Name:    "help",
			Aliases: []string{"h"},
			Usage:   "Print this help message",
		},
	}
	return append(appFlags, comp.GetFlags()...)
}
