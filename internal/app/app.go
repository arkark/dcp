package app

import (
	"github.com/ArkArk/dcp/internal/version"
	"github.com/urfave/cli"
)

func Build() *cli.App {
	app := cli.NewApp()

	app.Version = version.VERSION

	app.HideHelp = true // Disable built-in help
	app.EnableBashCompletion = true

	app.Name = "dcp"
	app.Usage = "An alias of `docker container cp` and useful completions"
	app.UsageText = "dcp [OPTIONS] CONTAINER:SRC_PATH DEST_PATH|-\n" +
		"   dcp [OPTIONS] SRC_PATH|- CONTAINER:DEST_PATH"

	app.Commands = []cli.Command{}
	app.Flags = getFlags()
	app.Action = getAction()
	app.BashComplete = getBashComplete()

	return app
}
