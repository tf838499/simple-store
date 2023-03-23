package clerk

import (
	"context"
	"simple-store/internal/adapter/repository/PostgresDB"
	"simple-store/internal/domain/common"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestClerkService_ListGoods(t *testing.T) {
	t.Parallel()

	// Args
	type Args struct {
		Page  GoodListParam
		Goods []PostgresDB.Good
	}

	var args Args
	_ = faker.FakeData(&args)

	// Init
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Test cases
	testCases := []struct {
		Name         string
		SetupService func(t *testing.T) *ClerkService
		ExpectError  bool
	}{
		{
			Name: "Success get goods",
			SetupService: func(t *testing.T) *ClerkService {
				mock := buildServiceMock(ctrl)
				Goods := args.Goods
				mock.GoodRepo.EXPECT().GetGoodListByPage(gomock.Any(), gomock.Any()).AnyTimes().Return(Goods, nil)
				// mock.GoodRepo.EXPECT().GetGoodListByPage(gomock.Any(), args.Good).Return(nil, common.DomainError{})
				service := buildService(mock)
				return service
			},
			ExpectError: false,
		},
		{
			Name: "Fail get goods",
			SetupService: func(t *testing.T) *ClerkService {
				mock := buildServiceMock(ctrl)
				mock.GoodRepo.EXPECT().GetGoodListByPage(gomock.Any(), gomock.Any()).AnyTimes().Return(nil, common.DomainError{})
				// mock.GoodRepo.EXPECT().GetGoodListByPage(gomock.Any(), args.Good).Return(nil, common.DomainError{})
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
			var param = GoodListParam{args.Page.Limit, args.Page.Offset}
			_, err := service.ListGoods(context.Background(), param)
			if c.ExpectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
func TestClerkService_AddGoods(t *testing.T) {
	t.Parallel()

	// Args
	type Args struct {
		Goods []PostgresDB.InsertGoodsParams
	}

	var args []Args

	// Init
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Test cases
	testCases := []struct {
		Name         string
		SetupService func(t *testing.T) *ClerkService
		ExpectError  bool
	}{
		{
			Name: "Success insert goods",
			SetupService: func(t *testing.T) *ClerkService {
				mock := buildServiceMock(ctrl)
				fakerArg := Args{Goods: []PostgresDB.InsertGoodsParams{}}
				_ = faker.FakeData(&fakerArg)
				args = append(args, fakerArg)
				mock.GoodRepo.EXPECT().InsertGoodsWithTx(gomock.Any(), args[0].Goods).Return(nil)
				// mock.GoodRepo.EXPECT().GetGoodListByPage(gomock.Any(), args.Good).Return(nil, common.DomainError{})
				service := buildService(mock)
				return service
			},
			ExpectError: false,
		},
		{
			Name: "Fail insert goods",
			SetupService: func(t *testing.T) *ClerkService {
				mock := buildServiceMock(ctrl)
				fakerArg := Args{Goods: []PostgresDB.InsertGoodsParams{}}
				// _ = faker.FakeData(&fakerArg)
				args = append(args, fakerArg)
				mock.GoodRepo.EXPECT().InsertGoodsWithTx(gomock.Any(), args[1].Goods).AnyTimes().Return(common.DomainError{})
				// mock.GoodRepo.EXPECT().GetGoodListByPage(gomock.Any(), args.Good).Return(nil, common.DomainError{})
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
			var param = args[i].Goods
			err := service.AddGoods(context.Background(), param)
			if c.ExpectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
func TestClerkService_ChangeGoods(t *testing.T) {
	t.Parallel()

	// Args
	type Args struct {
		Goods PostgresDB.UpdateGoodParams
	}

	var args []Args

	// Init
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Test cases
	testCases := []struct {
		Name         string
		SetupService func(t *testing.T) *ClerkService
		ExpectError  bool
	}{
		{
			Name: "Success change goods",
			SetupService: func(t *testing.T) *ClerkService {
				mock := buildServiceMock(ctrl)
				fakerArg := Args{Goods: PostgresDB.UpdateGoodParams{}}
				_ = faker.FakeData(&fakerArg)
				args = append(args, fakerArg)
				mock.GoodRepo.EXPECT().UpdateGood(gomock.Any(), args[0].Goods).Return(nil)
				// mock.GoodRepo.EXPECT().GetGoodListByPage(gomock.Any(), args.Good).Return(nil, common.DomainError{})
				service := buildService(mock)
				return service
			},
			ExpectError: false,
		},
		{
			Name: "Fail change goods",
			SetupService: func(t *testing.T) *ClerkService {
				mock := buildServiceMock(ctrl)
				fakerArg := Args{Goods: PostgresDB.UpdateGoodParams{}}
				// _ = faker.FakeData(&fakerArg)
				args = append(args, fakerArg)
				mock.GoodRepo.EXPECT().UpdateGood(gomock.Any(), args[1].Goods).AnyTimes().Return(common.DomainError{})
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
			var param = args[i].Goods
			err := service.ChangeGoods(context.Background(), param)
			if c.ExpectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestClerkService_RemoveGood(t *testing.T) {
	t.Parallel()

	// Args
	type Args struct {
		Goods GoodRomoveParam
	}

	var args []Args

	// Init
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Test cases
	testCases := []struct {
		Name         string
		SetupService func(t *testing.T) *ClerkService
		ExpectError  bool
	}{
		{
			Name: "Success delete goods",
			SetupService: func(t *testing.T) *ClerkService {
				mock := buildServiceMock(ctrl)
				fakerArg := Args{Goods: GoodRomoveParam{}}
				_ = faker.FakeData(&fakerArg)
				args = append(args, fakerArg)
				mock.GoodRepo.EXPECT().DeleteGood(gomock.Any(), gomock.Any()).Return(nil)
				// mock.GoodRepo.EXPECT().GetGoodListByPage(gomock.Any(), args.Good).Return(nil, common.DomainError{})
				service := buildService(mock)
				return service
			},
			ExpectError: false,
		},
		{
			Name: "Fail delete goods",
			SetupService: func(t *testing.T) *ClerkService {
				mock := buildServiceMock(ctrl)
				fakerArg := Args{Goods: GoodRomoveParam{}}
				// _ = faker.FakeData(&fakerArg)
				args = append(args, fakerArg)
				mock.GoodRepo.EXPECT().DeleteGood(gomock.Any(), gomock.Any()).AnyTimes().Return(common.DomainError{})
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
			var param = args[i].Goods
			err := service.RemoveGood(context.Background(), param)
			if c.ExpectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
