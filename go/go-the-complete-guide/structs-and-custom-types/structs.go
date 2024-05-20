package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// struct method
func (u *user) outputUserDetails() {

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

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// new is a convecion
// utility function that create a struct
func newUser(firstName, lastName, birthdate string) (*user, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name, last name and birthdate are required")
	}

	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user

	appUser, err := newUser(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	// appUser = user{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthDate: userBirthdate,
	// 	createdAt: time.Now(),
	// }

	//shorter notation
	// appUser = user{
	// 	userFirstName,
	// 	userLastName,
	// 	userBirthdate,
	// 	time.Now(),
	// }

	// ... do something awesome with that gathered data!

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
