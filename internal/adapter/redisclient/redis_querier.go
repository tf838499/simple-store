package redisclient

import "context"

type RedisQuerier interface {
	SetGood(ctx context.Context, arg []GoodInCartParams) error
	DeleteGood(ctx context.Context, arg []GoodInCartParams) error
	GetCartList(ctx context.Context, arg string) ([]GoodInRedisParams, error)
	GetGoodPrice(ctx context.Context, arg []string) ([]int, error)
	SetGoodPrice(ctx context.Context, arg GoodPriceInfo) error
}

const (
	prefixCustomer = "Customer:Cart:"
	prefixPrice    = "Good:Price:"

	// store    = "store:"
	// good     = "good:"
)

var _ RedisQuerier = (*RedisRepository)(nil)
