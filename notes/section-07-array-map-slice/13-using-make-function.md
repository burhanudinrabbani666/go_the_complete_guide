# using make function

```go
	userName := make([]any, 10)
	
	for index := range userName {
		userName[index] = index + 1
	}

	userName = append(userName, "Bani")
	userName = append(userName, "Aisa")

	fmt.Println(userName)
```

Next: [making maps](./14-making-maps.md)