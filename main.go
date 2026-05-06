package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	transformed := TransformedNumbers(&numbers, func(number int) int {
		return number * 2
	})

	fmt.Println(transformed)

}

func TransformedNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, val*2)
	}

	return dNumbers
}
