package testdata_test

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

func TestViper(t *testing.T) {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config")
	err := viper.ReadInConfig()
	if err != nil {
		print(err.Error())
	}
	// if err != nil {
	// 	panic("讀取設定檔出現錯誤，原因為：" + err.Errror())
	// }
	p := viper.Get("application.port")
	fmt.Println(p)
	fmt.Println("application port = " + viper.GetString("application.port"))
}
