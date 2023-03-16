package clerk

import (
	"context"
	"simple-store/internal/adapter/repository/PostgresDB"
)

//go:generate mockgen -destination automock/good_repository.go -package=automock . GoodRepository
type GoodRepository interface {
	GetGoodListByPage(ctx context.Context, GoodListByPageParams PostgresDB.GetGoodListByPageParams) ([]PostgresDB.Good, error)
	InsertGoods(ctx context.Context, GoodsParams PostgresDB.InsertGoodsParams) error
	InsertGoodsWithTx(ctx context.Context, GoodsParams []PostgresDB.InsertGoodsParams) error
	UpdateGood(ctx context.Context, GoodParams PostgresDB.UpdateGoodParams) error
	DeleteGood(ctx context.Context, id int32) error
}
