package main

import "fmt"

func main() {
	numbers := []int{5, 30, 40, 50, 60}
	sum := sumup(1, numbers...) //numbers... turn slice into a list of stand alone parameter values
	fmt.Println(sum)
}

func sumup(startingValue int, numbers ...int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}
