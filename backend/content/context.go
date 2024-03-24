package content

import (
	"bridge/app/blockchain/chainFactory"
	"bridge/config"
	"bridge/context"
	"github.com/redis/go-redis/v9"
	"github.com/stephenafamo/bob"
)

func SQLRepository() bob.DB {
	return bob.NewDB(context.GetContextSQL().DB)
}

func RedisRepository() *redis.Client {
	return context.GetContextRedisClient()
}

func ChainRepository(chainId string) chainFactory.IChain {
	return context.GetContextChainClient(chainId)
}

func ConfigRepository() *config.Config {
	return context.GetContextConfig()
}
