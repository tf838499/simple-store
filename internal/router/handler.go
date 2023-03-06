package router

import (
	"github.com/gin-gonic/gin"

	"simple-store/internal/app"
)

func RegisterHandlers(router *gin.Engine, app *app.Application) {
	registerAPIHandlers(router, app)
}

func registerAPIHandlers(router *gin.Engine, app *app.Application) {
	// Build middlewares
	// BearerToken := NewAuthMiddlewareBearer(app)

	// We mount all handlers under /api path
	r := router.Group("/api")
	v1 := r.Group("/v1")

	// Add health-check
	// v1.GET("/health", handlerHealthCheck())

	// Add auth namespace
	// authGroup := v1.Group("/auth")
	// {
	// 	authGroup.POST("/traders", RegisterTrader(app))
	// 	authGroup.POST("/traders/login", LoginTrader(app))
	// }

	// Add barter namespace
	clerkGroup := v1.Group("/clerk")
	{
		// barterGroup.POST("/goods", PostGood(app))
		clerkGroup.GET("/goods", ListMyGoods(app))
		// barterGroup.GET("/goods/traders", ListOthersGoods(app))
		// barterGroup.DELETE("/goods/:good_id", RemoveMyGood(app))
		// barterGroup.POST("/goods/exchange", ExchangeGoods(app))
	}
}
