package app

import (
	"context"
	"fmt"
	"log"
	"sync"

	// "github.com/golang-migrate/migrate/v4/database/postgres"
	"simple-store/internal/adapter/repository/PostgresDB"
	"simple-store/internal/app/service/clerk"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	// "github.com/chatbotgang/go-clean-architecture-template/internal/adapter/repository/postgres"
	// "github.com/chatbotgang/go-clean-architecture-template/internal/adapter/server"
	// "github.com/chatbotgang/go-clean-architecture-template/internal/app/service/auth"
	// "github.com/chatbotgang/go-clean-architecture-template/internal/app/service/barter"
)

type Application struct {
	Params       ApplicationParams
	ClerkService *clerk.ClerkService
	// BarterService *barter.BarterService
}

type ApplicationParams struct {
	Env         string
	DatabaseDSN string
}

func MustNewApplication(ctx context.Context, wg *sync.WaitGroup, params ApplicationParams) *Application {
	app, err := NewApplication(ctx, wg, params)
	if err != nil {
		log.Panicf("fail to new application, err: %s", err.Error())
	}
	return app
}

func NewApplication(ctx context.Context, wg *sync.WaitGroup, params ApplicationParams) (*Application, error) {
	// Create repositories
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		"0.0.0.0",
		5432,
		"postgres",
		"postgres",
		"postgres",
	)
	db := sqlx.MustOpen("postgres", dsn)
	if err := db.Ping(); err != nil {
		return nil, err
	}

	pgRepo := PostgresDB.NewPostgresRepository(db)
	fmt.Println(pgRepo)

	app := &Application{
		Params: params,
		ClerkService: clerk.NewClerkService(ctx, clerk.ClerkServiceParam{
			GoodRepo: pgRepo,
		}),
	}

	return app, nil
}
