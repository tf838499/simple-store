package RedisCache

import (
	"context"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/require"
	"gotest.tools/assert"
)

func TestRedisReposity_SetGood(t *testing.T) {
	db := getTestRedis()
	repo := initRepository(t, db)
	type Args struct {
		GoodFakes []GoodInCartParams
	}
	var args Args
	_ = faker.FakeData(&args)
	err := repo.SetGood(context.Background(), args.GoodFakes)
	require.NoError(t, err)
	var GoodNilFake []GoodInCartParams
	err = repo.SetGood(context.Background(), GoodNilFake)
	require.Error(t, err)
}
func TestRedisReposity_DeleteGood(t *testing.T) {
	db := getTestRedis()
	repo := initRepository(t, db)
	type Args struct {
		GoodFakes []GoodInCartParams
	}
	var args Args
	args.GoodFakes = []GoodInCartParams{
		{
			CustomerID: "testID1",
			Name:       "testGoodName1",
			Price:      100,
			Amount:     2,
		},
		{
			CustomerID: "testID1",
			Name:       "testGoodName2",
			Price:      20,
			Amount:     2,
		},
	}
	err := repo.SetGood(context.Background(), args.GoodFakes)
	require.NoError(t, err)
	err = repo.DeleteGood(context.Background(), args.GoodFakes)
	require.NoError(t, err)
}
func assertCartList(t *testing.T, expected []GoodInCartParams, actual []GoodInCartParams) {
	assert.Equal(t, len(expected), len(actual))
	for i := 0; i < len(expected); i++ {
		require.NotNil(t, actual)
		assert.Equal(t, expected[i].CustomerID, actual[i].CustomerID)
		assert.Equal(t, expected[i].Amount, actual[i].Amount)
		assert.Equal(t, expected[i].Name, actual[i].Name)
		assert.Equal(t, expected[i].Price, actual[i].Price)
		// assert.Equal(t, expected[i].Descript.String, actual[i].Descript.String)
		// assert.Equal(t, expected[i].ImageName.String, actual[i].ImageName.String)
		// assert.Equal(t, expected[i].Price.Int64, actual[i].Price.Int64)
	}
}
func TestRedisReposity_GetCartListCache(t *testing.T) {
	db := getTestRedis()
	repo := initRepository(t, db)
	type Args struct {
		GoodFakes []GoodInCartParams
	}
	var args Args
	args.GoodFakes = []GoodInCartParams{
		{
			CustomerID: "testID1",
			Name:       "testGoodName1",
			Price:      100,
			Amount:     2,
		},
		{
			CustomerID: "testID1",
			Name:       "testGoodName2",
			Price:      20,
			Amount:     2,
		},
	}
	err := repo.SetGood(context.Background(), args.GoodFakes)
	require.NoError(t, err)
	CartList, err := repo.GetCartListCache(context.Background(), "testID1")
	require.NoError(t, err)
	assertCartList(t, CartList, args.GoodFakes)
}
