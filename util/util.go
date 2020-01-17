package util

import (
	"fmt"

	"github.com/sysdevguru/checkout-service/common"
)

// CalculateAmount has the main discounts logic
func CalculateAmount(items []string) string {
	var amount float64
	var penAmount, tshirtAmount, mugAmount int
	for _, v := range items {
		switch v {
		case "PEN":
			penAmount++
		case "TSHIRT":
			tshirtAmount++
		case "MUG":
			mugAmount++
		}
	}

	amount += float64(penAmount-(penAmount/2)) * common.Server.PenPrice
	if tshirtAmount >= 3 {
		amount += common.Server.TshirtPrice * float64(tshirtAmount) * 0.75
	} else {
		amount += common.Server.TshirtPrice * float64(tshirtAmount)
	}
	amount += common.Server.MugPrice * float64(mugAmount)
	t := fmt.Sprintf("%.2f", amount)

	return t
}