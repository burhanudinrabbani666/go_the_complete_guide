package main

import "fmt"

type TransformFn func(int) int

func main() {

	numbers := []int{1, 2, 3, 4}
	moreNumber := []int{5, 1, 4}

	trnsforemer := getTrasformer(&numbers)
	trnsforemer2 := getTrasformer(&moreNumber)

	doubled := transformNumbers(&numbers, trnsforemer)
	tripled := transformNumbers(&moreNumber, trnsforemer2)

	fmt.Println(numbers)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform TransformFn) []int {

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

func getTrasformer(numbers *[]int) TransformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}

}
