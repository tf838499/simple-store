package clerk

import (
	"context"

	"github.com/rs/zerolog"
)

type ClerkService struct {
	goodRepo GoodRepository
}

type ClerkServiceParam struct {
	GoodRepo GoodRepository
}

func NewClerkService(_ context.Context, param ClerkServiceParam) *ClerkService {
	return &ClerkService{
		goodRepo: param.GoodRepo,
	}
}
func (c *ClerkService) logger(ctx context.Context) *zerolog.Logger {
	l := zerolog.Ctx(ctx).With().Str("component", "clerk-service").Logger()
	return &l
}
