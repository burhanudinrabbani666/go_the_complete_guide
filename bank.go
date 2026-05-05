package main

import (
	"fmt"
	"os"
	"strconv"
)

var ACCOUNT_BALANCE_FILE string = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)

	/*
		1. We need Account balance file
		2. change the string to array of byte
		3. add file primissions
	*/
	os.WriteFile(ACCOUNT_BALANCE_FILE, []byte(balanceText), 0666)
}

func getBelanceFromFile() float64 {
	bytes, _ := os.ReadFile(ACCOUNT_BALANCE_FILE)
	balanceText := string(bytes)
	balance, _ := strconv.ParseFloat(balanceText, 64)

	return balance
}

func main() {
	accountBalance := getBelanceFromFile()
	fmt.Println("Welcome GO Bank")

	// Infinite Loops
	for {

		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposite Money")
		fmt.Println("3. Draw Money")
		fmt.Println("4. Exit")

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
			writeBalanceToFile(accountBalance)

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
			writeBalanceToFile(accountBalance)

		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thank for choosing our bank")
			return
		}
	}

}
