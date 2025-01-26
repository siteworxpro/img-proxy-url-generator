package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/siteworxpro/img-proxy-url-generator/config"
	"strconv"
)

type Redis struct {
	initialized redisStatus
	client      *redis.Client
}

type redisStatus uint8

const (
	redisStatusUninitialized redisStatus = iota
	redisStatusInitialized
)

var singleton *Redis

func New(config *config.Config) (*Redis, error) {
	if singleton != nil && singleton.initialized == redisStatusUninitialized {
		return singleton, nil
	}

	db, err := strconv.ParseInt(config.Redis.DB, 10, 64)
	if err != nil {
		db = 0
	}

	port := config.Redis.Port
	if port == "" {
		port = "6379"
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Redis.Host, port),
		DB:       int(db),
		Password: config.Redis.Password,
	})

	_, err = rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	singleton = &Redis{
		initialized: redisStatusInitialized,
		client:      rdb,
	}

	return singleton, nil
}

func (r *Redis) GetClient() *redis.Client {
	return r.client
}

func (r *Redis) Close() error {
	return r.client.Close()
}
