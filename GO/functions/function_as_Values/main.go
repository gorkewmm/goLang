package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)
	fmt.Println(doubled)
	fmt.Println(tripled)

}

func transformNumbers(x *[]int, transform transformFn) []int {
	var newSlice = []int{}
	for _, value := range *x {
		newSlice = append(newSlice, transform(value))
	}
	return newSlice
}

func double(a int) int {
	return a * 2
}

func triple(a int) int {
	return a * 3
}
