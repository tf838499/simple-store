package customer

import (
	"context"
	"database/sql"
	"simple-store/internal/adapter/redisclient"
	"simple-store/internal/adapter/repository/PostgresDB"
)

//go:generate mockgen -destination automock/good_repository.go -package=automock . GoodRepository
type CartRepository interface {
	SetGood(ctx context.Context, arg []redisclient.GoodInCartParams) error
	DeleteGood(ctx context.Context, arg []redisclient.GoodInCartParams) error
	GetCartList(ctx context.Context, arg string) ([]redisclient.GoodInRedisParams, error)
	GetGoodPrice(ctx context.Context, arg []string) (redisclient.GoodPrice, error)
	SetGoodPrice(ctx context.Context, arg redisclient.GoodPriceInfo) error
}

type OrderRepository interface {
	InsertOrder(ctx context.Context, arg PostgresDB.InsertOrderParams) error
	GetGoodByName(ctx context.Context, imageName sql.NullString) (PostgresDB.Good, error)
}
