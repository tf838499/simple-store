package clerk

import (
	"context"
	"simple-store/internal/adapter/repository/PostgresDB"
)

type GoodRepository interface {
	// GetGoodList(ctx context.Context) ([]PostgresDB.Good, error)
	GetGoodListByPage(ctx context.Context, GoodListByPageParams PostgresDB.GetGoodListByPageParams) ([]PostgresDB.Good, error)
	InsertGoods(ctx context.Context, GoodsParams PostgresDB.InsertGoodsParams) error
	InsertGoodsWithTx(ctx context.Context, GoodsParams []PostgresDB.InsertGoodsParams) error
	UpdateGood(ctx context.Context, GoodParams PostgresDB.UpdateGoodParams) error
	DeleteGood(ctx context.Context, id int32) error
}
