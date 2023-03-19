package customer

import (
	"context"

	"github.com/rs/zerolog"
)

type CustomerService struct {
	cartRepo CartRepository
}

type CustomerServiceParam struct {
	CartRepo CartRepository
}

func NewCustomerService(_ context.Context, param CustomerServiceParam) *CustomerService {
	return &CustomerService{
		cartRepo: param.CartRepo,
	}
}
func (c *CustomerService) logger(ctx context.Context) *zerolog.Logger {
	l := zerolog.Ctx(ctx).With().Str("component", "customer-service").Logger()
	return &l
}
