package main

import (
	"fmt"
	filesops "go_the_complete_guide/bank/files_ops"

	"github.com/brianvoe/gofakeit"
)

var ACCOUNT_BALANCE_FILE string = "balance.txt"

func main() {
	accountBalance, err := filesops.GetFloatFromFile(ACCOUNT_BALANCE_FILE)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------------------")

		panic(err)
	}

	fmt.Println("Welcome GO Bank")
	fmt.Printf("Reach us 24/7 at %s\n", gofakeit.Phone())

	// Infinite Loops
	for {

		PresentOptions()

		var userChoice int
		fmt.Print("Your choice: ")
		fmt.Scan(&userChoice)

		switch userChoice {

		case 1:
			fmt.Printf("Your Balance: %f\n", accountBalance)

		case 2:
			fmt.Print("Your Deposit: ")

			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invlid amount. Must be greater than 0")
				continue
			}

			accountBalance += depositAmount
			fmt.Printf("Balance Update! new Amount: %.2f\n", accountBalance)
			filesops.WriteValueToFile(accountBalance, ACCOUNT_BALANCE_FILE)

		case 3:
			fmt.Print("Withdrawal Amount: ")

			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invlid amount. Must be greater than 0")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Printf("Invlid amount. Must be less than Your account balance: %.2f\n", accountBalance)
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Printf("Balance Update! new Amount: %.2f\n", accountBalance)
			filesops.WriteValueToFile(accountBalance, ACCOUNT_BALANCE_FILE)

		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thank for choosing our bank")
			return
		}
	}

}
