package main

import "fmt"

func main() {
	userNames := make([]string, 2)
	userNames[0] = "Julie"
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	for index, value := range userNames {
		fmt.Println(index)
		fmt.Println(value)
	}

	// for i := 0; i < len(userNames); i++ {
	// 	fmt.Println(i)
	// 	fmt.Println(userNames[i])
	// }

}
