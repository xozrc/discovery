package ctl

import (
	"os"
)
import (
	"github.com/codegangsta/cli"
)

func Start() {
	app := cli.NewApp()
	app.Name = "discoveryctl"
	app.Usage = "A simple command line client for discovery."
	app.Run(os.Args)

	app.Flags = []cli.Flag{}

	app.Commands = []cli.Command{}
}
