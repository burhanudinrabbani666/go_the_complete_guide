package main

import "fmt"

type Product struct {
	id    string
	Title string
	Price float64
}

func main() {

	productNames := [4]string{"A Book"}
	fmt.Println(productNames[3])

	prices := [4]float64{10, 11, 12, 13} // Length is declare first
	fmt.Println(prices[1])               // output: [11]

	// Set value in array
	productNames[3] = "Meditations"
	fmt.Println(productNames)

}
