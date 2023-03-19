package v1

import (
	"fmt"
	"log"
	"simple-store/internal/app"
	"simple-store/internal/domain/common"
	"simple-store/internal/router/api/reponse"

	"github.com/gin-gonic/gin"
)

// @Summary Clerk Get Goods
// @Tags Clerk
// @version 1.0
// @produce application/json
// @param token Json string true "token"
// @param page Json string true "page"
// @Success 200 string string "success"
// @Failure 404 {page} errcode.error "no_found_item"
// @Failure 400 {page} errcode.error "invalid_parameter"
// @Router api/v1/clerk/goods [get]
func CreateOrder(app *app.Application) gin.HandlerFunc {

	type CarGoods struct {
		Good_id    string `json:"image_name"`
		Descript   string `json:"descript"`
		Totalprice int    `json:"price"`
		Amount     string `json:"class"`
		Owner      string `json:"owner"`
	}
	// "id" integer PRIMARY KEY,
	// "created_at" timestamptz DEFAULT 'now()',
	// "updated_at" timestamptz,
	// "delete_at" timestamptz,
	// "email" text,
	// "good_id" int[],
	// "good_price" int[],
	// "Allimage_name" text
	type Body struct {
		Email   string  `json:"email" form:"email" binding:"required"`
		Good_id []int32 `json:"good_id" form:"good_id" binding:"required"`
	}

	// type Response struct {
	// 	Goods []Good `json:"goods"`
	// }

	return func(c *gin.Context) {
		ctx := c.Request.Context()
		var body Body
		fmt.Println(ctx)
		err := c.ShouldBind(&body)
		if err != nil {
			log.Panicf(err.Error())
			reponse.RespondWithError(c,
				common.NewError(common.ErrorCodeParameterInvalid, err, common.WithMsg("invalid parameter")))
			return
		}

		// Invoke service
		// goods, err := app.ClerkService.ListGoods(ctx, clerk.GoodListParam{Limit: 15, Offset: body.Page})
		// if err != nil {
		// 	log.Panicf(err.Error())
		// 	msg := "no found item"
		// 	reponse.RespondWithError(c,
		// 		common.NewError(common.ErrorCodeResourceNotFound, errors.New(msg), common.WithMsg(msg)))
		// 	return
		// }
		// response result
		// resp := Response{Goods: []Good{}}
		// for i := range goods {
		// 	g := goods[i]
		// 	hashID := backstage.EncodeIDKey(int(g.ID))
		// 	resp.Goods = append(resp.Goods, Good{
		// 		ID:        hashID,
		// 		CreatedAt: g.CreatedAt.Time,
		// 		ImageName: g.ImageName.String,
		// 		Descript:  g.Descript.String,
		// 		Price:     int(g.Price.Int64),
		// 		Class:     g.Class.String,
		// 	})
		// }
		// reponse.RespondWithJSON(c, http.StatusOK, resp)
	}
}

// @Summary Clerk Get Goods
// @Tags Clerk
// @version 1.0
// @produce application/json
// @param token Json string true "token"
// @param page Json string true "page"
// @Success 200 string string "success"
// @Failure 404 {page} errcode.error "no_found_item"
// @Failure 400 {page} errcode.error "invalid_parameter"
// @Router api/v1/clerk/goods [get]
func AddCartGoods(app *app.Application) gin.HandlerFunc {

	type CarGoods struct {
		Good_id    string `json:"image_name"`
		Descript   string `json:"descript"`
		Totalprice int    `json:"price"`
		Amount     string `json:"class"`
		Owner      string `json:"owner"`
	}
	// "id" integer PRIMARY KEY,
	// "created_at" timestamptz DEFAULT 'now()',
	// "updated_at" timestamptz,
	// "delete_at" timestamptz,
	// "email" text,
	// "good_id" int[],
	// "good_price" int[],
	// "Allimage_name" text
	type Body struct {
		Email       string `json:"email" form:"email" binding:"required"`
		Good_ID     int32  `json:"good_id" form:"good_id" binding:"required"`
		Good_Price  int    `json:"good_price" form:"good_price" binding:"required"`
		Good_Amount int    `json:"good_amount" form:"good_amount" binding:"required"`
		Good_Name   int32  `json:"good_name" form:"good_name" binding:"required"`
	}

	// type Response struct {
	// 	Goods []Good `json:"goods"`
	// }

	return func(c *gin.Context) {
		ctx := c.Request.Context()
		var body Body
		fmt.Println(ctx)
		err := c.ShouldBind(&body)
		if err != nil {
			log.Panicf(err.Error())
			reponse.RespondWithError(c,
				common.NewError(common.ErrorCodeParameterInvalid, err, common.WithMsg("invalid parameter")))
			return
		}

		// Invoke service
		// goods, err := app.ClerkService.ListGoods(ctx, clerk.GoodListParam{Limit: 15, Offset: body.Page})
		// if err != nil {
		// 	log.Panicf(err.Error())
		// 	msg := "no found item"
		// 	reponse.RespondWithError(c,
		// 		common.NewError(common.ErrorCodeResourceNotFound, errors.New(msg), common.WithMsg(msg)))
		// 	return
		// }
		// response result
		// resp := Response{Goods: []Good{}}
		// for i := range goods {
		// 	g := goods[i]
		// 	hashID := backstage.EncodeIDKey(int(g.ID))
		// 	resp.Goods = append(resp.Goods, Good{
		// 		ID:        hashID,
		// 		CreatedAt: g.CreatedAt.Time,
		// 		ImageName: g.ImageName.String,
		// 		Descript:  g.Descript.String,
		// 		Price:     int(g.Price.Int64),
		// 		Class:     g.Class.String,
		// 	})
		// }
		// reponse.RespondWithJSON(c, http.StatusOK, resp)
	}
}
