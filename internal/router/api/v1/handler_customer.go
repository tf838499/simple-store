package v1

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"simple-store/internal/adapter/redisclient"
	"simple-store/internal/app"
	"simple-store/internal/app/service/customer"
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

	type Body struct {
		Email      string   `json:"email" form:"email" binding:"required"`
		GoodAmount []int32  `json:"good_amount" form:"good_amount" binding:"required"`
		GoodName   []string `json:"good_name" form:"good_name" binding:"required"`
		Message    string   `json:"message" form:"message" binding:"required"`
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
		var InserOrderParm customer.OrderParams
		InserOrderParm.Email = body.Email
		InserOrderParm.GoodAmount = body.GoodAmount
		InserOrderParm.GoodName = body.GoodName
		InserOrderParm.Message = body.Message

		OrderInfo, err := app.CustomerService.InsertGoodInCart(ctx, InserOrderParm)
		if err != nil {
			log.Panicf(err.Error())
			msg := "fail create order"
			reponse.RespondWithError(c,
				common.NewError(common.ErrorCodeResourceNotFound, errors.New(msg), common.WithMsg(msg)))
			return
		}
		resp := OrderInfo
		reponse.RespondWithJSON(c, http.StatusOK, resp)
	}
}

// @Summary Customer Add Cart
// @Tags Customer
// @version 1.0
// @produce application/json
// @param email Json string true "email"
// @param good_price Json int true "good_price"
// @param good_amount Json int true "good_amount"
// @param good_name Json string true "good_name"
// @Success 200 string string "success"
// @Failure 404 {page} errcode.error "no_found_item"
// @Failure 400 {page} errcode.error "invalid_parameter"
// @Router api/v1/clerk/goods [get]
func AddCartGoods(app *app.Application) gin.HandlerFunc {

	type Body struct {
		Email      string `json:"email" form:"email" binding:"required"`
		GoodPrice  int    `json:"good_price" form:"good_price" binding:"required"`
		GoodAmount int    `json:"good_amount" form:"good_amount" binding:"required"`
		GoodName   string `json:"good_name" form:"good_name" binding:"required"`
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
		var goodsParams []redisclient.GoodInCartParams
		for i := range body {
			goodsParams = append(goodsParams, redisclient.GoodInCartParams{
				CustomerID: body[i].Email,
				GoodName:   body[i].GoodName,
				GoodPrice:  body[i].GoodPrice,
				GoodAmount: body[i].GoodAmount,
			})
		}

		// Invoke service
		err = app.CustomerService.SetGoodInCart(ctx, goodsParams)
		if err != nil {
			log.Panicf(err.Error())
			msg := "fail insert to cart"
			reponse.RespondWithError(c,
				common.NewError(common.ErrorCodeResourceNotFound, errors.New(msg), common.WithMsg(msg)))
			return
		}
		reponse.RespondWithoutBody(c, http.StatusOK)
	}
}

// @Summary Customer Delete Cart
// @Tags Customer
// @version 1.0
// @produce application/json
// @param email Json string true "email"
// @param good_price Json int true "good_price"
// @param good_amount Json int true "good_amount"
// @param good_name Json string true "good_name"
// @Success 200 string string "success"
// @Failure 404 {page} errcode.error "no_found_item"
// @Failure 400 {page} errcode.error "invalid_parameter"
// @Router api/v1/clerk/goods [get]
func DeleteCartGoods(app *app.Application) gin.HandlerFunc {

	type Body struct {
		Email      string `json:"email" form:"email" binding:"required"`
		GoodPrice  int    `json:"good_price" form:"good_price" binding:"required"`
		GoodAmount int    `json:"good_amount" form:"good_amount" binding:"required"`
		GoodName   string `json:"good_name" form:"good_name" binding:"required"`
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
		var goodsParams []redisclient.GoodInCartParams
		for i := range body {
			goodsParams = append(goodsParams, redisclient.GoodInCartParams{
				CustomerID: body[i].Email,
				GoodName:   body[i].GoodName,
				GoodPrice:  body[i].GoodPrice,
				GoodAmount: body[i].GoodAmount,
			})
		}

		// Invoke service
		err = app.CustomerService.DeleteGoodInCart(ctx, goodsParams)
		if err != nil {
			log.Panicf(err.Error())
			msg := "fail insert to cart"
			reponse.RespondWithError(c,
				common.NewError(common.ErrorCodeResourceNotFound, errors.New(msg), common.WithMsg(msg)))
			return
		}
		reponse.RespondWithoutBody(c, http.StatusOK)
	}
}

// @Summary Customer Delete Cart
// @Tags Customer
// @version 1.0
// @produce application/json
// @param email Json string true "email"
// @param good_price Json int true "good_price"
// @param good_amount Json int true "good_amount"
// @param good_name Json string true "good_name"
// @Success 200 string string "success"
// @Failure 404 {page} errcode.error "no_found_item"
// @Failure 400 {page} errcode.error "invalid_parameter"
// @Router api/v1/clerk/goods [get]
func CartLists(app *app.Application) gin.HandlerFunc {
	type GoodInCart struct {
		ImageName string `json:"image_name"`
		Amount    int    `json:"amount"`
		Price     int    `json:"price"`
	}
	type Response struct {
		Goods []GoodInCart `json:"goods"`
	}

	type Body struct {
		Email string `json:"email" form:"email" binding:"required"`
	}

	return func(c *gin.Context) {
		ctx := c.Request.Context()
		var body Body
		// errors.create(arg == 10)
		// fmt.Scanln()
		err := c.ShouldBind(&body)
		if err != nil {
			log.Panicf(err.Error())
			reponse.RespondWithError(c,
				common.NewError(common.ErrorCodeParameterInvalid, err, common.WithMsg("invalid parameter")))
			return
		}

		// Invoke service
		CartGodds, err := app.CustomerService.GetCartList(ctx, customer.CartParams{Email: body.Email})
		if err != nil {
			log.Panicf(err.Error())
			msg := "fail get goods in cart"
			reponse.RespondWithError(c,
				common.NewError(common.ErrorCodeResourceNotFound, errors.New(msg), common.WithMsg(msg)))
			return
		}
		resp := Response{Goods: []GoodInCart{}}
		for i := range CartGodds {
			g := CartGodds[i]
			resp.Goods = append(resp.Goods, GoodInCart{
				ImageName: g.ImageName,
				Amount:    g.Amount,
				Price:     g.Amount,
			})
		}
		reponse.RespondWithJSON(c, http.StatusOK, CartGodds)
	}
}
