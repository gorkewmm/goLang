package main

import "fmt"

func main() {
	// numbers := []int{1, 10, 15}

	numbers := sumup(1, 10, 15, 40, -2)
	fmt.Println(numbers)
}
func sumup(startingValue int, numbers ...int) int {
	// sum := 0
	// for _, v := range numbers {
	// 	sum = sum + v
	// }
	// return sum
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}
