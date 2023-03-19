package backstage

import (
	"fmt"
	"testing"

	"github.com/speps/go-hashids/v2"
)

func TestSymmetricKey(t *testing.T) {
	hd := hashids.NewData()
	hd.Salt = "simpleStorePkId"
	// hd.MinLength = 8
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{99911394939, 023445})
	fmt.Println(e)
	d, _ := h.DecodeWithError(e)
	fmt.Println(d)
}
