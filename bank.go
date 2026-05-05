package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome GO Bank")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposite Money")
	fmt.Println("3. Draw Money")
	fmt.Println("4. Exit")

	var userChoice int
	fmt.Print("Your choice: ")
	fmt.Scan(&userChoice)

	/*
		= is for asign to variable, example: var name = bani
		== is for check condition
	*/
	if userChoice == 1 {
		fmt.Printf("Your Balance: %f\n", accountBalance)
	} else if userChoice == 2 {
		fmt.Print("Your Deposit: ")

		var depositAmount float64
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Invlid amount. Must be greater than 0")
			return
		}

		accountBalance += depositAmount
		fmt.Printf("Balance Update! new Amount: %.2f\n", accountBalance)
	} else if userChoice == 3 {
		fmt.Print("Withdrawal Amount: ")

		var withdrawalAmount float64
		fmt.Scan(&withdrawalAmount)

		if withdrawalAmount <= 0 {
			fmt.Println("Invlid amount. Must be greater than 0")
			return
		}

		if withdrawalAmount > accountBalance {
			fmt.Printf("Invlid amount. Must be less than Your account balance: %.2f\n", accountBalance)
			return

		}

		accountBalance -= withdrawalAmount
		fmt.Printf("Balance Update! new Amount: %.2f\n", accountBalance)

	} else {
		fmt.Println("Goodbye!")
	}

}
