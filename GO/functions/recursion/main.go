package main

import "fmt"

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)

	// result := 1
	// for i := 1; i <= number; i++ {
	// 	result = result * i
	// 	fmt.Println(result)
	// }
	// return result
}

func main() {
	fmt.Println(factorial(4))
}
