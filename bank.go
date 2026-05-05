package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var ACCOUNT_BALANCE_FILE string = "balance.txt"

func writeValueToFile(value float64, fileName string) {
	balanceText := fmt.Sprint(value)
	err := os.WriteFile(fileName, []byte(balanceText), 0666)

	if err != nil {
		panic(err)
	}

}

func getFloatFromFile(fileName string) (float64, error) {
	bytes, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("Failed to Find file.")
	}

	valueText := string(bytes)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("Failed to Parse Storage Value.")
	}

	return value, nil
}

func main() {
	accountBalance, err := getFloatFromFile(ACCOUNT_BALANCE_FILE)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------------------")

		panic(err)
	}

	fmt.Println("Welcome GO Bank")

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
			writeValueToFile(accountBalance, ACCOUNT_BALANCE_FILE)

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
			writeValueToFile(accountBalance, ACCOUNT_BALANCE_FILE)

		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thank for choosing our bank")
			return
		}
	}

}
