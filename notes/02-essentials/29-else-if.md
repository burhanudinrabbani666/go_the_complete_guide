# else if


```go
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
	} else {
		fmt.Println("Goodbye!")
	}

```

Next: [if statements exercise](./30-if-statements-exercise.md)