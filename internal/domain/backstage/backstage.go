package backstage

import "time"

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
