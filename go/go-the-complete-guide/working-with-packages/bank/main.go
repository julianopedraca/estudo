package main

import (
	"fmt"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = getFloatFromFile(accountBalanceFile)

	presentOptions()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---")
		panic("Can't continue, sorry.")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")

	for {

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		// wantsCheckBalance := choice == 1

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fmt.Printf("You deposited: %.2f and your new balance is: %.2f\n", depositAmount, accountBalance)
			writeFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("Withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("insufficient funds")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Printf("You Withdraw: %.2f and your new balance is: %.2f\n", withdrawAmount, accountBalance)
			writeFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			return
		}
	}
}

func getBalanceFromFile(accountBalanceFile string) {
	panic("unimplemented")
}