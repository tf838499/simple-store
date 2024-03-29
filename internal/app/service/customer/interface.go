package customer

import (
	"context"
	"database/sql"
	"simple-store/internal/adapter/repository/PostgresDB"
	"simple-store/internal/adapter/repository/RedisCache"
)

//go:generate mockgen -destination automock/cart_repository.go -package=automock . CartRepository
type CartRepository interface {
	SetGood(ctx context.Context, arg []RedisCache.GoodInCartParams) error
	DeleteGood(ctx context.Context, arg []RedisCache.GoodInCartParams) error
	GetCartListCache(ctx context.Context, arg string) ([]RedisCache.GoodInCartParams, error)
	GetGoodPrice(ctx context.Context, arg []string) ([]int, error)
	SetGoodPrice(ctx context.Context, arg RedisCache.GoodPriceInfo) error
	MSetGoodPrice(ctx context.Context, arg []RedisCache.GoodInCartParams) error
}

//go:generate mockgen -destination automock/order_repository.go -package=automock . OrderRepository
type OrderRepository interface {
	InsertOrder(ctx context.Context, arg PostgresDB.InsertOrderParams) error
	GetGetOrderByOwner(ctx context.Context, owner sql.NullString) ([]PostgresDB.Order, error)
	GetGoodByName(ctx context.Context, imageName sql.NullString) (PostgresDB.Good, error)
}

type AuthRepository interface {
	SetUserInfo(ctx context.Context, arg RedisCache.OauthInfo) error
	GetUserInfo(ctx context.Context, arg RedisCache.OauthInfo) (string, error)
}
