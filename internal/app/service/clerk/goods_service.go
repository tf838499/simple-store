package clerk

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"simple-store/internal/adapter/repository/PostgresDB"
	"simple-store/internal/domain/common"
)

type GoodListParam struct {
	Page int32
}

func (c *ClerkService) ListGoods(ctx context.Context, param GoodListParam) ([]PostgresDB.Good, error) {
	var PageOffset int32 = 15
	PageLimit := 1 + PageOffset*(param.Page-1)
	goods, err := c.goodRepo.GetGoodListByPage(ctx, PostgresDB.GetGoodListByPageParams{Limit: PageLimit, Offset: PageOffset})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, common.NewError(common.ErrorCodeResourceNotFound, err)
		}
		c.logger(ctx).Error().Err(err).Msg("failed to get good")
		return nil, common.NewError(common.ErrorCodeInternalProcess, err)
	}

	return goods, err
}

type GoodInfoParam struct {
	ImageName string
	Descript  string
	Price     int64
	Class     string
}

func (c *ClerkService) AddGoods(ctx context.Context, param []PostgresDB.InsertGoodsParams) error {

	err := c.goodRepo.InsertGoodsWithTx(ctx, param)
	if err != nil {
		log.Println(err.Error())
		c.logger(ctx).Error().Err(err).Msg("failed to insert good")
		return err
	}

	return err
}
func (c *ClerkService) ChangeGoods(ctx context.Context, param PostgresDB.UpdateGoodParams) error {

	err := c.goodRepo.UpdateGood(ctx, param)
	if err != nil {
		c.logger(ctx).Error().Err(err).Msg("failed to update good")
		return err
	}

	return err
}

func (c *ClerkService) RemoveGood(ctx context.Context, goodIDParam int32) error {

	err := c.goodRepo.DeleteGood(ctx, goodIDParam)
	if err != nil {
		c.logger(ctx).Error().Err(err).Msg("failed to delete good")
		return err
	}

	return err
}
