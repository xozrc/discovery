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
	heartbeat = time.Second
)
var (
	watchCmd = cli.Command{
		Name:        "watch",
		Aliases:     []string{"w"},
		Usage:       "watch --heartbeat=${heartbeat(second)}",
		Description: "watch specific service",
		Action:      watch,
		Flags: []cli.Flag{
			cli.DurationFlag{
				Name:        "heartbeat",
				Value:       heartbeat,
				Usage:       "error heartbeat",
				Destination: &heartbeat,
			},
		},
	}
)

func init() {
	appendCmd(watchCmd)
}

func watch(c *cli.Context) {
	//args no enough
	if c.NArg() == 0 {
		log.Fatalln("need service name arg")
	}

	serv := c.Args()[0]

	st, err := discovery.NewEtcdStore(cafile, certfile, keyfile, clusterAddrs())
	if err != nil {
		log.Fatalln(err)
	}

	bke := discovery.NewBackend(st, types.EntryFactoryInstance, prefix, serv, heartbeat)

	entriesCh, errCh, err := bke.Watch()
	if err != nil {
		log.Fatalln(err)
	}

	for {
		select {
		case entries := <-entriesCh:
			{
				log.Println("get change:")
				if len(entries) == 0 {
					log.Println("no record")
					continue
				}

				for i, entry := range entries {
					mc, err := entry.Marshal()
					if err != nil {
						break
					}
					log.Printf("%d:%v\n", i, string(mc))
				}
			}
		case err = <-errCh:
			{
				log.Printf("get error:%v\n", err.Error())
			}
		}
	}

}
