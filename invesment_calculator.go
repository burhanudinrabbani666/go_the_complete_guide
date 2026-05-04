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

	fmt.Print("Investment Amount: ")
	fmt.Scan(&invesmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("years: ")
	fmt.Scan(&years)

	var futureValue = invesmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// fmt.Println("Future Value: ", futureValue)

	fmt.Printf("Future Value: %v\n", futureValue)
	fmt.Println("Future Value (adjusted for inflation): ", futureRealValue)

}
