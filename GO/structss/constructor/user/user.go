package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time //provided by time package
}

type Admin struct { //BAŞKA BİR STRUCTU İÇİNE ALAN STRUCT
	email    string
	password string
	User     //Anonymus embadding (anonim gömme)   //	User     User      Explicit embadding (açık gömme)
}

func (x *User) OutputUserDetails() {
	fmt.Println(x.firstName, x.lastName, x.birthdate)
}

func (x *User) ClearUserName() {
	x.firstName = ""
	x.lastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthdate: "---",
			createdAt: time.Now(),
		},
	}
}

func NewUser(firstName, lastName, birthDate string) (*User, error) { //constructor
	if (firstName == "") || (lastName == "") || (birthDate == "") {
		return nil, errors.New("firstname, lastname and birthdate are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthDate,
		createdAt: time.Now(),
	}, nil
}
