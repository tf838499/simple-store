// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: store_repository.sql

package PostgresDB

import (
	"context"
	"database/sql"
	"simple-store/testdata"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)
func assertGoodsList(t *testing.T, expected []InsertGoodsParams, actual []Good) {
	assert.Equal(t,len(expected),len(actual))
	for i:= 0;i< len(expected);i++{
		require.NotNil(t, actual)
		assert.Equal(t, expected[i].Class.String , actual[i].Class.String)
		assert.Equal(t, expected[i].Descript.String , actual[i].Descript.String)
		assert.Equal(t, expected[i].ImageName.String, actual[i].ImageName.String)
		assert.Equal(t, expected[i].Price.Int64, actual[i].Price.Int64)
	}
}
type GoodFake struct{
	ID        int32          
	Image_Name string 
	Class string
	Descript  string 
	Price     int64
}
func TestStoreReposity_GetGoodList(t *testing.T) {
	db := getTestPostgresDB()
	repo := initRepository(t, db)
	// repo := initRepository(t, db,testdata.Path(testdata.TestDataGood))
	type Args struct {
		Page GetGoodListByPageParams
		GoodFakes []InsertGoodsParams
	}
	var args Args	
	fakedata:=testdata.Data(testdata.TestDataGood)
	// var c []GoodFake
	yaml.Unmarshal(fakedata, &args.GoodFakes)
	// var args Args	
	_ = faker.FakeData(&args)
	err := repo.InsertGoodsWithTx(context.Background(),args.GoodFakes)
	require.NoError(t, err)
	// assert.AnError
	args.Page.Limit = 15
	args.Page.Offset = 1
	goods, err := repo.GetGoodListByPage(context.Background(),args.Page) 
	require.NoError(t, err)
	assertGoodsList(t, args.GoodFakes[args.Page.Offset:args.Page.Offset+args.Page.Limit], goods)
}

func TestStoreReposity_InsertGoods(t *testing.T) {
	db := getTestPostgresDB()
	repo := initRepository(t, db)
	// repo := initRepository(t, db,testdata.Path(testdata.TestDataGood))
	type Args struct {
		Page GetGoodListByPageParams
		GoodFakes []InsertGoodsParams
	}
	var args Args	
	// fakedata:=testdata.Data(testdata.TestDataGood)
	// // var c []GoodFake
	// yaml.Unmarshal(fakedata, &args.GoodFakes)
	// var args Args	
	
	err := repo.InsertGoodsWithTx(context.Background(),args.GoodFakes)
	require.NoError(t, err)

	_ = faker.FakeData(&args)
	err = repo.InsertGoodsWithTx(context.Background(),args.GoodFakes)
	require.NoError(t, err)

}

func TestStoreReposity_UpdateGood(t *testing.T) {
	db := getTestPostgresDB()
	repo := initRepository(t, db,testdata.Path(testdata.TestDataGood))
	type Args struct {
		Page GetGoodListByPageParams
		GoodFakes UpdateGoodParams
	}
	var args Args	
	args.GoodFakes = UpdateGoodParams{
		ImageName: sql.NullString{String: "test",Valid: true},
		Descript: sql.NullString{String: "test",Valid: true},
		Price: sql.NullInt64{Int64: 10,Valid: true},
		Class: sql.NullString{String: "test",Valid: true},
		ID: 1,
	}
	err := repo.UpdateGood(context.Background(),args.GoodFakes)
	require.NoError(t, err)
}
func TestStoreReposity_DeleteGood(t *testing.T) {
	db := getTestPostgresDB()
	repo := initRepository(t, db,testdata.Path(testdata.TestDataGood))
	type Args struct {
		Page GetGoodListByPageParams
		GoodFakesID int32
	}
	var args Args	
	args.GoodFakesID = 1
	
	err := repo.DeleteGood(context.Background(),args.GoodFakesID)
	require.NoError(t, err)
}