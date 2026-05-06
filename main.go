package main

import "fmt"

func main() {
	result := add(1, 2.7)
	fmt.Println(result)
}

// This T is Generic type placeholder
func add[T int | float64 | string](a, b T) T {
	return a + b
}
