# writing code without pointers

```go
func main() {
	age := 32 // Reguler Variable
	fmt.Println("Age:", age)
	fmt.Println(getAdultYears(32))
}

func getAdultYears(age int) int {
	return age - 18
}
```

Next: [Creating pointer](./03-creating-pointer.md)