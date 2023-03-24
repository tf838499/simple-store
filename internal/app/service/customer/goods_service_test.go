package customer

import (
	"context"
	"fmt"
	"simple-store/internal/adapter/redisclient"
	"simple-store/internal/adapter/repository/PostgresDB"
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
		GoodInCart []redisclient.GoodInCartParams
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
			Name: "Success insert goods",
			SetupService: func(t *testing.T) *CustomerService {
				mock := buildServiceMock(ctrl)
				// fakerArg := Args}
				_ = faker.FakeData(&args)
				// args = fakerArg)
				// args.GoodInCart
				mock.CartRepo.EXPECT().SetGood(gomock.Any(), gomock.Any()).Return(nil)
				mock.CartRepo.EXPECT().MSetGoodPrice(gomock.Any(), gomock.Any()).Return(nil)
				service := buildService(mock)
				return service
			},
			ExpectError: false,
		},
		{
			Name: "fail insert goods",
			SetupService: func(t *testing.T) *CustomerService {
				mock := buildServiceMock(ctrl)
				// fakerArg := Args}
				_ = faker.FakeData(&args)
				// args = fakerArg)
				// args.GoodInCart
				mock.CartRepo.EXPECT().SetGood(gomock.Any(), gomock.Any()).Return(common.DomainError{})
				mock.CartRepo.EXPECT().MSetGoodPrice(gomock.Any(), gomock.Any()).Return(nil)
				service := buildService(mock)
				return service
			},
			ExpectError: true,
		},
		{
			Name: "fail insert goods price",
			SetupService: func(t *testing.T) *CustomerService {
				mock := buildServiceMock(ctrl)
				// fakerArg := Args}
				_ = faker.FakeData(&args)
				// args = fakerArg)
				// args.GoodInCart
				mock.CartRepo.EXPECT().SetGood(gomock.Any(), gomock.Any()).Return(nil)
				mock.CartRepo.EXPECT().MSetGoodPrice(gomock.Any(), gomock.Any()).Return(common.DomainError{})
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
			var param = args.GoodInCart
			err := service.SetGoodInCart(context.Background(), param)

			if c.ExpectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestClerkService_DeleteGoodInCart(t *testing.T) {
	t.Parallel()

	// Args
	type Args struct {
		GoodInCart []redisclient.GoodInCartParams
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
			Name: "Success delete goods",
			SetupService: func(t *testing.T) *CustomerService {
				mock := buildServiceMock(ctrl)
				// fakerArg := Args}
				_ = faker.FakeData(&args)
				// args = fakerArg)
				// args.GoodInCart
				mock.CartRepo.EXPECT().DeleteGood(gomock.Any(), gomock.Any()).Return(nil)
				service := buildService(mock)
				return service
			},
			ExpectError: false,
		},
		{
			Name: "fail delete goods",
			SetupService: func(t *testing.T) *CustomerService {
				mock := buildServiceMock(ctrl)
				// fakerArg := Args}
				_ = faker.FakeData(&args)
				// args = fakerArg)
				// args.GoodInCart
				mock.CartRepo.EXPECT().DeleteGood(gomock.Any(), gomock.Any()).Return(common.DomainError{})
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
			var param = args.GoodInCart
			err := service.DeleteGoodInCart(context.Background(), param)
			if c.ExpectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestClerkService_InsertGoodInCart(t *testing.T) {
	t.Parallel()

	// Args
	type Args struct {
		GoodInCart OrderParams
		PriceList  []int
		GoodInfo   PostgresDB.Good
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
			Name: "Success create order",
			SetupService: func(t *testing.T) *CustomerService {
				mock := buildServiceMock(ctrl)
				// fakerArg := Args}
				_ = faker.FakeData(&args)
				args.GoodInCart.GoodAmount = args.GoodInCart.GoodAmount[:2]
				args.GoodInCart.GoodName = args.GoodInCart.GoodName[:2]
				args.PriceList = args.PriceList[:2]
				// args = fakerArg)
				// args.GoodInCart
				mock.CartRepo.EXPECT().GetGoodPrice(gomock.Any(), gomock.Any()).Return(args.PriceList, nil)
				mock.OrderRepo.EXPECT().InsertOrder(gomock.Any(), gomock.Any()).Return(nil)
				service := buildService(mock)
				return service
			},
			ExpectError: false,
		},
		{
			Name: "Success create order & miss price",
			SetupService: func(t *testing.T) *CustomerService {
				mock := buildServiceMock(ctrl)
				// fakerArg := Args}
				_ = faker.FakeData(&args)
				args.GoodInCart.GoodAmount = args.GoodInCart.GoodAmount[:2]
				args.GoodInCart.GoodName = args.GoodInCart.GoodName[:2]
				args.PriceList = args.PriceList[:2]
				args.PriceList[0] = -1
				// args = fakerArg)
				// args.GoodInCart
				mock.CartRepo.EXPECT().GetGoodPrice(gomock.Any(), gomock.Any()).Return(args.PriceList, nil)
				mock.OrderRepo.EXPECT().GetGoodByName(gomock.Any(), gomock.Any()).Return(args.GoodInfo, nil)
				mock.CartRepo.EXPECT().SetGoodPrice(gomock.Any(), gomock.Any()).Return(nil)
				mock.OrderRepo.EXPECT().InsertOrder(gomock.Any(), gomock.Any()).Return(nil)
				service := buildService(mock)
				return service
			},
			ExpectError: false,
		},
		{
			Name: "Fail create order : get good price",
			SetupService: func(t *testing.T) *CustomerService {
				mock := buildServiceMock(ctrl)
				// fakerArg := Args}
				_ = faker.FakeData(&args)
				args.GoodInCart.GoodAmount = args.GoodInCart.GoodAmount[:2]
				args.GoodInCart.GoodName = args.GoodInCart.GoodName[:2]
				args.PriceList = args.PriceList[:2]
				args.PriceList[0] = -1
				// args = fakerArg)
				// args.GoodInCart
				mock.CartRepo.EXPECT().GetGoodPrice(gomock.Any(), gomock.Any()).Return(nil, common.DomainError{})
				service := buildService(mock)
				return service
			},
			ExpectError: true,
		},
		{
			Name: "Fail create order : get good by name",
			SetupService: func(t *testing.T) *CustomerService {
				mock := buildServiceMock(ctrl)
				// fakerArg := Args}
				_ = faker.FakeData(&args)
				args.GoodInCart.GoodAmount = args.GoodInCart.GoodAmount[:2]
				args.GoodInCart.GoodName = args.GoodInCart.GoodName[:2]
				args.PriceList = args.PriceList[:2]
				args.PriceList[0] = -1
				// args = fakerArg)
				// args.GoodInCart
				mock.CartRepo.EXPECT().GetGoodPrice(gomock.Any(), gomock.Any()).Return(args.PriceList, nil)
				mock.OrderRepo.EXPECT().GetGoodByName(gomock.Any(), gomock.Any()).Return(args.GoodInfo, common.DomainError{})
				service := buildService(mock)
				return service
			},
			ExpectError: true,
		},
		{
			Name: "Fail create order : set good price by name",
			SetupService: func(t *testing.T) *CustomerService {
				mock := buildServiceMock(ctrl)
				// fakerArg := Args}
				_ = faker.FakeData(&args)
				args.GoodInCart.GoodAmount = args.GoodInCart.GoodAmount[:2]
				args.GoodInCart.GoodName = args.GoodInCart.GoodName[:2]
				args.PriceList = args.PriceList[:2]
				args.PriceList[0] = -1
				// args = fakerArg)
				// args.GoodInCart
				mock.CartRepo.EXPECT().GetGoodPrice(gomock.Any(), gomock.Any()).Return(args.PriceList, nil)
				mock.OrderRepo.EXPECT().GetGoodByName(gomock.Any(), gomock.Any()).Return(args.GoodInfo, nil)
				mock.CartRepo.EXPECT().SetGoodPrice(gomock.Any(), gomock.Any()).Return(common.DomainError{})
				mock.OrderRepo.EXPECT().InsertOrder(gomock.Any(), gomock.Any()).Return(nil)
				service := buildService(mock)
				return service
			},
			ExpectError: false,
		},
		{
			Name: "Fail create order : InsertOrder",
			SetupService: func(t *testing.T) *CustomerService {
				mock := buildServiceMock(ctrl)
				// fakerArg := Args}
				_ = faker.FakeData(&args)
				args.GoodInCart.GoodAmount = args.GoodInCart.GoodAmount[:2]
				args.GoodInCart.GoodName = args.GoodInCart.GoodName[:2]
				args.PriceList = args.PriceList[:2]
				args.PriceList[0] = -1
				// args = fakerArg)
				// args.GoodInCart
				mock.CartRepo.EXPECT().GetGoodPrice(gomock.Any(), gomock.Any()).Return(args.PriceList, nil)
				mock.OrderRepo.EXPECT().GetGoodByName(gomock.Any(), gomock.Any()).Return(args.GoodInfo, nil)
				mock.CartRepo.EXPECT().SetGoodPrice(gomock.Any(), gomock.Any()).Return(nil)
				mock.OrderRepo.EXPECT().InsertOrder(gomock.Any(), gomock.Any()).Return(common.DomainError{})
				service := buildService(mock)
				return service
			},
			ExpectError: true,
		},
	}

	for i := range testCases {
		c := testCases[i]
		fmt.Println(i)
		t.Run(c.Name, func(t *testing.T) {
			service := c.SetupService(t)
			var param = args.GoodInCart
			_, err := service.InsertGoodInCart(context.Background(), param)
			if c.ExpectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
