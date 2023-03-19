package redis_sever

import (
	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	Client *redis.Client
}
type RedisQueriers struct {
}

func NewPostgresRepository(client *redis.Client) *RedisRepository {

	return &RedisRepository{
		Client: client,
	}
}
