# Simple store

**Simple store**是基於*Crescendo Lab* go-clean-arch 模板實現基本的網頁商店後端功能

## Table of Content:

- [About The project](#about-the-project)
- [Technologies](#technologies)
- [Architecture](#screenshots)
- [Setup](#setup)
- [Approach](#approach)
- [Status](#status)
- [License](#license)

## About The project

這個 Side project 是用來練習跟實現的小專案，後端架構參考 Crescendo Lab 的 DDD(Domain-Driven Design)實現，裡面使用到 Postgres 資料庫來存儲商品目錄以及顧客訂單，和使用 Redis 實現簡易的購物車，以及 OAuth 跟 JWT 確認身份以及驗證，來管理商品以及更新訂單

## Technologies

<details>
<summary>後端</summary>

- [Golang](https://go.dev): ^1.17
- [gin](https://github.com/gin-gonic/gin): ~1.7.7
- [viper](https://github.com/spf13/viper): ~1.15.0
- [cobra](https://github.com/spf13/cobra): ~1.6.1
- [oauth](https://github.com/golang/oauth2): ~7.0

<details>

<details>
<summary>日誌文件</summary>

- [swagger](https://github.com/swaggo/swag): ~1.8.1
- [zerolog](https://github.com/rs/zerolog): ~1.26.1
<details>

<details><summary>資料庫</summary>

- [postgres](https://github.com/swaggo/swag):
- [redis](https://github.com/rs/zerolog):
- [sqlc](https://github.com/rs/zerolog):
- [migrations](https://github.com/rs/zerolog):
<details>

<details><summary>測試</summary>

- [testify](https://github.com/stretchr/testify): ^1.8.0
- [mockgen](https://github.com/golang/mock): ~1.6.0
- [testfixtures](https://github.com/go-testfixtures/testfixtures): ^3.8.0
- [migrate](https://github.com/golang-migrate/migrate): ^4.15.0
- [dockertest](https://github.com/ory/dockertest): ^3.9.0
<details>ㄋ

## Architecture

## Setup

- 啟動資料庫 : 啟動 postgres 跟 redis
  ```sh
  npm install npm@latest -g
  ```
  docker-compose up
- 建立資料庫表格
  make migrateup
- 運行後端
  go run main.go

## Approach

- 用 migrations 來創建跟版控資料庫，以及不使用 GORM，改用 sqlc 來自動產生 sql 程式碼跟提升對 db 操作效能。
- 在 router 內寫 swagger 的註解讓自動產生 api 的相關資訊，以及在使用 zerolog 日誌庫，搜集 API 相關訊息
- 使用 gin 框架來開發後端，以及使用 viper 跟 cobra 自動產生 CMD 與抓取 config，在 middleware 裡添加 google 的 ouath 認證

## Status
