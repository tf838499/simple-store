package v1

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"simple-store/internal/app"
	"simple-store/internal/app/service/clerk"
	"simple-store/internal/domain/common"

	// "github.com/bytedance/sonic"

	// "github.com/bytedance/sonic/decoder"
	"github.com/gin-gonic/gin"
	// "github.com/chatbotgang/go-clean-architecture-template/internal/app"
	"simple-store/internal/router/api/reponse"
)

// @Summary Clerk Get Goods
// @Tags Clerk
// @version 1.0
// @produce application/json
// @param token formData string true "token"
// @param page formData string true "page"
// @Success 200 string string "success"
// @Failure 400 {page} errcode.error "invalid_parameter"
// @Failure 500 {page} errcode.error "invalid_parameter"
// @Router api/v1/clerk/goods [get]
func ListGoods(app *app.Application) gin.HandlerFunc {

	type Good struct {
		ID        int       `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		ImageName string    `json:"image_name"`
		Descript  string    `json:"descript"`
		Price     int       `json:"price"`
		Class     string    `json:"class"`
	}

	type Body struct {
		Page int32 `json:"page" binding:"page"`
	}

	type Response struct {
		Goods []Good `json:"goods"`
	}

	return func(c *gin.Context) {
		ctx := c.Request.Context()
		var body Body
		err := c.ShouldBind(&body)
		if err != nil {
			reponse.RespondWithError(c,
				common.NewError(common.ErrorCodeParameterInvalid, err, common.WithMsg("invalid parameter")))
			return
		}

		// Invoke service
		goods, err := app.ClerkService.ListGoods(ctx, clerk.GoodParam{Page: body.Page})
		if err != nil {
			msg := "no found item"
			reponse.RespondWithError(c,
				common.NewError(common.ErrorCodeResourceNotFound, errors.New(msg), common.WithMsg(msg)))
			return
		}
		// response result
		resp := Response{Goods: []Good{}}
		for i := range goods {
			g := goods[i]
			resp.Goods = append(resp.Goods, Good{
				ID:        int(g.ID),
				CreatedAt: g.CreatedAt.Time,
				ImageName: g.ImageName.String,
				Descript:  g.Descript.String,
				Price:     int(g.Price.Int64),
				Class:     g.Class.String,
			})
		}
		reponse.RespondWithJSON(c, http.StatusOK, resp)
	}
}

// @Summary Add Goods
// @Tags Clerk
// @version 1.0
// @produce application/json
// @param ImageName formData string true "image name"
// @param Descript formData string true "descript"
// @param Price formData string true "price"
// @param Class formData string true "class"
// @Success 200 string string ""
// @Failure 400 參數類型 數據類型 註釋
// @Router api/v1/clerk/goods [post]
func AddNewGoods(app *app.Application) gin.HandlerFunc {

	type Good struct {
		ID        int       `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		ImageName string    `json:"image_name"`
		Descript  string    `json:"descript"`
		Price     int       `json:"price"`
		Class     string    `json:"class"`
	}

	type Body struct {
		Page int32 `json:"page" binding:"required"`
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
		goods, err := app.ClerkService.ListGoods(ctx, clerk.GoodParam{Page: body.Page})
		// goods, err := app.BarterService.ListMyGoods(ctx, barter.ListMyGoodsParam{

		resp := Response{
			Goods: []Good{},
		}
		for i := range goods {
			g := goods[i]
			resp.Goods = append(resp.Goods, Good{ImageName: g.Class.String})
		}

		// respondWithJSON(c, http.StatusOK, resp)
	}
}

// @Summary Update Goods
// @Tags Clerk
// @version 1.0
// @produce application/json
// @param ImageName formData string true "image name"
// @param Descript formData string true "descript"
// @param Price formData string true "price"
// @param Class formData string true "class"
// @Success 200 string string 成功後返回的值
// @Failure 狀態碼 參數類型 數據類型 註釋
// @Router api/v1/clerk/goods/{id} [put]
func UpdateGoods(app *app.Application) gin.HandlerFunc {

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
		Page int32 `json:"page" binding:"required"`
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
		goods, err := app.ClerkService.ListGoods(ctx, clerk.GoodParam{Page: body.Page})
		// goods, err := app.BarterService.ListMyGoods(ctx, barter.ListMyGoodsParam{

		resp := Response{
			Goods: []Good{},
		}
		for i := range goods {
			g := goods[i]
			resp.Goods = append(resp.Goods, Good{ImageName: g.Class.String})
		}

		// respondWithJSON(c, http.StatusOK, resp)
	}
}

// @Summary Update Goods
// @Tags Clerk
// @version 1.0
// @produce application/json
// @param ImageName formData string true "image name"
// @param Descript formData string true "descript"
// @param Price formData string true "price"
// @param Class formData string true "class"
// @Success 200 string string 成功後返回的值
// @Failure 狀態碼 參數類型 數據類型 註釋
// @Router api/v1/clerk/goods/{id} [delete]
func DeleteGoods(app *app.Application) gin.HandlerFunc {

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
		Page int32 `json:"page" binding:"required"`
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
		goods, err := app.ClerkService.ListGoods(ctx, clerk.GoodParam{Page: body.Page})
		// goods, err := app.BarterService.ListMyGoods(ctx, barter.ListMyGoodsParam{

		resp := Response{
			Goods: []Good{},
		}
		for i := range goods {
			g := goods[i]
			resp.Goods = append(resp.Goods, Good{ImageName: g.Class.String})
		}

		// respondWithJSON(c, http.StatusOK, resp)
	}
}
