package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	expectedReturnRate := 5.5

	var invesmentAmount float64
	var years float64

	/*
	 Scan scans text read from standard input, storing successive space-separated values into successive arguments.
	 Newlines count as space. It returns the number of items successfully scanned.
	 If that is less than the number of arguments, err will report why.
	*/

	fmt.Print("Investment Amount: ")
	fmt.Scan(&invesmentAmount) // & symbols is pointer for changing value of var

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("years: ")
	fmt.Scan(&years)

	var futureValue = invesmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

}
