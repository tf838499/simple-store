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

func TestBarterService_RegisterTrader(t *testing.T) {
	t.Parallel()

	// Args
	type Args struct {
		Good []PostgresDB.Good
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
			Name: "trader does not exist",
			SetupService: func(t *testing.T) *ClerkService {
				mock := buildServiceMock(ctrl)
				mock.GoodRepo.EXPECT().GetGoodListByPage(gomock.Any(), args.Good).Return(nil, common.DomainError{})
				// mock.TraderRepo.EXPECT().GetTraderByEmail(gomock.Any(), args.Trader.Email).Return(nil, common.DomainError{})
				// mock.AuthServer.EXPECT().RegisterAccount(gomock.Any(), args.Trader.Email, args.Password).Return(args.Trader.UID, nil)
				// mock.TraderRepo.EXPECT().CreateTrader(gomock.Any(), gomock.Any()).Return(&args.Trader, nil)

				service := buildService(mock)
				return service
			},
			ExpectError: false,
		},
		// {
		// 	Name: "failed to register trader",
		// 	SetupService: func(t *testing.T) *AuthService {
		// 		mock := buildServiceMock(ctrl)

		// 		mock.TraderRepo.EXPECT().GetTraderByEmail(gomock.Any(), args.Trader.Email).Return(nil, common.DomainError{})
		// 		mock.AuthServer.EXPECT().RegisterAccount(gomock.Any(), args.Trader.Email, args.Password).Return(args.Trader.UID, nil)
		// 		mock.TraderRepo.EXPECT().CreateTrader(gomock.Any(), gomock.Any()).Return(nil, common.DomainError{})

		// 		service := buildService(mock)
		// 		return service
		// 	},
		// 	ExpectError: true,
		// },
		// {
		// 	Name: "failed to register account",
		// 	SetupService: func(t *testing.T) *AuthService {
		// 		mock := buildServiceMock(ctrl)
		// 		mock.GoodRepo.EXPECT().GetGoodListByPage(gomock.Any(),args.Good)
		// 		mock.GoodRepo.EXPECT().InsertGoods()

		// 		// mock.TraderRepo.EXPECT().GetTraderByEmail(gomock.Any(), args.Trader.Email).Return(nil, common.DomainError{})
		// 		// mock.AuthServer.EXPECT().RegisterAccount(gomock.Any(), args.Trader.Email, args.Password).Return("", common.DomainError{})

		// 		service := buildService(mock)
		// 		return service
		// 	},
		// 	ExpectError: true,
		// },
		// {
		// 	Name: "trader exist",
		// 	SetupService: func(t *testing.T) *AuthService {
		// 		mock := buildServiceMock(ctrl)

		// 		mock.TraderRepo.EXPECT().GetTraderByEmail(gomock.Any(), args.Trader.Email).Return(&args.Trader, nil)

		// 		service := buildService(mock)
		// 		return service
		// 	},
		// 	ExpectError: true,
		// },
	}

	for i := range testCases {
		c := testCases[i]
		t.Run(c.Name, func(t *testing.T) {
			service := c.SetupService(t)
			param := RegisterTraderParam{
				Email:    args.Trader.Email,
				Name:     args.Trader.Name,
				Password: args.Password,
			}

			_, err := service.RegisterTrader(context.Background(), param)

			if c.ExpectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
