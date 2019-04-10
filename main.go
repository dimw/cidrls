package main

import (
	"fmt"
	"github.com/dimw/cidrls/cidr"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Version = "0.1.0"
	app.Name = "cidrls"
	app.Usage = "List IPs based on CIDR"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "clean, c",
			Usage: "remove network and broadcast IPs",
		},
	}
	app.Action = func(c *cli.Context) error {
		for argumentIndex := 0; argumentIndex < c.NArg(); argumentIndex++ {
			hostname := c.Args().Get(argumentIndex)

			removeNetworkAndBroadcastIps := c.Bool("clean")
			ips := cidr.CalculateIps(hostname, removeNetworkAndBroadcastIps)
			for _, ip := range ips {
				fmt.Println(ip)
			}
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
