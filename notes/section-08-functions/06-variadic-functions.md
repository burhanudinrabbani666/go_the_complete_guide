# Variadic functions

A variadic function in Go (Golang) is a function that can accept an arbitrary number of arguments of a specific type. Common examples in the standard library include fmt.Println and the built-in append function.

```go
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
```

Next: [Splitting slice parameter values](./07-splitting-slices-parameter-values.md)