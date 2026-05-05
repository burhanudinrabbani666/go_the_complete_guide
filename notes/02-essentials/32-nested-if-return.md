# nested if return


```go
	else if userChoice == 2 {
		fmt.Print("Your Deposit: ")

		var depositAmount float64
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Invlid amount. Must be greater than 0")
			return
		}
  }

```
Next: [for loops](./33-for-loops.md)