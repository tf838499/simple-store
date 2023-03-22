package cartlist

import "strconv"

type GoodInCarts struct {
	ImageName string
	Price     int
	Amount    int
}

func SpiltNamePrice(NameAndPrice string) (string, int) {
	index := findSplitIndex(NameAndPrice)
	name, PriceStr := NameAndPrice[:index], NameAndPrice[index:]

	Price, err := strconv.Atoi(PriceStr)
	if err != nil {
		return "", -1
	}
	return name, Price
}

func findSplitIndex(str string) int {
	r := len(str) - 1
	for i := r; i > 0; i-- {
		if str[i] == '_' {
			return i
		}
	}
	return -1
}

func CheckoutPrice(price map[string]int, amount []int32, name []string) int {
	totalPrice := 0
	for i := range name {
		totalPrice = totalPrice + int(amount[i])*price[name[i]]
	}
	return totalPrice
}
