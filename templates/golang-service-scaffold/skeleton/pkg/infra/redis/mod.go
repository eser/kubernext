package redis

import (
	"context"

	"github.com/eser/go-service/pkg/infra/config"
	"github.com/redis/go-redis/v9"
)

type (
	RedisClient = redis.Client
)

func NewRedisClient(conf *config.Config) (*RedisClient, error) {
	if conf.RedisAddr == "" {
		return nil, nil
	}

	redisOptions, err := redis.ParseURL(conf.RedisAddr)
	if err != nil {
		return nil, err
	}

	if conf.RedisPwd != "" {
		redisOptions.Password = conf.RedisPwd
	}

	redisClient := redis.NewClient(redisOptions)

	if conf.RedisConnCheck {
		_, err = redisClient.Ping(context.Background()).Result()
		if err != nil {
			return nil, err
		}
	}

	return redisClient, nil
}
