package main

import (
	"bridge/app/blockchain/chainFactory"
	"bridge/app/utils"
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/stephenafamo/bob"
)

var ctx = context.Background()

func SQLRepository() bob.DB {
	return bob.NewDB(utils.GetContextSQL().Client.DB)
}

func RedisRepository() *redis.Client {
	return utils.GetContextRedisClient().Client
}

func ChainRepository(chainId string) chainFactory.IChain {
	return utils.GetContextChainClient(chainId)
}
