package app

import (
	"github.com/arkark/dcp/internal/version"
	"github.com/urfave/cli/v2"
)

func Build() *cli.App {
	cli.AppHelpTemplate = getHelpTemplate()

	app := &cli.App{
		Name:     "dcp",
		Usage:    "An alias of `docker container cp` and useful completions",
		Action:   getAction(),
		Flags:    getFlags(),
		Version:  version.VERSION,
		HideHelp: true, // Disable built-in help

		EnableBashCompletion: true,
		BashComplete:         getBashComplete(),
	}

	return app
}

func getHelpTemplate() string {
	return `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   dcp [OPTIONS] CONTAINER:SRC_PATH DEST_PATH|-
   dcp [OPTIONS] SRC_PATH|- CONTAINER:DEST_PATH

   Use '-' as the source to read a tar archive from stdin and extract it to a directory destination in a container.
   Use '-' as the destination to stream a tar archive of a container source to stdout.

OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}
`
}
