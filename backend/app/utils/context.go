package utils

import (
	"bridge/app/blockchain"
	"bridge/app/blockchain/chainFactory"
	"bridge/app/database"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	contextSQLRepository = "CONTEXT_SQL_REPOSITORY"
	contextRedisClient   = "CONTEXT_REDIS_CLIENT"
	contextChainClient   = "CONTEXT_CHAIN_CLIENT_"
	envChainIdList       = "CHAIN_ID_LIST"
)

var (
	background = context.Background()
)

func GetContextSQL() *database.SQLRepository {
	client, _ := background.Value(contextSQLRepository).(*database.SQLRepository)
	return client
}

func GetContextRedisClient() *database.RedisRepository {
	client, _ := background.Value(contextRedisClient).(*database.RedisRepository)
	return client
}

func SetContextRedisClient() {
	redisClient, err := database.NewRedisRepository()
	if err != nil {
		return
	}

	background = context.WithValue(background, contextRedisClient, redisClient)
	log.Println("redis", redisClient.Client.Ping(context.Background()).String())
}

func SetContextSQL() {
	client, err := database.NewRepository()
	if err != nil {
		fmt.Println("Set context SQL error", err)
		return
	}
	background = context.WithValue(background, contextSQLRepository, client)
	log.Println("connected to database.")
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
		background = context.WithValue(background, fmt.Sprintf("%s%s", contextChainClient, chainId), chain)
	}
}

func GetContextChainClient(chainId string) chainFactory.IChain {
	client, _ := background.Value(fmt.Sprintf("%s%s", contextChainClient, chainId)).(chainFactory.IChain)
	return client
}
