package main

import (
	"bridge/app/blockchain"
	"bridge/app/blockchain/chainFactory"
	"bridge/db"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/stephenafamo/bob"
	"github.com/uptrace/bun"
	"log"
	"os"
	"strings"
)

var ctx = context.Background()

const (
	contextSQLRepository = "CONTEXT_SQL_REPOSITORY"
	contextRedisClient   = "CONTEXT_REDIS_CLIENT"
	contextChainClient   = "CONTEXT_CHAIN_CLIENT_"
)

func SQLRepository() bob.DB {
	return bob.NewDB(GetContextSQL().DB)
}

func RedisRepository() *redis.Client {
	return GetContextRedisClient()
}

func ChainRepository(chainId string) chainFactory.IChain {
	return GetContextChainClient(chainId)
}

func GetContextSQL() *bun.DB {
	client, _ := ctx.Value(contextSQLRepository).(*bun.DB)
	return client
}

func GetContextRedisClient() *redis.Client {
	client, _ := ctx.Value(contextRedisClient).(*redis.Client)
	return client
}

func SetContextRedisClient(cfg db.RedisConfig) {
	redisClient := db.NewRedis(cfg)
	ctx = context.WithValue(ctx, contextRedisClient, redisClient)
	log.Println("redis", redisClient.Ping(ctx).String())
}

func SetContextSQL(cfg db.DatabaseConfig) {
	client := db.NewSQLDB(cfg)
	ctx = context.WithValue(ctx, contextSQLRepository, client)
	err := client.Ping()
	if err != nil {
		log.Fatal("Set context SQL error", err)
		return
	}
	log.Println("postgres: connected to database")
}

func SetChainClient() {
	chainList := strings.Split(os.Getenv(envChainIdList), ".")
	for _, chainId := range chainList {
		chain, err := blockchain.NewChain(chainId)
		log.Println("init chain id :", chainId)
		if err != nil {
			fmt.Println("Set context Chain error", err)
			return
		}
		ctx = context.WithValue(ctx, fmt.Sprintf("%s%s", contextChainClient, chainId), chain)
	}
}

func GetContextChainClient(chainId string) chainFactory.IChain {
	client, _ := ctx.Value(fmt.Sprintf("%s%s", contextChainClient, chainId)).(chainFactory.IChain)
	return client
}
