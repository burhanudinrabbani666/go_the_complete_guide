package main

import "fmt"

type Product struct {
	id    string
	Title string
	Price float64
}

/*
func main() {

	pricesArray := [4]float64{10, 11, 12, 13}

	// Slices
	featuredPrices := pricesArray[1:]     // [11, 12, 13]
	highlightPrices := featuredPrices[:1] // Slices can create reference slice

	fmt.Println(pricesArray)
	fmt.Println(featuredPrices)
	fmt.Println(highlightPrices)

	featuredPrices[0] = 199  // This is overwrite the original Array
	fmt.Println(pricesArray) // [10 199 12 13]

	fmt.Println(len(highlightPrices)) // Get Length of Array: Number item in Array
	fmt.Println(cap(highlightPrices)) // Get capicity of Array

}
*/

func main() {
	prices := []float64{10.99, 8.99}
	prices = append(prices, 13.55, 19.55) // Adding new Value

	fmt.Println(prices)

}
