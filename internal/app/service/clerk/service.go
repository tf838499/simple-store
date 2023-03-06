package clerk

import "context"

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
