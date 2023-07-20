package crawler

import (
	"bridge/app/utils"
	"github.com/urfave/cli/v2"
)

func NewCommand() *cli.Command {
	return &cli.Command{
		Name:  "crawler",
		Usage: "Start crawler server",
		Flags: []cli.Flag{},

		Action: func(c *cli.Context) error {
			return startBlockchainServer(c)
		},

		Before: func(c *cli.Context) error {
			return beforeStartBlockchainServer(c)
		},
	}
}

func beforeStartBlockchainServer(c *cli.Context) error {
	utils.SetContextSQL()
	utils.SetContextRedisClient()
	return nil
}

func startBlockchainServer(c *cli.Context) error {
	return nil

}
