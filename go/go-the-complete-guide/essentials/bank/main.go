package main

import "fmt"

func main() {
	var accountBalance = 1000.0

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposity money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	// wantsCheckBalance := choice == 1

	if choice == 1 {
		fmt.Println("Your balance is", accountBalance)
	} else if choice == 2 {
		fmt.Print("Your deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		accountBalance += depositAmount
		fmt.Printf("You deposited: %.2f and your new balance is: %.2f\n", depositAmount, accountBalance)
	} else if choice == 3 {
		fmt.Print("Withdraw: ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)

		accountBalance -= withdrawAmount
		fmt.Printf("You Withdraw: %.2f and your new balance is: %.2f\n", withdrawAmount, accountBalance)
	} else {
		fmt.Println("Goodbye!")
	}
}
