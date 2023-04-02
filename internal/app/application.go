package app

import (
	"context"
	"fmt"
	"log"
	"sync"

	// "github.com/golang-migrate/migrate/v4/database/postgres"
	"simple-store/internal/adapter/repository/PostgresDB"
	"simple-store/internal/adapter/repository/RedisCache"
	"simple-store/internal/app/service/clerk"
	"simple-store/internal/app/service/customer"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	// "github.com/chatbotgang/go-clean-architecture-template/internal/adapter/repository/postgres"
	// "github.com/chatbotgang/go-clean-architecture-template/internal/adapter/server"
	// "github.com/chatbotgang/go-clean-architecture-template/internal/app/service/auth"
	// "github.com/chatbotgang/go-clean-architecture-template/internal/app/service/barter"
)

type Application struct {
	Params          ApplicationParams
	ClerkService    *clerk.ClerkService
	CustomerService *customer.CustomerService
	// BarterService *barter.BarterService
}

type ApplicationParams struct {
	Env         string
	DatabaseDSN string
	DBHost      string
	DBPort      string
	DBUser      string
	DBname      string
	DBPassword  string

	RedisHost     string
	RedisPort     string
	Redisname     int
	RedisPassword string
	RedisPoolSize int
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
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		params.DBHost,
		params.DBPort,
		params.DBUser,
		params.DBname,
		params.DBPassword,
	)
	db := sqlx.MustOpen("postgres", dsn)
	if err := db.Ping(); err != nil {
		return nil, err
	}

	pgRepo := PostgresDB.NewPostgresRepository(db)

	client := redis.NewClient(&redis.Options{
		Addr:     params.RedisHost + ":" + params.RedisPort,
		Password: params.RedisPassword, // no password set
		DB:       params.Redisname,     // use default DB
		PoolSize: params.RedisPoolSize, // 連接詞數量
	})
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	redisRepo := RedisCache.NewRedisRepository(client)
	app := &Application{
		Params: params,
		ClerkService: clerk.NewClerkService(ctx, clerk.ClerkServiceParam{
			GoodRepo: pgRepo,
		}),
		CustomerService: customer.NewCustomerService(ctx, customer.CustomerServiceParam{
			CartRepo:  redisRepo,
			OrderRepo: pgRepo,
			AuthRepo:  redisRepo,
		}),
	}

	return app, nil
}
