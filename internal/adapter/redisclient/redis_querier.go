package redisclient

import "context"

type RedisQuerier interface {
	SetGood(ctx context.Context, arg []GoodInCartParams) error
	DeleteGood(ctx context.Context, arg []GoodInCartParams) error
	GetCartListCache(ctx context.Context, arg string) ([]GoodInCartParams, error)
	GetGoodPrice(ctx context.Context, arg []string) ([]int, error)
	SetGoodPrice(ctx context.Context, arg GoodPriceInfo) error
	MSetGoodPrice(ctx context.Context, arg []GoodInCartParams) error
}

const (
	prefixCustomer = "Customer:Cart:"
	prefixPrice    = "Good:Price:"

	// store    = "store:"
	// good     = "good:"
)

var _ RedisQuerier = (*RedisRepository)(nil)
