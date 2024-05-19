package main

import "fmt"

func main() {
	age := 32

	var agePointer *int

	agePointer = &age

	//use the asteristic if you want to get the value behind the pointer
	fmt.Println("Age:", *agePointer)

	adultYears := getAdultYears(agePointer)
	fmt.Println(adultYears)
}

func getAdultYears(age *int) int {
	return *age - 18
}
