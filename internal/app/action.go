package app

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/ArkArk/dcp/internal/logger"
	"github.com/urfave/cli/v2"
)

func getAction() cli.ActionFunc {
	return func(c *cli.Context) error {
		if c.Bool("help") {
			return cli.ShowAppHelp(c)
		}

		args := []string{"container", "cp"}
		args = append(args, os.Args[1:]...)
		logger.Write(
			logger.DEBUG,
			"Exec: %s %v",
			"docker",
			args,
		)
		out, err := exec.Command("docker", args...).CombinedOutput()

		message := string(out)
		message = strings.Replace(message, "docker container cp", "dcp", -1)
		message = strings.Replace(message, "docker cp", "dcp", -1)
		if err != nil {
			fmt.Fprint(os.Stderr, message)
			os.Exit(1)
		}
		fmt.Fprint(os.Stdout, message)
		return nil
	}
}
