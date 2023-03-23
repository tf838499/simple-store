package redisclient

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
	SetKey := prefixCustomer + arg[0].CustomerID
	data := make(map[string]interface{})
	for i := range arg {
		field := arg[i].GoodName
		data[field] = arg[i].GoodAmount
		err := q.SetGoodPrice(ctx, GoodPriceInfo{Name: field, Price: arg[i].GoodPrice})
		if err != nil {
			return err
		}
	}
	err := q.Client.HMSet(ctx, SetKey, data).Err()
	if err != nil {
		return err
	}
	return nil
}
func (q *RedisRepository) DeleteGood(ctx context.Context, arg []GoodInCartParams) error {
	SetKey := prefixCustomer + arg[0].CustomerID
	// data := make(map[string]interface{})
	for i := range arg {
		field := arg[i].GoodName
		err := q.Client.HDel(ctx, SetKey, field).Err()
		if err != nil {
			return err
		}
		err = q.Client.Del(ctx, field).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

type GoodInRedisParams struct {
	CustomerID string
	GoodName   string
	GoodAmount int
	GoodPrice  string
}

func (q *RedisRepository) GetCartListCache(ctx context.Context, arg string) ([]GoodInRedisParams, error) {
	// SetKey :=
	SetKey := prefixCustomer + arg
	fields, err := q.Client.HGetAll(ctx, SetKey).Result()
	if err != nil {
		return nil, err
	}
	var GoodsCart []GoodInRedisParams
	for Name, amountString := range fields {
		vInt, err := strconv.Atoi(amountString)
		if err != nil {
			return nil, err
		}
		GoodsCart = append(GoodsCart, GoodInRedisParams{
			CustomerID: arg,
			GoodName:   Name,
			GoodAmount: vInt,
		})
	}
	return GoodsCart, nil
}

func (q *RedisRepository) GetGoodPrice(ctx context.Context, arg []string) ([]int, error) {
	// SetKey :=
	var GetKey []string
	for i := range arg {
		GetKey = append(GetKey, prefixPrice+arg[i])
	}
	var goodPrice []int
	Values, err := q.Client.MGet(ctx, GetKey...).Result()
	if err != nil {
		return goodPrice, err
	}
	for i := range Values {
		Price, err := strconv.Atoi(Values[i].(string))
		if err != nil {
			goodPrice = append(goodPrice, -1)
		} else {
			goodPrice = append(goodPrice, Price)
		}
	}
	return goodPrice, nil
}

type GoodPriceInfo struct {
	Name  string
	Price int
}

func (q *RedisRepository) SetGoodPrice(ctx context.Context, arg GoodPriceInfo) error {
	// SetKey
	Setkey := prefixPrice + arg.Name
	err := q.Client.Set(ctx, Setkey, arg.Price, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
func (q *RedisRepository) DeleteGoodPrice(ctx context.Context, arg GoodPriceInfo) error {
	// SetKey
	Setkey := prefixPrice + arg.Name
	err := q.Client.Del(ctx, Setkey).Err()
	if err != nil {
		return err
	}
	return nil
}
