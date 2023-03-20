package customer

import (
	"context"
	"simple-store/internal/adapter/RedisClient"
)

//go:generate mockgen -destination automock/good_repository.go -package=automock . GoodRepository
type CartRepository interface {
	SetGood(ctx context.Context, arg []RedisClient.GoodInCartParams) error
	DeleteGood(ctx context.Context, arg []RedisClient.GoodInCartParams) error
	GetCartList(ctx context.Context, arg RedisClient.GoodInCartParams) ([]RedisClient.GoodInRedisParams, error)
}
