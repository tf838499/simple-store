package RedisClient

import (
	"context"
	"strconv"
)

type GoodInCartParams struct {
	CustomerID string
	GoodName   string
	GoodPrice  int
	GoodAmount int
}

func (q *RedisRepository) SetGood(ctx context.Context, arg []GoodInCartParams) error {
	SetKey := customer + arg[0].CustomerID
	data := make(map[string]interface{})
	for i := range arg {
		field := arg[i].GoodName + "_" + strconv.Itoa(arg[i].GoodPrice)
		data[field] = arg[i].GoodAmount
	}
	err := q.Client.HMSet(ctx, SetKey, data).Err()
	if err != nil {
		return err
	}
	return nil
}
func (q *RedisRepository) DeleteGood(ctx context.Context, arg []GoodInCartParams) error {
	SetKey := customer + arg[0].CustomerID
	// data := make(map[string]interface{})
	for i := range arg {
		field := arg[i].GoodName + "_" + strconv.Itoa(arg[i].GoodPrice)
		err := q.Client.HDel(ctx, SetKey, field).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

type GoodInRedisParams struct {
	CustomerID       string
	GoodNameAndPrice string
	GoodAmount       int
}

func (q *RedisRepository) GetCartList(ctx context.Context, arg GoodInCartParams) ([]GoodInRedisParams, error) {
	// SetKey :=
	SetKey := customer + arg.CustomerID
	fields, err := q.Client.HGetAll(context.Background(), SetKey).Result()
	if err != nil {
		return nil, err
	}
	var GoodsCart []GoodInRedisParams
	for k, vString := range fields {
		vInt, err := strconv.Atoi(vString)
		if err != nil {
			return nil, err
		}
		GoodsCart = append(GoodsCart, GoodInRedisParams{
			CustomerID:       SetKey,
			GoodNameAndPrice: k,
			GoodAmount:       vInt,
		})
	}
	return GoodsCart, nil
}
