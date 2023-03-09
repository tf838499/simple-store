package router

import (
	"github.com/gin-gonic/gin"

	"simple-store/internal/app"
	v1 "simple-store/internal/router/api/v1"
)

func RegisterHandlers(router *gin.Engine, app *app.Application) {
	registerAPIHandlers(router, app)
}

func registerAPIHandlers(router *gin.Engine, app *app.Application) {
	// Build middlewares
	// BearerToken := NewAuthMiddlewareBearer(app)

	// We mount all handlers under /api path
	r := router.Group("/api")
	v := r.Group("/v1")

	// Add health-check
	// v1.GET("/health", handlerHealthCheck())

	// Add auth namespace
	// authGroup := v1.Group("/auth")
	// {
	// 	authGroup.POST("/traders", RegisterTrader(app))
	// 	authGroup.POST("/traders/login", LoginTrader(app))
	// }

	// Add barter namespace
	clerkGroup := v.Group("/clerk")
	{
		// barterGroup.POST("/goods", PostGood(app))
		clerkGroup.GET("/goods", v1.ListGoods(app))
		clerkGroup.POST("/goods", v1.AddNewGoods(app))
		clerkGroup.PUT("/goods", v1.UpdateGoods(app))
		clerkGroup.DELETE("/goods", v1.DeleteGoods(app))

		// clerkGroup.GET("/order_list", v1.ListMyGoods(app))

		// barterGroup.DELETE("/goods/:good_id", RemoveMyGood(app))
		// barterGroup.POST("/goods/exchange", ExchangeGoods(app))
	}
}
