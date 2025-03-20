package main

import "fmt"

func main() {
	prices := []float64{10.99, 8.99} //we are creating a slice

	prices[1] = 9.87
	fmt.Println(prices)

	updatedPrice := append(prices, 45.99) //orjinal dilime ekleme yapmaz, yeni slice döndürür

	fmt.Println(prices, updatedPrice)
}
