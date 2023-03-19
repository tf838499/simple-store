package testdata

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
)

var basepath string

const (
	// TestDataTrader = "trader.yml"
	TestDataGood = "goods.yml"
)

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	basepath = filepath.Dir(currentFile)
}

func Path(rel string) string {
	return filepath.Join(basepath, rel)
}
func Data(rel string) []uint8 {
	yfile, _ := ioutil.ReadFile(basepath + "/" + rel)
	return yfile
}
