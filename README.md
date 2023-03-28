# Simple store

**Simple store**是基於*Crescendo Lab* go-clean-arch模板實現基本的網頁商店後端功能

## 概述
此架構參考Crescendo Lab的DDD(Domain-Driven Design)實現，裡面使用到Postgres資料庫來存儲商品目錄以及顧客訂單，和使用redis實現簡易的購物車，以及OAuth跟JWT確認驗證是否為店員，來管理商品以及更新顧客訂單

## API


## Architecture



## Crescendo Barter

Crescendo Barter is a second-hand goods exchange application in which people can post their old goods and exchange them with others.

### User Stories

Account management:
- As a client, I want to register a trader account.
- As a client, I want to log in to the application through the registered trader account.

Second-hand Goods:
- As a trader, I want to post my old goods to the application so that others can see what I have.
- As a trader, I want to see all my posted goods.
- As a trader, I want to see others’ posted goods.
- As a trader, I want to remove some of my goods from the application.

Goods Exchange:
- As a trader, I want to exchange my own goods with others.
    
### Project Dependencies

<details><summary>Main application</summary>
  
- [Golang](https://go.dev): ^1.17
- [gin](https://github.com/gin-gonic/gin): ~1.7.7
- [viper](https://github.com/swaggo/swag): ~1.15.0
- [cobra](https://github.com/swaggo/swag): ~1.6.1
- [swagger](https://github.com/swaggo/swag): ~1.8.1
- [zerolog](https://github.com/rs/zerolog): ~1.26.1
- [sqlc](https://github.com/jmoiron/sqlx): ~1.15.0
- [PostgreSQL](https://www.postgresql.org/docs/13/index.html): ^13
- [Redis] (https://redis.io/): ~7.0
- [OAuth] (https://redis.io/): ~7.0
</details>

<details><summary>Test usage</summary>
  
- [testify](https://github.com/stretchr/testify): ^1.8.0
- [mockgen](https://github.com/golang/mock): ~1.6.0
- [testfixtures](https://github.com/go-testfixtures/testfixtures): ^3.8.0
- [migrate](https://github.com/golang-migrate/migrate): ^4.15.0
- [dockertest](https://github.com/ory/dockertest): ^3.9.0

</details>

### Development Guideline

See [development guideline](./docs/development-guideline.md).
