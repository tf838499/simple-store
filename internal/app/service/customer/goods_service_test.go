package customer

import (
	"context"
	"simple-store/internal/adapter/redisclient"
	"simple-store/internal/domain/common"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestClerkService_GetCartList(t *testing.T) {
	t.Parallel()

	// Args
	type Args struct {
		Goods      CartParams
		GoodInCart []redisclient.GoodInRedisParams
		FakePrice  []int
	}

	var args Args

	// Init
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Test cases
	testCases := []struct {
		Name         string
		SetupService func(t *testing.T) *CustomerService
		ExpectError  bool
	}{
		{
			Name: "Success get goods",
			SetupService: func(t *testing.T) *CustomerService {
				mock := buildServiceMock(ctrl)
				// fakerArg := Args}
				_ = faker.FakeData(&args)
				// args = fakerArg)
				// args.GoodInCart
				mock.CartRepo.EXPECT().GetCartListCache(gomock.Any(), gomock.Any()).Return(args.GoodInCart, nil)
				mock.CartRepo.EXPECT().GetGoodPrice(gomock.Any(), gomock.Any()).Return(args.FakePrice, nil)
				service := buildService(mock)
				return service
			},
			ExpectError: false,
		},
		{
			Name: "Fail get goods",
			SetupService: func(t *testing.T) *CustomerService {
				mock := buildServiceMock(ctrl)
				// fakerArg := Args}
				_ = faker.FakeData(&args)
				// args = fakerArg)
				// args.GoodInCart
				mock.CartRepo.EXPECT().GetCartListCache(gomock.Any(), gomock.Any()).Return(nil, common.DomainError{})
				mock.CartRepo.EXPECT().GetGoodPrice(gomock.Any(), gomock.Any()).Return(args.FakePrice, nil)
				service := buildService(mock)
				return service
			},
			ExpectError: true,
		},
	}

	for i := range testCases {
		c := testCases[i]
		t.Run(c.Name, func(t *testing.T) {
			service := c.SetupService(t)
			var param = args.Goods
			_, err := service.GetCartList(context.Background(), param)

			if c.ExpectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestClerkService_SetGoodInCart(t *testing.T) {
	t.Parallel()

	// Args
	type Args struct {
		Goods      CartParams
		GoodInCart []redisclient.GoodInCartParams
		FakePrice  []int
	}

	var args Args

	// Init
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Test cases
	testCases := []struct {
		Name         string
		SetupService func(t *testing.T) *CustomerService
		ExpectError  bool
	}{
		{
			Name: "Success get goods",
			SetupService: func(t *testing.T) *CustomerService {
				mock := buildServiceMock(ctrl)
				// fakerArg := Args}
				_ = faker.FakeData(&args)
				// args = fakerArg)
				// args.GoodInCart
				mock.CartRepo.EXPECT().SetGood(gomock.Any(), gomock.Any()).Return(nil)

				service := buildService(mock)
				return service
			},
			ExpectError: false,
		},
		{
			Name: "Fail get goods",
			SetupService: func(t *testing.T) *CustomerService {
				mock := buildServiceMock(ctrl)
				// fakerArg := Args}
				_ = faker.FakeData(&args)
				// args = fakerArg)
				// args.GoodInCart
				mock.CartRepo.EXPECT().GetCartListCache(gomock.Any(), gomock.Any()).Return(nil, common.DomainError{})
				mock.CartRepo.EXPECT().GetGoodPrice(gomock.Any(), gomock.Any()).Return(args.FakePrice, nil)
				service := buildService(mock)
				return service
			},
			ExpectError: true,
		},
	}

	for i := range testCases {
		c := testCases[i]
		t.Run(c.Name, func(t *testing.T) {
			service := c.SetupService(t)
			var param = args.Goods
			_, err := service.GetCartList(context.Background(), param)

			if c.ExpectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
