package main

import (
	"fmt"
	"math"
)

const INFLATION_RATE = 2.5

func main() {

	expectedReturnRate := 5.5

	var invesmentAmount float64
	var years float64

	// ------------------ Logic Start ---------------------- //

	printQuestion("Invesment Amount: ")
	fmt.Scan(&invesmentAmount)

	printQuestion("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	printQuestion("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(invesmentAmount, expectedReturnRate, years)

	fmt.Printf("Future Value: %.2f\n", futureValue)
	fmt.Printf("Future value adjusat for inflation: %.2f\n", futureRealValue)
}

func printQuestion(question string) {
	fmt.Print(question)
}

func calculateFutureValues(invesmentAmount, expectedReturnRate, years float64) (futureValue float64, realFeatureValue float64) {
	futureValue = invesmentAmount * math.Pow((1+expectedReturnRate/100), years)
	realFeatureValue = futureValue / math.Pow(1+INFLATION_RATE/100, years)

	return futureValue, realFeatureValue
}
