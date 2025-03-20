package main

import "fmt"

func main() {
	var age = 32
	agePointer := &age
	fmt.Println(age)

	getAdultYear(agePointer)
	fmt.Println(age)
}

func getAdultYear(a *int) {
	*a = *a - 18
}
