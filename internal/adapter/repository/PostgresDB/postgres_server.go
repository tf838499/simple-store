package PostgresDB

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PostgresRepository struct {
	*Queries
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {

	return &PostgresRepository{
		Queries: New(db),
		db:      db,
	}
}

func (store *PostgresRepository) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		fmt.Println("Error occur, Start Rollback")
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
