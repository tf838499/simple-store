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
	SetKey := prefixCustomer + arg[0].CustomerID
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
	SetKey := prefixCustomer + arg[0].CustomerID
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

func (q *RedisRepository) GetCartList(ctx context.Context, arg string) ([]GoodInRedisParams, error) {
	// SetKey :=
	SetKey := prefixCustomer + arg
	fields, err := q.Client.HGetAll(ctx, SetKey).Result()
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
			CustomerID:       arg,
			GoodNameAndPrice: k,
			GoodAmount:       vInt,
		})
	}
	return GoodsCart, nil
}

type GoodPrice struct {
	Price map[string]int
}

func (q *RedisRepository) GetGoodPrice(ctx context.Context, arg []string) (GoodPrice, error) {
	// SetKey :=
	var GetKey []string
	for i := range arg {
		GetKey = append(GetKey, prefixPrice+arg[i])
	}
	var goodMap GoodPrice
	Values, err := q.Client.MGet(ctx, GetKey...).Result()
	if err != nil {
		return goodMap, err
	}
	for i := range Values {
		Price, err := strconv.Atoi(Values[i].(string))
		if err != nil {
			goodMap.Price[arg[i]] = -1
		} else {
			goodMap.Price[arg[i]] = Price
		}
	}
	return goodMap, nil
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
