package app

import (
	"github.com/ArkArk/dcp/internal/comp"
	"github.com/urfave/cli"
)

func getFlags() []cli.Flag {
	cli.VersionFlag = cli.BoolFlag{
		Name:  "v, version", // not "version, v"
		Usage: "Print the version",
	}

	appFlags := []cli.Flag{
		cli.BoolFlag{
			Name:  "a, archive",
			Usage: "Archive mode (copy all uid/gid information)",
		},
		cli.BoolFlag{
			Name:  "L, follow-link",
			Usage: "Always follow symbol link in SRC_PATH",
		},
		cli.BoolFlag{
			Name:  "h, help",
			Usage: "Print this help message",
		},
	}
	return append(appFlags, comp.GetFlags()...)
}
