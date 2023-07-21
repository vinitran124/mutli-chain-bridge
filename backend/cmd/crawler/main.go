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
	envChainIdList = "CHAIN_ID_LIST"
	envPath        = ".env"
)

func init() {
	if err := godotenv.Overload(strings.Split(envPath, ",")...); err != nil {
		fmt.Println("Load env error", err.Error())
	}
}

func main() {
	utils.SetContextSQL()
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
			_, err := tx.InsertTransaction(context.Background(), SQLRepository(), &event)
			if err != nil {
				log.Println(err)
			}
		default:
			time.Sleep(time.Second)
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
