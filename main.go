package main

import (
	"fmt"
	"strings"
)

var productPrices = map[string]float64{
	"TSHIRT": 20.00,
	"MUG":    12.50,
	"HAT":    18.00,
	"BOOK":   25.99,
}

func getPrice(itemCode string) (float64, bool) {
	price, found := productPrices[itemCode]
	if !found {
		if strings.HasSuffix(itemCode, "_SALE") {
			slicedItemCode := strings.TrimSuffix(itemCode, "_SALE")
			price, found = productPrices[slicedItemCode]
			if found {
				return price * 0.9, true
			}
		}
		return 0, false
	}
	return price, true
}

func calculateTotal(items []string) float64 {
	var total float64
	for _, item := range items {
		price, found := getPrice(item)
		if !found {
			fmt.Printf("- Item: %s => Product Not Found!\n", item)
		}
		total += price
	}
	return total
}

func main() {
	items := []string{"TSHIRT", "MUG_SALE", "HAT", "BOOK"}

	fmt.Println("-------------- Simple Sales Order Processor------------")
	fmt.Printf("Total price: %.2f\n", calculateTotal(items))
}
