package redis_sever

import "context"

type RedisQuerier interface {
	SetGood(ctx context.Context, arg SetGoodInCartParams) error
	GetCartList(ctx context.Context, arg SetGoodInCartParams) error
}

const (
	//cart
	customer = "customer"
	cart     = "cart:"
	store    = "store:"
	good     = "good:"
)

var _ RedisQuerier = (*RedisRepository)(nil)
