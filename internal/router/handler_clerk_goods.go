package router

import (
	"fmt"
	"time"

	"simple-store/internal/app"
	"simple-store/internal/app/service/clerk"

	// "github.com/bytedance/sonic"

	"github.com/bytedance/sonic/decoder"
	"github.com/gin-gonic/gin"
	// "github.com/chatbotgang/go-clean-architecture-template/internal/app"
)

func ListMyGoods(app *app.Application) gin.HandlerFunc {

	type Good struct {
		ID        int       `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		DeleteAt  time.Time `json:"delete_at"`
		ImageName string    `json:"image_name"`
		Descript  string    `json:"descript"`
		Price     int       `json:"price"`
		Class     string    `json:"class"`
	}

	type Body struct {
		Name string `json:"name" binding:"required"`
	}

	type Response struct {
		Goods []Good `json:"goods"`
	}

	return func(c *gin.Context) {
		ctx := c.Request.Context()
		fmt.Println(ctx)
		var body Body
		err := c.ShouldBind(&body)

		// err = sonic.Unmarshal(p, &body)
		if err != nil {
			return
		}

		// Invoke service
		goods, err := app.ClerkService.ListMyGoods(ctx, clerk.PostGoodParam{GoodName: body.Name})
		// goods, err := app.BarterService.ListMyGoods(ctx, barter.ListMyGoodsParam{

		resp := Response{
			Goods: []Good{},
		}
		for i := range goods {
			g := goods[i]
			resp.Goods = append(resp.Goods, Good{ImageName: g.Name})
		}

		// respondWithJSON(c, http.StatusOK, resp)
	}
}
