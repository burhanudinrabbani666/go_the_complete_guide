# Understanding Recursion

```go
func main() {

	fact := factorial(3)
	fmt.Println(fact)
}

func factorial(numbers int) int {
	if numbers == 0 {
		return 1
	}

	return numbers * factorial(numbers-1)
}

// Recursion is function call himself
```

Next: [Varadic Function](./06-variadic-functions.md)