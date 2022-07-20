package main

import (
	"github.com/gancc6/ksubdomain/core/conf"
	"github.com/gancc6/ksubdomain/core/gologger"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := &cli.App{
		Name:    conf.AppName,
		Version: conf.Version,
		Usage:   conf.Description,
		Commands: []*cli.Command{
			enumCommand,
			verifyCommand,
			testCommand,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		gologger.Fatalf(err.Error())
	}
}
