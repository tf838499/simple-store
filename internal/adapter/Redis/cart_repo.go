package redis_sever

import "context"

type SetGoodInCartParams struct {
	CustomerID string
	GoodName   string
	GoodPrice  int
	GoodAmount int
}

func (q *RedisRepository) SetGood(ctx context.Context, arg SetGoodInCartParams) error {
	q.Client.HSet(ctx, "Hash_key", "user1", "ken", "age", 25, "height", 190).Result()
	return nil
}
func (q *RedisRepository) GetCartList(ctx context.Context, arg SetGoodInCartParams) error {

	return nil
}
