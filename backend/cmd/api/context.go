package api

import (
	"bridge/app/blockchain/chainFactory"
	"bridge/app/utils"
	"github.com/redis/go-redis/v9"
	"github.com/uptrace/bun"
)

func SQLRepository() *bun.DB {
	return utils.GetContextSQL().Client
}

func RedisRepository() *redis.Client {
	return utils.GetContextRedisClient().Client
}

func ChainRepository(chainId string) chainFactory.IChain {
	return utils.GetContextChainClient(chainId)
}
