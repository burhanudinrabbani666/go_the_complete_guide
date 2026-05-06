# Type switches

```go
	printSomething(1)
	printSomething(1.5)
	printSomething("Hello")

func printSomething(value any) {
	switch value.(type) {
	case string:
		fmt.Println("String: ", value)

	case float64:
		fmt.Println("Float64: ", value)

	case int:
		fmt.Println("Integer: ", value)

	}
}

```


Next: [Extracting type information](./06-extracting-type-information.md)