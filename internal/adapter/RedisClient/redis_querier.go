package RedisClient

import "context"

type RedisQuerier interface {
	SetGood(ctx context.Context, arg []GoodInCartParams) error
	DeleteGood(ctx context.Context, arg []GoodInCartParams) error
	GetCartList(ctx context.Context, arg GoodInCartParams) ([]GoodInRedisParams, error)
}

const (
	//cart
	customer = "Customer:Cart:"
	// cart     = "cart:"
	// store    = "store:"
	// good     = "good:"
)

var _ RedisQuerier = (*RedisRepository)(nil)
