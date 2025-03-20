package main

import "fmt"

func main() {
	var age = 32

	agePointer := &age // & ile pointer oluşturdun

	fmt.Println("The pointer of age is:", agePointer)
	fmt.Println(*agePointer)

	adultYears := getAdultYear(agePointer)
	fmt.Println(adultYears)

	fmt.Println(age)
}

func getAdultYear(a *int) int {
	return *a - 18
}
