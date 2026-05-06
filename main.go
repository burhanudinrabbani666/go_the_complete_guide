package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(numbers)
	fmt.Println(doubled)
	fmt.Println(tripled)

}

func transformNumbers(numbers *[]int, transform func(int) int) []int {

	dNumber := []int{}
	for _, number := range *numbers {
		dNumber = append(dNumber, transform(number))
	}

	return dNumber
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
