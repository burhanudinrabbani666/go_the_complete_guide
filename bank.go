package main

import "fmt"

func main() {
	fmt.Println("Welcome GO Bank")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposite Money")
	fmt.Println("3. Draw Money")
	fmt.Println("4. Exit")

	var userChoice int
	fmt.Print("Your choice: ")
	fmt.Scan(&userChoice)

	fmt.Printf("Your Choice: %d\n", userChoice)
}
