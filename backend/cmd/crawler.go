package main

import (
	"bridge/app/content/bob"
	"bridge/app/content/datastore"
	"bridge/app/utils"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strings"
	"time"
)

const (
	envChainIdList       = "CHAIN_ID_LIST"
	redisNewDepositEvent = "NEW_DEPOSIT_EVENT"
)

func beforeStartCrawler(c *cli.Context) error {
	utils.SetContextSQL()
	utils.SetContextRedisClient()
	utils.SetChainClient()
	return nil
}

func startCrawler(c *cli.Context) error {
	events := make(chan bob.Transaction)
	chainList := strings.Split(os.Getenv(envChainIdList), ".")

	for _, chainId := range chainList {
		go func(chainId string) {
			err := crawl(chainId, events)
			if err != nil {
				log.Println(err)
			}
		}(chainId)
	}

	for {
		select {
		case event := <-events:
			tx := datastore.DatastoreTransaction{}
			transaction, err := tx.Insert(ctx, SQLRepository(), &event)
			if err != nil {
				log.Println(err)
				continue
			}

			err = RedisRepository().Publish(ctx, redisNewDepositEvent, transaction.ID.String()).Err()
			if err != nil {
				return err
			}

		default:
			time.Sleep(10 * time.Second)
		}
	}
}

func crawl(chainId string, events chan bob.Transaction) error {
	chain := ChainRepository(chainId)
	err := chain.TrackDeposit(events)
	if err != nil {
		return err
	}
	return err
}
