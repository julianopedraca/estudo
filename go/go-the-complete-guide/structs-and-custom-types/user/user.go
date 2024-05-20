package user

import (
	"errors"
	"fmt"
	"time"
)

// when exporting a struct the fields inside the struct also must have uppercase first letter
// otherwise it wont be exported (yes, you can export partial struct)
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// struct method
func (u *User) OutputUserDetails() {

	// the "correct" way to access a field of a struct, that's a exception for pointers to struct
	//fmt.Println((*u).firstName, (*u).lastName, (*u).birthDate)

	//synthatic sugar
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

// function with pointer to struct
// func outputUserDetails(u *user) {

// the "correct" way to access a field of a struct, that's a exception for pointers to struct
//fmt.Println((*u).firstName, (*u).lastName, (*u).birthDate)

//synthatic sugar
// fmt.Println(u.firstName, u.lastName, u.birthDate)
// }

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// new is a convecion
// utility function that create a struct
func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name, last name and birthdate are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}, nil
}
