package main

import (
	"fmt"
	"math"
)

func main() {
	inflationRate := 2.5

	// This all type float.
	invesmentAmount := 1000.0
	expectedReturnRate := 5.5
	years := 10.0

	var futureValue = invesmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(invesmentAmount)
	fmt.Println(futureRealValue)

}
