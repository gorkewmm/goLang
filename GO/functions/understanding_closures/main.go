package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	transformed := transformNumbers(&numbers, func(number int) int { // ikinci parametre olarak anonymus function verdik
		if number%2 == 0 {
			return number * 2
		}
		return number * 3
	})

	fmt.Println(transformed)

	x := createTransformer(5)
	y := createTransformer(3)
	list1 := transformNumbers(&numbers, x)
	list2 := transformNumbers(&numbers, y)
	fmt.Println(list1, list2)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
