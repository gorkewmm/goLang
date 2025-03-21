package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)
	fmt.Println(doubled)
	fmt.Println(tripled)

	moreNumbers := []int{5, 1, 3, 4}

	transform1 := getTransformetFunction(&numbers)
	transform2 := getTransformetFunction(&moreNumbers)

	a := transformNumbers(&numbers, transform1)
	b := transformNumbers(&moreNumbers, transform2)

	fmt.Println(a, b)
	fmt.Println(numbers, moreNumbers)
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
func getTransformetFunction(x *[]int) transformFn {
	if (*x)[0] == 1 {
		return double
	}
	return triple
}
