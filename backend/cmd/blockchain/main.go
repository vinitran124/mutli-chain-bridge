package main

import (
	"bridge/app/content/datastore"
	"bridge/app/utils"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"strings"
)

const (
	redisNewDepositEvent = "NEW_DEPOSIT_EVENT"
	redisNewRequest      = "NEW_BRIDGE_REQUEST"
	envPath              = ".env,.env.local"
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

	eventSub := RedisRepository().Subscribe(ctx, redisNewDepositEvent)
	defer eventSub.Close()

	eventCh := eventSub.Channel()
	bridgeStr := datastore.DatastoreBridgeRequest{}
	transactionStr := datastore.DatastoreTransaction{}
	tokenStr := datastore.DatastoreToken{}

	go func() {
		for msg := range eventCh {
			bridgeRq, err := bridgeStr.CheckValidForWithdrawal(ctx, SQLRepository(), msg.Payload)
			if err != nil {
				log.Println(err)
				continue
			}

			if bridgeRq == nil {
				log.Println(fmt.Sprintf("event id %s: not valid with request", msg.Payload))
				continue
			}

			token, err := tokenStr.FindTokenInOutputChain(ctx, SQLRepository(), bridgeRq.Token, bridgeRq.InputChain, bridgeRq.OutputChain)
			if err != nil {
				log.Println(fmt.Sprintf("event id %s: %e", msg.Payload, err))
				continue
			}

			err = ChainRepository(token.ChainID).CallWithdrawal(token.Address, bridgeRq.UserAddress, bridgeRq.RawAmount)
			if err != nil {
				log.Println(fmt.Sprintf("event id %s: %e", msg.Payload, err))
				continue
			}

			err = bridgeStr.SetComplete(ctx, SQLRepository(), bridgeRq.ID.String())
			if err != nil {
				log.Println(fmt.Sprintf("event id %s: %e", msg.Payload, err))
				continue
			}

			err = transactionStr.SetComplete(ctx, SQLRepository(), msg.Payload)
			if err != nil {
				log.Println(fmt.Sprintf("event id %s: %e", msg.Payload, err))
				continue
			}
		}
	}()

	//go func() {
	//	for msg := range requestCh {
	//		requestId = append(requestId, msg.Payload)
	//	}
	//}()

	for {

	}
}
