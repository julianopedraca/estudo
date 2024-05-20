package main

import (
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
func (u user) outputUserDetails() {

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

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user

	// appUser = user{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthDate: userBirthdate,
	// 	createdAt: time.Now(),
	// }

	//shorter notation
	appUser = user{
		userFirstName,
		userLastName,
		userBirthdate,
		time.Now(),
	}

	// ... do something awesome with that gathered data!

	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
