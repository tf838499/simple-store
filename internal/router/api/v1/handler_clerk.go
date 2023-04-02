package v1

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"time"

	"simple-store/internal/adapter/repository/PostgresDB"
	"simple-store/internal/app"
	"simple-store/internal/app/service/clerk"
	"simple-store/internal/domain/backstage"
	"simple-store/internal/domain/common"
	"simple-store/internal/router/api/reponse"

	"github.com/gin-gonic/gin"
)

// @Summary Get Goods
// @Tags Clerk
// @version 1.0
// @produce application/json
// @param token formData string true "token"
// @param page formData string true "page"
// @Success 200 string string "success get"
// @Failure 404 string errcode.error "no_found_item"
// @Failure 400 string errcode.error "invalid_parameter"
// @Router /api/v1/clerk/goods/ [get]
func ListGoods(app *app.Application) gin.HandlerFunc {

	type Good struct {
		ID        string    `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		ImageName string    `json:"image_name"`
		Descript  string    `json:"descript"`
		Price     int       `json:"price"`
		Class     string    `json:"class"`
	}

	type Body struct {
		Page int32 `json:"page" form:"page" binding:"required"`
	}

	type Response struct {
		Goods []Good `json:"goods"`
	}

	return func(c *gin.Context) {
		ctx := c.Request.Context()
		var body Body
		err := c.ShouldBind(&body)
		if err != nil {
			log.Panicf(err.Error())
			reponse.RespondWithError(c,
				common.NewError(common.ErrorCodeParameterInvalid, err, common.WithMsg("invalid parameter")))
			return
		}

		// Invoke service
		goods, err := app.ClerkService.ListGoods(ctx, clerk.GoodListParam{Limit: 15, Offset: body.Page})
		if err != nil {
			log.Panicf(err.Error())
			msg := "no found item"
			reponse.RespondWithError(c,
				common.NewError(common.ErrorCodeResourceNotFound, errors.New(msg), common.WithMsg(msg)))
			return
		}
		// response result
		resp := Response{Goods: []Good{}}
		for i := range goods {
			g := goods[i]
			hashID := backstage.EncodeIDKey(int(g.ID))
			resp.Goods = append(resp.Goods, Good{
				ID:        hashID,
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
// @Failure 400 string errcode.error "invalid_insert_parameter"
// @Router /api/v1/clerk/goods [post]
func AddNewGoods(app *app.Application) gin.HandlerFunc {

	type Body struct {
		ImageName string `json:"image_name" form:"page" binding:"required"`
		Descript  string `json:"descript" form:"descript" binding:"required"`
		Price     int    `json:"price" form:"price" binding:"required"`
		Class     string `json:"class" form:"class" binding:"required"`
	}

	return func(c *gin.Context) {
		ctx := c.Request.Context()
		var body []Body
		err := c.ShouldBind(&body)

		if err != nil {
			log.Panicf(err.Error())
			reponse.RespondWithError(c,
				common.NewError(common.ErrorCodeParameterInvalid, err, common.WithMsg("invalid parameter")))
			return
		}
		GoodsInfo := []PostgresDB.InsertGoodsParams{}
		for i := range body {
			GoodsInfo = append(GoodsInfo,
				PostgresDB.InsertGoodsParams{
					ImageName: sql.NullString{String: body[i].ImageName, Valid: true},
					Descript:  sql.NullString{String: body[i].Descript, Valid: true},
					Price:     sql.NullInt64{Int64: int64(body[i].Price), Valid: true},
					Class:     sql.NullString{String: body[i].Class, Valid: true},
				})
		}
		// Invoke service

		err = app.ClerkService.AddGoods(ctx, GoodsInfo)
		if err != nil {
			reponse.RespondWithError(c,
				common.NewError(common.ErrorCodeResourceOperateFail, err, common.WithMsg("invalid insert parameter")))
			return
		}

		reponse.RespondWithoutBody(c, http.StatusOK)
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
// @Success 200 string string "success"
// @Failure 500 string string "invalid update parameter"
// @Router /api/v1/clerk/goods/{id} [put]
func UpdateGoods(app *app.Application) gin.HandlerFunc {

	type Body struct {
		ID        string `json:"id"`
		ImageName string `json:"image_name"`
		Descript  string `json:"descript"`
		Price     int    `json:"price"`
		Class     string `json:"class"`
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

		hashID := backstage.DncodeIDKey(body.ID)
		// Invoke service
		err = app.ClerkService.ChangeGoods(ctx,
			PostgresDB.UpdateGoodParams{
				ImageName: sql.NullString{String: body.ImageName, Valid: true},
				Descript:  sql.NullString{String: body.Descript, Valid: true},
				Price:     sql.NullInt64{Int64: int64(body.Price), Valid: true},
				Class:     sql.NullString{String: body.Class, Valid: true},
				ID:        int32(hashID),
			})

		if err != nil {
			reponse.RespondWithError(c,
				common.NewError(common.ErrorCodeResourceOperateFail, err, common.WithMsg("invalid update parameter")))
			return
		}

		reponse.RespondWithoutBody(c, http.StatusOK)
	}
}

// @Summary Delete Goods
// @Tags Clerk
// @version 1.0
// @produce application/json
// @param id formData string true "id"
// @Success 200 string string success
// @Failure 500 string string "invalid delete parameter"
// @Router /api/v1/clerk/goods/{id} [delete]
func DeleteGoods(app *app.Application) gin.HandlerFunc {

	type Body struct {
		ID string `json:"id"`
	}

	return func(c *gin.Context) {
		ctx := c.Request.Context()
		var body Body
		err := c.ShouldBind(&body)

		// err = sonic.Unmarshal(p, &body)
		if err != nil {
			return
		}
		hashID := backstage.DncodeIDKey(body.ID)
		// Invoke service
		err = app.ClerkService.RemoveGood(ctx, clerk.GoodRomoveParam{GoodID: int32(hashID)})
		// goods, err := app.BarterService.ListMyGoods(ctx, barter.ListMyGoodsParam{
		if err != nil {
			reponse.RespondWithError(c,
				common.NewError(common.ErrorCodeResourceOperateFail, err, common.WithMsg("invalid delete parameter")))
			return
		}

		reponse.RespondWithoutBody(c, http.StatusOK)
		// respondWithJSON(c, http.StatusOK, resp)
	}
}
