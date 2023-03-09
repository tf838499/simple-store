package clerk

import (
	"context"
	"simple-store/internal/adapter/repository/PostgresDB"
)

type GoodRepository interface {
	GetGoodList(ctx context.Context) ([]PostgresDB.Good, error)
	GetGoodListByPage(ctx context.Context, GoodListByPageParams PostgresDB.GetGoodListByPageParams) ([]PostgresDB.Good, error)
}
