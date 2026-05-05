# Writing File

```go
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

```

Next: []()