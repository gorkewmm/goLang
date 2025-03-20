package main

import "fmt"

func add[T int | float64 | string](a, b T) T {
	return a + b
}

func main() {
	result := add(2, 4)
	fmt.Println(result)
}
