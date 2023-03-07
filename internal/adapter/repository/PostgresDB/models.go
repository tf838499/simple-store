// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package PostgresDB

import (
	"database/sql"
)

type CarList struct {
	ID        int32          `json:"id"`
	CreatedAt sql.NullTime   `json:"created_at"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
	DeleteAt  sql.NullTime   `json:"delete_at"`
	Account   sql.NullString `json:"account"`
	Phone     sql.NullString `json:"phone"`
	Email     sql.NullString `json:"email"`
	Total     sql.NullInt64  `json:"total"`
	Status    sql.NullInt32  `json:"status"`
	OrderID   sql.NullInt32  `json:"order_id"`
}

type Good struct {
	ID        int32          `json:"id"`
	CreatedAt sql.NullTime   `json:"created_at"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
	DeleteAt  sql.NullTime   `json:"delete_at"`
	ImageName sql.NullString `json:"image_name"`
	Descript  sql.NullString `json:"descript"`
	Price     sql.NullInt64  `json:"price"`
	Class     sql.NullString `json:"class"`
}

type Order struct {
	ID        int32         `json:"id"`
	CreatedAt sql.NullTime  `json:"created_at"`
	UpdatedAt sql.NullTime  `json:"updated_at"`
	DeleteAt  sql.NullTime  `json:"delete_at"`
	GoodID    sql.NullInt32 `json:"good_id"`
	Amount    sql.NullInt32 `json:"amount"`
}