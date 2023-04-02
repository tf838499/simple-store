package customer

import (
	"context"
	"database/sql"
	"log"
	"simple-store/internal/adapter/repository/PostgresDB"
	"simple-store/internal/adapter/repository/RedisCache"
	"simple-store/internal/domain/cartlist"
)

type CartParams struct {
	Email string
}

func (c *CustomerService) GetCartList(ctx context.Context, param CartParams) (cartlist.GoodInCart, error) {
	var GoodInCart = cartlist.GoodInCart{}
	Goods, err := c.cartRepo.GetCartListCache(ctx, param.Email)
	if err != nil {
		log.Println(err.Error())
		c.logger(ctx).Error().Err(err).Msg("failed to get good in cartlist")
		return GoodInCart, err
	}

	for i := 0; i < len(Goods); i++ {
		GoodInCart.Name = append(GoodInCart.Name, Goods[i].Name)
		GoodInCart.Amount = append(GoodInCart.Amount, Goods[i].Amount)
	}
	price, err := c.cartRepo.GetGoodPrice(ctx, GoodInCart.Name)
	GoodInCart.Price = price
	if err != nil {
		log.Println(err.Error())
		c.logger(ctx).Error().Err(err).Msg("failed to get good")
		return GoodInCart, err
	}

	return GoodInCart, err
}
func (c *CustomerService) SetGoodInCart(ctx context.Context, param []RedisCache.GoodInCartParams) error {

	err := c.cartRepo.SetGood(ctx, param)
	if err != nil {
		log.Println(err.Error())
		c.logger(ctx).Error().Err(err).Msg("failed to insert good")
		return err
	}
	err = c.cartRepo.MSetGoodPrice(ctx, param)
	if err != nil {
		log.Println(err.Error())
		c.logger(ctx).Error().Err(err).Msg("failed to cache good price")
	}

	return err
}
func (c *CustomerService) DeleteGoodInCart(ctx context.Context, param []RedisCache.GoodInCartParams) error {

	err := c.cartRepo.DeleteGood(ctx, param)
	if err != nil {
		log.Println(err.Error())
		c.logger(ctx).Error().Err(err).Msg("failed to delete good")
		return err
	}

	return err
}

type OrderParams struct {
	Email      string
	GoodAmount []int32
	GoodName   []string
	Message    string
}
type OrderInfo struct {
	Email      string
	GoodAmount []int32
	GoodName   []string
	Message    string
	TotalPrice int
	Status     string
}

func (c *CustomerService) InsertGoodInCart(ctx context.Context, param OrderParams) (OrderInfo, error) {
	var orderInfo OrderInfo
	var orderDbParam = PostgresDB.InsertOrderParams{
		Owner:   sql.NullString{String: param.Email, Valid: true},
		GoodID:  param.GoodName,
		Amount:  param.GoodAmount,
		Message: sql.NullString{String: param.Message, Valid: true},
		Status:  sql.NullInt32{Int32: 3, Valid: true},
	}
	priceList, err := c.cartRepo.GetGoodPrice(ctx, param.GoodName)
	if err != nil {
		log.Println(err.Error())
		c.logger(ctx).Error().Err(err).Msg("failed to get good price")
		return orderInfo, err
	}
	for i := range param.GoodName {
		if i < len(priceList) && priceList[i] == -1 {
			GoodInfo, err := c.orderRepo.GetGoodByName(ctx, sql.NullString{String: param.GoodName[i], Valid: true})
			if err != nil {
				log.Println(err.Error())
				c.logger(ctx).Error().Err(err).Msg("failed to get good")
				return orderInfo, err
			}
			priceList[i] = int(GoodInfo.Price.Int64)

			err = c.cartRepo.SetGoodPrice(ctx, RedisCache.GoodPriceInfo{Name: param.GoodName[i], Price: priceList[i]})
			if err != nil {
				log.Println(err.Error())
				c.logger(ctx).Error().Err(err).Msg("failed to cache good price")
			}
		}
	}

	totalPrice := cartlist.CheckoutPrice(priceList, param.GoodAmount, param.GoodName)
	orderDbParam.TotalPrice = sql.NullInt32{Int32: int32(totalPrice), Valid: true}
	err = c.orderRepo.InsertOrder(ctx, orderDbParam)
	if err != nil {
		log.Println(err.Error())
		c.logger(ctx).Error().Err(err).Msg("failed to delete good")
		return orderInfo, err
	}

	orderInfo = OrderInfo{
		Email:      param.Email,
		GoodAmount: param.GoodAmount,
		GoodName:   param.GoodName,
		Message:    param.Message,
		TotalPrice: totalPrice,
		Status:     "下單成功",
	}

	return orderInfo, err
}
