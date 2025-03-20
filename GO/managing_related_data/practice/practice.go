package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	myHobbies := [3]string{"coding", "swimming", "hiking"}
	fmt.Println(myHobbies)

	fmt.Println(myHobbies[0]) //1

	newList := myHobbies[1:3] //1
	fmt.Println(newList)

	slice := myHobbies[0:2] //3

	slice = slice[1:cap(slice)] //4
	fmt.Println(slice)

	dynamicArray := []string{"improving skills", "learning english"} //5
	fmt.Println(dynamicArray)
	dynamicArray[1] = "how to code"
	dynamicArray = append(dynamicArray, "being faster") // 6
	fmt.Println(dynamicArray)

	dynamicProducts := []Product{ //product tipinde verileri olan bir slice oluşturduk.
		{title: "t-shirt",
			id:    "AF91L6G9DJ",
			price: 1599.99},
		{title: "hoodie",
			id:    "P22MLPS415",
			price: 2799.99}} //7
	fmt.Println(dynamicProducts)

	var a Product
	a = Product{
		title: "fas",
		id:    "dasfa",
		price: 214,
	}

	dynamicProducts = append(dynamicProducts, a)
	fmt.Println(dynamicProducts)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
