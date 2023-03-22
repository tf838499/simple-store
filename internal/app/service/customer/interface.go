package customer

import (
	"context"
	"database/sql"
	"simple-store/internal/adapter/RedisClient"
	"simple-store/internal/adapter/repository/PostgresDB"
)

//go:generate mockgen -destination automock/good_repository.go -package=automock . GoodRepository
type CartRepository interface {
	SetGood(ctx context.Context, arg []RedisClient.GoodInCartParams) error
	DeleteGood(ctx context.Context, arg []RedisClient.GoodInCartParams) error
	GetCartList(ctx context.Context, arg string) ([]RedisClient.GoodInRedisParams, error)
	GetGoodPrice(ctx context.Context, arg []string) (RedisClient.GoodPrice, error)
	SetGoodPrice(ctx context.Context, arg RedisClient.GoodPriceInfo) error
}

type OrderRepository interface {
	InsertOrder(ctx context.Context, arg PostgresDB.InsertOrderParams) error
	GetGoodByName(ctx context.Context, imageName sql.NullString) (PostgresDB.Good, error)
}
