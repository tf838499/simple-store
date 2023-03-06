package backstage

type Good struct {
	ID   int
	Name string
}

func NewGood(goodName string) Good {
	return Good{
		Name: goodName,
	}
}
func AlterGood(goodName string) Good {
	return Good{
		Name: goodName,
	}
}
