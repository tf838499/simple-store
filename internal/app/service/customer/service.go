package customer

import (
	"context"

	"github.com/rs/zerolog"
)

type CustomerService struct {
	cartRepo  CartRepository
	orderRepo OrderRepository
}

type CustomerServiceParam struct {
	CartRepo  CartRepository
	OrderRepo OrderRepository
}

func NewCustomerService(_ context.Context, param CustomerServiceParam) *CustomerService {
	return &CustomerService{
		cartRepo:  param.CartRepo,
		orderRepo: param.OrderRepo,
	}
}
func (c *CustomerService) logger(ctx context.Context) *zerolog.Logger {
	l := zerolog.Ctx(ctx).With().Str("component", "customer-service").Logger()
	return &l
}
