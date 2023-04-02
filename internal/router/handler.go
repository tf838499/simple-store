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
	OauthToken := NewOAuthMiddleware(app)
	// We mount all handlers under /api path
	r := router.Group("/api")
	v := r.Group("/v1")
	v.GET("/callback", OauthToken.Callback)

	v.GET("/goods", v1.ListGoods(app))
	// Add barter namespace
	clerkGroup := v.Group("/clerk")
	{
		clerkGroup.POST("/goods", v1.AddNewGoods(app))
		clerkGroup.PUT("/goods", v1.UpdateGoods(app))
		clerkGroup.DELETE("/goods", v1.DeleteGoods(app))
	}
	customerGroup := v.Group("/customer").Use(OauthToken.AuthMiddleware())
	{
		customerGroup.POST("/order", v1.CreateOrder(app))
		customerGroup.GET("/cardlist/", v1.CartLists(app))
		customerGroup.POST("/cardlist/good/", v1.AddCartGoods(app))
		customerGroup.DELETE("/cardlist/good/", v1.DeleteCartGoods(app))
	}

}
