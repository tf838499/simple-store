package customer

import (
	"context"
	"database/sql"
	"simple-store/internal/adapter/redisclient"
	"simple-store/internal/adapter/repository/PostgresDB"
)

//go:generate mockgen -destination automock/cart_repository.go -package=automock . CartRepository
type CartRepository interface {
	SetGood(ctx context.Context, arg []redisclient.GoodInCartParams) error
	DeleteGood(ctx context.Context, arg []redisclient.GoodInCartParams) error
	GetCartListCache(ctx context.Context, arg string) ([]redisclient.GoodInCartParams, error)
	GetGoodPrice(ctx context.Context, arg []string) ([]int, error)
	SetGoodPrice(ctx context.Context, arg redisclient.GoodPriceInfo) error
	MSetGoodPrice(ctx context.Context, arg []redisclient.GoodInCartParams) error
}

//go:generate mockgen -destination automock/order_repository.go -package=automock . OrderRepository
type OrderRepository interface {
	InsertOrder(ctx context.Context, arg PostgresDB.InsertOrderParams) error
	GetGetOrderByOwner(ctx context.Context, owner sql.NullString) ([]PostgresDB.Order, error)
	GetGoodByName(ctx context.Context, imageName sql.NullString) (PostgresDB.Good, error)
}

type AuthRepository interface {
	SetUserInfo(ctx context.Context, arg redisclient.OauthInfo) error
	GetUserInfo(ctx context.Context, arg redisclient.OauthInfo) (string, error)
}
