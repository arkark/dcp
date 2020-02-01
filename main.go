package main

import (
	"log"
	"os"

	"github.com/arkark/dcp/internal/app"
	"github.com/arkark/dcp/internal/logger"
)

func main() {
	if err := logger.Init("debug/dcp.log", logger.DEBUG); err != nil {
		log.Fatal(err)
	}

	if err := app.Build().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
