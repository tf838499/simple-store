package customer

import (
	"context"
	"testing"

	"simple-store/internal/app/service/customer/automock"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
)

type serviceMock struct {
	CartRepo  *automock.MockCartRepository
	OrderRepo *automock.MockOrderRepository
}

func buildServiceMock(ctrl *gomock.Controller) serviceMock {
	return serviceMock{
		CartRepo:  automock.NewMockCartRepository(ctrl),
		OrderRepo: automock.NewMockOrderRepository(ctrl),
	}
}
func buildService(mock serviceMock) *CustomerService {
	param := CustomerServiceParam{
		CartRepo:  mock.CartRepo,
		OrderRepo: mock.OrderRepo,
	}
	return NewCustomerService(context.Background(), param)
}

// nolint
func TestMain(m *testing.M) {
	// To avoid getting an empty object slice
	_ = faker.SetRandomMapAndSliceMinSize(2)

	// To avoid getting a zero random number
	_ = faker.SetRandomNumberBoundaries(1, 100)

	m.Run()
}
