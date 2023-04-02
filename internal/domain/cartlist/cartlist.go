package cartlist

type GoodInCart struct {
	Name   []string
	Price  []int
	Amount []int
}

func CheckoutPrice(price []int, amount []int32, name []string) int {
	totalPrice := 0
	for i := range name {
		totalPrice = totalPrice + int(amount[i])*price[i]
		if price[i] == 0 {
			return -1
		}
	}
	return totalPrice
}
