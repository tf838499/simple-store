package customer

import (
	"context"

	"github.com/rs/zerolog"
)

type CustomerService struct {
	cartRepo  CartRepository
	orderRepo OrderRepository
	authRepo  AuthRepository
}

type CustomerServiceParam struct {
	CartRepo  CartRepository
	OrderRepo OrderRepository
	AuthRepo  AuthRepository
}

func NewCustomerService(_ context.Context, param CustomerServiceParam) *CustomerService {
	return &CustomerService{
		cartRepo:  param.CartRepo,
		orderRepo: param.OrderRepo,
		authRepo:  param.AuthRepo,
	}
}
func (c *CustomerService) logger(ctx context.Context) *zerolog.Logger {
	l := zerolog.Ctx(ctx).With().Str("component", "customer-service").Logger()
	return &l
}
