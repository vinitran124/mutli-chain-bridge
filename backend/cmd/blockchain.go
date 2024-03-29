package main

import (
	"fmt"
	"log"

	"bridge/config"
	"bridge/content/datastore"
	"bridge/context"

	"github.com/urfave/cli/v2"
)

func beforeStartBlockchain(c *cli.Context) error {
	cfg, err := config.Load(c)
	if err != nil {
		return err
	}
	context.SetContextConfig(cfg)
	context.SetContextSQL(cfg.Database)
	context.SetContextRedisClient(cfg.Redis)
	context.SetChainClient()
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
