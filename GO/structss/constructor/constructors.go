package main

import (
	"example/user"
	"fmt"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!

	var appUser *user.User

	appUser, err := user.NewUser(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

	adminUser := user.NewAdmin("gorkemuysal2002@hotmail.com", "123")
	adminUser.OutputUserDetails() //adminUser.User.OutputUserDetails()  --explicit embadd
	adminUser.ClearUserName()     //adminUser.User.ClearUserName()      --explicit embadd
	adminUser.OutputUserDetails() //adminUser.User.OutputUserDetails()  --explicit embadd

	// (&appUser).outputUserDetails()
	// (&appUser).clearUserName()
	// (&appUser).outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
