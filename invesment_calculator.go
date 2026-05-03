package main

import (
	"fmt"
	"math"
)

func main() {

	// This all type float.
	invesmentAmount := 1000.0
	expectedReturnRate := 5.5
	years := 10.0

	var future = invesmentAmount * math.Pow((1+expectedReturnRate/100), years)
	fmt.Println(future)

}
