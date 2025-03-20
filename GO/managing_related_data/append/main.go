package main

import "fmt"

func main() {
	prices := []float64{10.99, 8.99} //we are creating a slice

	prices[1] = 9.87
	fmt.Println(prices)
	prices = append(prices, 2.45, 5.87, 4.30)

	fmt.Println(prices)

	discountPrices := []float64{101.99, 80.99, 20.59}
	//prices = append(prices, discountPrices) //append needs seperate elements so you can't use like this
	prices = append(prices, discountPrices...)

	fmt.Println(prices)
}
