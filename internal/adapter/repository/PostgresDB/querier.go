// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package PostgresDB

import (
	"context"
	"database/sql"
)

type Querier interface {
	DeleteGood(ctx context.Context, id int32) error
	GetGetOrderByOwner(ctx context.Context, owner sql.NullString) ([]Order, error)
	GetGoodByName(ctx context.Context, imageName sql.NullString) (Good, error)
	GetGoodList(ctx context.Context) ([]Good, error)
	GetGoodListByPage(ctx context.Context, arg GetGoodListByPageParams) ([]Good, error)
	InsertGoods(ctx context.Context, arg InsertGoodsParams) error
	InsertOrder(ctx context.Context, arg InsertOrderParams) error
	UpdateGood(ctx context.Context, arg UpdateGoodParams) error
}

var _ Querier = (*Queries)(nil)
