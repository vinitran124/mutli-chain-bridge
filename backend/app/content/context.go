package content

import (
	"bridge/app/blockchain/chainFactory"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/stephenafamo/bob"
	"github.com/uptrace/bun"
)

const (
	contextSQLRepository = "CONTEXT_SQL_REPOSITORY"
	contextRedisClient   = "CONTEXT_REDIS_CLIENT"
	contextChainClient   = "CONTEXT_CHAIN_CLIENT_"
)

var ctx = context.Background()

func SQLRepository() bob.DB {
	client, _ := ctx.Value(contextSQLRepository).(*bun.DB)
	return bob.NewDB(client.DB)
}

func RedisRepository() *redis.Client {
	client, _ := ctx.Value(contextRedisClient).(*redis.Client)
	return client
}

func ChainRepository(chainId string) chainFactory.IChain {
	client, _ := ctx.Value(fmt.Sprintf("%s%s", contextChainClient, chainId)).(chainFactory.IChain)
	return client
}
