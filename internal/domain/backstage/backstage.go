package backstage

import (
	"time"

	"github.com/speps/go-hashids/v2"
)

type Good struct {
	ID        int
	CreatedAt time.Time
	ImageName string
	Descript  string
	Price     int
	Class     string
}

func NewGood(goodName string) Good {
	return Good{
		ImageName: goodName,
	}
}
func AlterGood(goodName string) Good {
	return Good{
		ImageName: goodName,
	}
}

var Salt = "simpleStorePkId"
var keylength = 16

func EncodeIDKey(id int) string {
	hd := hashids.NewData()
	hd.Salt = Salt
	hd.MinLength = keylength
	h, err := hashids.NewWithData(hd)
	if err != nil {
		// add log
		return ""
	}

	e, err := h.Encode([]int{id})
	if err != nil {
		// add log
		return ""
	}
	return e
}
func DncodeIDKey(id string) int {
	hd := hashids.NewData()
	hd.Salt = Salt
	hd.MinLength = keylength
	h, err := hashids.NewWithData(hd)
	if err != nil {
		// add log
		return -1
	}

	e, err := h.DecodeWithError(id)
	if err != nil {
		// add log
		return -1
	}
	return e[0]
}
