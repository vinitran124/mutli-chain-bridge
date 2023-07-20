package database

import (
	"github.com/redis/go-redis/v9"
	"os"
)

const (
	envRDDataSourceAddress  = "RD_ADDRESS"
	envRDDataSourcePassword = "RD_PASSWORD"
)

type RedisRepository struct {
	Client *redis.Client
}

func (r *RedisRepository) Init() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv(envRDDataSourceAddress),
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
