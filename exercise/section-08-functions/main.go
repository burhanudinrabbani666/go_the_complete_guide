package main

import "fmt"

func main() {

	numbers := []int{1, 10, 15}
	sum := sumUp(numbers...)
	sum2 := sumUp(2, 4, 6, 6)

	fmt.Println(sum)
	fmt.Println(sum2)

}

// ...int make dynamic
func sumUp(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
