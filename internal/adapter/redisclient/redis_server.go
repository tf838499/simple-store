package redisclient

import (
	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	Client *redis.Client
}

func NewRedisRepository(client *redis.Client) *RedisRepository {

	return &RedisRepository{
		Client: client,
	}
}
