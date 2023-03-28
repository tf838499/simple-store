package cartlist

import (
	"testing"

	"gotest.tools/assert"
)

func TestCheckoutPrice(t *testing.T) {
	type Args struct {
		price  []int
		amount []int32
		name   []string
	}
	var args Args
	tests := []struct {
		name         string
		SetupArgs    func(t *testing.T) Args
		ExpectResult int
	}{
		// TODO: Add test cases.
		{
			name: "good[1,2,3] amount[1,1,1] goodprice[10,20,30] total price=60",
			SetupArgs: func(t *testing.T) Args {
				a := args
				a.name = []string{"good1", "good2", "good3"}
				a.amount = []int32{1, 1, 1}
				a.price = []int{10, 20, 30}
				return a
			},
			ExpectResult: 60,
		},
		{
			name: "good[1,2,3] amount[1,2,1] goodprice[10,20,30] total price=80",
			SetupArgs: func(t *testing.T) Args {
				a := args
				a.name = []string{"good1", "good2", "good3"}
				a.amount = []int32{1, 2, 1}
				a.price = []int{10, 20, 30}
				return a
			},
			ExpectResult: 80,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := CheckoutPrice(tt.args.price, tt.args.amount, tt.args.name); got != tt.want {
			// 	t.Errorf("CheckoutPrice() = %v, want %v", got, tt.want)
			// }
			a := tt.SetupArgs(t)
			total := CheckoutPrice(a.price, a.amount, a.name)
			assert.Equal(t, tt.ExpectResult, total)
		})
	}
}
