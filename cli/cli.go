package cli

import (
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/xozrc/discovery/version"
)

var (
	clusterStore = "127.0.0.1:2379,127.0.0.1:4001"
	prefix       = "xozrc_discovery"
	certfile     = ""
	keyfile      = ""
	cafile       = ""
)

func Start() {
	app := cli.NewApp()
	app.Name = "disocvery"
	app.Usage = "disocvery [global options] command [command options] [arguments...]."
	app.Version = version.VERSION + " (" + version.GITCOMMIT + ")"

	app.Author = ""
	app.Email = ""

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "cluster-store,C",
			Value:       clusterStore,
			Usage:       " a comma-delimited list of machine addresses in the cluster",
			Destination: &clusterStore,
		},
		cli.StringFlag{
			Name:        "cert-file",
			Value:       certfile,
			Usage:       "identify HTTPS client using this SSL certificate file",
			Destination: &certfile,
		},
		cli.StringFlag{
			Name:        "key-file",
			Value:       keyfile,
			Usage:       "identify HTTPS client using this SSL key file",
			Destination: &keyfile,
		},
		cli.StringFlag{
			Name:        "ca-file",
			Value:       cafile,
			Usage:       "verify certificates of HTTPS-enabled servers using this CA bundle",
			Destination: &cafile,
		},

		cli.StringFlag{
			Name:        "prefix,p",
			Value:       prefix,
			Usage:       "prefix difference vendor",
			Destination: &prefix,
		},
	}

	app.Commands = commands
	app.Run(os.Args)
}

func clusterAddrs() []string {
	return strings.Split(clusterStore, ",")
}
