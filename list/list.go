package main

import (
	"fmt"
)

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

// Time to practice what you learned!

//  1. Create a new array (!) that contains three hobbies you have
//     Output (print) that array in the command line.
//  2. Also output more data about that array:
//     - The first element (standalone)
//     - The second and third element combined as a new list
//  3. Create a slice based on the first element that contains
//     the first and second elements.
//     Create that slice in two different ways (i.e. create two slices in the end)
//  4. Re-slice the slice from (3) and change it to contain the second
//     and last element of the original array.
//  5. Create a "dynamic array" that contains your course goals (at least 2 goals)
//  6. Set the second goal to a different one AND then add a third goal to that existing dynamic array
//  7. Bonus: Create a "Product" struct with title, id, price and create a
//     dynamic list of products (at least 2 products).
//     Then add a third product to the existing list of products.
type ProductPractice struct {
	Id    string
	Title string
	Price float64
}

func main() {
	// 1.
	hobbies := []string{"Coding", "Gaming", "Sleeping"}
	fmt.Println(hobbies)

	// 2.
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	// 3.
	hobbiesSlice := hobbies[:2]

	// 4.
	hobbiesSlice = []string{hobbies[1], hobbies[2]}
	fmt.Println(hobbiesSlice)

	// 5.
	goals := []string{"Backend Developer", "Professional Software Enginering"}

	// 6.
	goals[1] = "Being Stoik"
	goals = append(goals, "Becoming Rich")

	// 7.
	products := []ProductPractice{
		{
			Id:    "1",
			Title: "Shampo",
			Price: 13.00,
		},
		{
			Id:    "2",
			Title: "Brush",
			Price: 7.00,
		},
	}

	products = append(products, ProductPractice{
		Id:    "3",
		Title: "Conditioner",
		Price: 10.00,
	})

	fmt.Println(products)

	newHobbies := []string{"Drinking", "Watch Anime"}
	hobbies = append(hobbies, newHobbies...)

}
