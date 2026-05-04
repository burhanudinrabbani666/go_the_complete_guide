package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - (taxRate / 100))
	ratio := ebt / profit

	// For Clearing terminal
	execute := exec.Command("clear")
	execute.Stdout = os.Stdout
	execute.Run()

	fmt.Printf("EBT: %f \n", ebt)
	fmt.Printf("Profit: %f \n", profit)
	fmt.Printf("EBT: %f \n", ratio)
}
