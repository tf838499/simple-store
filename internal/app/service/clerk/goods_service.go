package clerk

import (
	"context"
	"fmt"
	"simple-store/internal/domain/backstage"
)

type PostGoodParam struct {
	GoodName string
}

func (s *ClerkService) ListMyGoods(ctx context.Context, param PostGoodParam) ([]backstage.Good, error) {
	goods, err := s.goodRepo.GetGoodList(ctx)
	if err != nil {
		return nil, err
	}
	var alterGoods []backstage.Good
	for i := range goods {
		g := goods[i]
		alterGoods = append(alterGoods, backstage.Good{
			ID:   int(g.ID),
			Name: g.ImageName.String,
		})
	}

	fmt.Println(param.GoodName)
	fmt.Println(goods)
	return alterGoods, err
}
