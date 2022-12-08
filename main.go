package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"test_ticket/common"
	"test_ticket/common/database"
	"test_ticket/web_app"
	"test_ticket/worker"
)

func main() {
	config := common.InitConf()

	var executableName string
	app := &cli.App{
		Name:  "test ticket",
		Usage: "test ticket executable for workers and web api",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "run",
				Aliases:     []string{"r"},
				Usage:       "Run executable. Accepted names: runWorkers, runWebApp and runMigrations",
				Destination: &executableName,
			},
		},
		Action: func(c *cli.Context) error {
			if executableName == "runWorkers" {
				worker.RunWorker(config)
			} else if executableName == "runWebApp" {
				web_app.RunServer(config)
			} else if executableName == "runMigrations" {
				database.RunMigration()
			} else {
				return fmt.Errorf("unknown executable %s", executableName)
			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
