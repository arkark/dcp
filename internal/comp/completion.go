package comp

import (
	"github.com/arkark/dcp/internal/logger"
	"github.com/urfave/cli/v2"
)

type Completion struct {
	Line    string
	Point   int
	Current string
}

func GetFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:   "completion-line",
			Hidden: true,
		},
		&cli.IntFlag{
			Name:   "completion-point",
			Hidden: true,
		},
		&cli.StringFlag{
			Name:   "completion-current",
			Hidden: true,
		},
	}
}

func GetCompletion(c *cli.Context) Completion {
	flags := Completion{
		Line:    c.String("completion-line"),
		Point:   c.Int("completion-point"),
		Current: c.String("completion-current"),
	}
	logger.Write(logger.DEBUG, "Completion: %#v", flags)
	return flags
}
