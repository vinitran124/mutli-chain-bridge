package main

import (
	"bridge/app/content/bob"
	"bridge/app/content/datastore"
	"bridge/app/utils"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
	"time"
)

const (
	envChainIdList       = "CHAIN_ID_LIST"
	envPath              = ".env,.env.local"
	redisNewDepositEvent = "NEW_DEPOSIT_EVENT"
)

func init() {
	if err := godotenv.Overload(strings.Split(envPath, ",")...); err != nil {
		fmt.Println("Load env error", err.Error())
	}
}

func main() {
	ctx := context.Background()
	utils.SetContextSQL()
	utils.SetContextRedisClient()
	utils.SetChainClient()

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
				return
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
