package clerk

import (
	"context"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"

	"simple-store/internal/app/service/clerk/automock"
)

type serviceMock struct {
	GoodRepo *automock.MockGoodRepository
}

func buildServiceMock(ctrl *gomock.Controller) serviceMock {
	return serviceMock{
		GoodRepo: automock.NewMockGoodRepository(ctrl),
	}
}
func buildService(mock serviceMock) *ClerkService {
	param := ClerkServiceParam{
		GoodRepo: mock.GoodRepo,
	}
	return NewClerkService(context.Background(), param)
}

// nolint
func TestMain(m *testing.M) {
	// To avoid getting an empty object slice
	_ = faker.SetRandomMapAndSliceMinSize(2)

	// To avoid getting a zero random number
	_ = faker.SetRandomNumberBoundaries(1, 100)

	m.Run()
}
