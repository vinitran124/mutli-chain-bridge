package main

import (
	"bridge/app/content/datastore"
	"bridge/config"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
)

func beforeStartBlockchain(c *cli.Context) error {
	cfg, err := config.Load(c)
	if err != nil {
		return err
	}
	SetContextSQL(cfg.Database)
	SetContextRedisClient(cfg.Redis)
	SetChainClient()
	return nil
}

func startBlockchain(c *cli.Context) error {
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
				log.Println("blockchain: ", err)
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

	for {
	}
}
