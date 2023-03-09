package clerk

import (
	"context"
	"database/sql"
	"errors"
	"simple-store/internal/adapter/repository/PostgresDB"
	"simple-store/internal/domain/common"
)

type GoodParam struct {
	Page int32
}

func (c *ClerkService) ListGoods(ctx context.Context, param GoodParam) ([]PostgresDB.Good, error) {
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

	// var alterGoods []backstage.Good
	// for i := range goods {
	// 	g := goods[i]
	// 	alterGoods = append(alterGoods, backstage.Good{
	// 		ID:        int(g.ID),
	// 		CreatedAt: g.CreatedAt.Time,
	// 		ImageName: g.ImageName.String,
	// 		Descript:  g.Descript.String,
	// 		Price:     int(g.Price.Int64),
	// 		Class:     g.Class.String,
	// 	})
	// }

	return goods, err
}
