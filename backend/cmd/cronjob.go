package main

import (
	"bridge/app/content/bob"
	"bridge/app/utils"
	"context"
	"github.com/urfave/cli/v2"
	"log"
	"time"
)

func beforeStartCronjob(c *cli.Context) error {
	utils.SetContextSQL()
	return nil
}

func startCronjob(c *cli.Context) error {
	for {
		expiredRequest, err := bob.BridgeRequests(context.Background(),
			SQLRepository(),
			bob.SelectWhere.BridgeRequests.CreatedAt.LTE(time.Now().Add(-10*time.Minute)),
		).All()
		if err != nil {
			log.Println(err)
			continue
		}

		_, err = expiredRequest.DeleteAll(context.Background(), SQLRepository())
		if err != nil {
			log.Println(err)
			continue
		}
		time.Sleep(10 * time.Minute)
	}
}
