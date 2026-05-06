# extracting type information

```go
func printSomething(value any) {
	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer: ", intVal)
	}

	float64Val, ok := value.(float64)
	if ok {
		fmt.Println("Float: ", float64Val)
	}

	stringVal, ok := value.(string)
	if ok {
		fmt.Println("String: ", stringVal)
	}

}
```

Next: [Interfaces dynamic types limitations](./07-interfaces-dynamic-types-limitations.md)