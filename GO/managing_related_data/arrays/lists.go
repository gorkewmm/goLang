package main

import (
	"fmt"
)

func main() {
	var productNames = [4]string{"A Book"}
	price := [4]float64{10.99, 9.99, 45.99, 20.0}
	price[1] = 1.99
	fmt.Println(price[1])
	fmt.Println(price)
	fmt.Println(productNames)

	featuredPrices := price[1:3]
	highlightedPrices := featuredPrices[1:]

	fmt.Println(featuredPrices)
	fmt.Println(len(featuredPrices), cap(featuredPrices))
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
}
