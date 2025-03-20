package main

import (
	"errors"
	"fmt"
	"time"
)

func getUser(text string) string {
	fmt.Println(text)
	var value string
	fmt.Scanln(&value)
	return value
}

func outputUserInfo(x *user) { //yani, fonksiyon bir user nesnesinin adresini alıyor.
	fmt.Println((*x).firstName, (*x).lastName, (*x).birthdate)
}

func clearUserName(x *user) { //yani, fonksiyon bir user nesnesinin adresini alıyor.
	(*x).firstName = ""
	(*x).lastName = ""
}

func addNewUser(userFirstName, userLastName, userBirthDate string) (*user, error) {
	if (userFirstName == "") || (userLastName == "") || (userBirthDate == "") {
		return nil, errors.New("firstname,lastname and birthdate are required")
	}
	return &user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthDate,
	}, nil
}

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func main() {

	firstname := getUser("Please enter your firstname")
	lastname := getUser("Please enter your lastname")
	birthdate := getUser("Please enter your birthdate")
	if firstname == "" || lastname == "" || birthdate == "" {
		fmt.Println("hatalı giriş")
		return
	}

	var denemeUser user

	denemeUser = user{
		firstName: firstname,
		lastName:  lastname,
		birthdate: birthdate,
		createdAt: time.Now(),
	}

	outputUserInfo(&denemeUser)
	clearUserName(&denemeUser)
	outputUserInfo(&denemeUser)
	NewUser, err := addNewUser("tolga", "bal", "2000")
	if err != nil {
		fmt.Println(err)
		return
	}
	outputUserInfo(NewUser)

}
