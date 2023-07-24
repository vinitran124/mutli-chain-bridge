package database

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

const (
	envRDDataSourceHost     = "REDIS_HOST"
	envRDDataSourcePort     = "REDIS_PORT"
	envRDDataSourcePassword = "RD_PASSWORD"
)

type RedisRepository struct {
	Client *redis.Client
}

func (r *RedisRepository) Init() error {
	redisAddr := fmt.Sprintf("%s:%s", os.Getenv(envRDDataSourceHost), os.Getenv(envRDDataSourcePort))
	log.Println("redis: ", redisAddr)
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: os.Getenv(envRDDataSourcePassword), // no password set
		DB:       0,                                  // use default DB
	})

	r.Client = rdb
	return nil
}

func NewRedisRepository() (*RedisRepository, error) {
	r := new(RedisRepository)
	err := r.Init()
	if err != nil {
		return nil, err
	}

	return r, nil
}
