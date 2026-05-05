package main

import "fmt"

func main() {
	age := 32 // Reguler Variable

	var agePointer *int
	agePointer = &age
	fmt.Println("Age: ", *agePointer) // Output: 32

	calculateAge(agePointer)
	fmt.Println(age) // ouput: 14
}

func calculateAge(age *int) {
	*age -= 18
}
