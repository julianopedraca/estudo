package main

import "fmt"

func main() {
	age := 32

	var agePointer *int

	agePointer = &age

	//use the asteristic if you want to get the value behind the pointer
	fmt.Println("Age:", *agePointer)

	getAdultYears(agePointer)
	fmt.Println(age)
}

func getAdultYears(age *int) {
	*age = *age - 18
}
