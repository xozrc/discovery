package cli

import (
	"log"
	"time"
)
import (
	"github.com/codegangsta/cli"
	"github.com/xozrc/discovery/discovery"
	"github.com/xozrc/discovery/types"
)

var (
	ttl = time.Minute
)
var (
	registerCmd = cli.Command{
		Name:        "register",
		Usage:       "register --ttl=${ttl(second)} srvName val",
		Description: "register specific service",
		Action:      register,
		Flags: []cli.Flag{
			cli.DurationFlag{
				Name:        "ttl",
				Value:       ttl,
				Usage:       "time to live in store",
				Destination: &heartbeat,
			},
		},
	}
)

func init() {
	appendCmd(registerCmd)
}

func register(c *cli.Context) {
	//args no enough
	if c.NArg() <= 1 {
		log.Fatalln("need service name arg")
	}

	serv := c.Args()[0]
	servVal := c.Args()[1]

	st, err := discovery.NewEtcdStore(cafile, certfile, keyfile, clusterAddrs())
	if err != nil {
		log.Fatalln(err)
	}

	bke := discovery.NewBackend(st, types.EntryFactoryInstance, prefix, serv, heartbeat)

	tmpttl := time.Duration(int64(0))

	for {
		select {
		case <-time.After(tmpttl):
			{
				se := types.StringEntry(servVal)
				err = bke.Register(&se, ttl)
				if err != nil {
					log.Fatalln(err)
				}
				log.Printf("register service %s,val %s\n", serv, servVal)
				tmpttl = ttl
			}
		}
	}

}
