package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	double := createTrasnformer(2)
	triple := createTrasnformer(3)

	transformed := TransformedNumbers(&numbers, func(number int) int {
		return number * 2
	})

	doubled := TransformedNumbers(&numbers, double)
	tripled := TransformedNumbers(&numbers, triple)

	fmt.Println(transformed)
	fmt.Println(doubled)
	fmt.Println(tripled)

}

func TransformedNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTrasnformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
