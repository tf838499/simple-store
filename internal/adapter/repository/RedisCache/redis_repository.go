package RedisCache

import (
	"context"
	"errors"
	"strconv"
	"time"
)

type GoodInCartParams struct {
	CustomerID string
	Name       string
	Price      int
	Amount     int
}

func (q *RedisRepository) SetGood(ctx context.Context, arg []GoodInCartParams) error {

	if arg == nil {
		return errors.New("param is invalid")
	}
	SetKey := prefixCustomer + arg[0].CustomerID
	data := make(map[string]interface{})
	for i := range arg {
		field := arg[i].Name
		data[field] = arg[i].Amount
	}
	err := q.Client.HMSet(ctx, SetKey, data).Err()
	if err != nil {
		return err
	}
	return nil
}
func (q *RedisRepository) DeleteGood(ctx context.Context, arg []GoodInCartParams) error {
	if arg == nil {
		return errors.New("param is invalid")
	}
	SetKey := prefixCustomer + arg[0].CustomerID
	// data := make(map[string]interface{})
	for i := range arg {
		field := arg[i].Name
		err := q.Client.HDel(ctx, SetKey, field).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func (q *RedisRepository) GetCartListCache(ctx context.Context, arg string) ([]GoodInCartParams, error) {
	SetKey := prefixCustomer + arg
	fields, err := q.Client.HGetAll(ctx, SetKey).Result()
	if err != nil {
		return nil, err
	}
	var GoodsCart []GoodInCartParams
	for Name, amountString := range fields {
		vInt, err := strconv.Atoi(amountString)
		if err != nil {
			return nil, err
		}
		GoodsCart = append(GoodsCart, GoodInCartParams{
			CustomerID: arg,
			Name:       Name,
			Amount:     vInt,
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

func (q *RedisRepository) MSetGoodPrice(ctx context.Context, arg []GoodInCartParams) error {
	// SetKey
	data := make(map[string]interface{})
	for i := range arg {
		field := prefixPrice + arg[i].Name
		data[field] = arg[i].Price
	}
	err := q.Client.MSet(ctx, data).Err()
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

type OauthInfo struct {
	Name string
	Code string
}

func (q *RedisRepository) SetUserInfo(ctx context.Context, arg OauthInfo) error {
	// SetKey
	Setkey := prefixUser + arg.Name
	err := q.Client.Set(ctx, Setkey, arg.Code, time.Hour*1).Err()
	if err != nil {
		return err
	}
	return nil
}
func (q *RedisRepository) GetUserInfo(ctx context.Context, arg OauthInfo) (string, error) {
	// SetKey
	Setkey := prefixUser + arg.Name
	value, err := q.Client.Get(ctx, Setkey).Result()
	if err != nil {
		return value, err
	}
	return value, nil
}
