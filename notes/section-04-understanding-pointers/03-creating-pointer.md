# Creating pointer

```go
func main() {
	age := 32 // Reguler Variable

	var agePointer *int // <-- this * tell you dealing with pointer
	agePointer = &age   // <-- This & is pointer

	fmt.Println("Age", agePointer)  // Return: Age 0x2891dfed2148 <-- This address to memory
	fmt.Println("Age", *agePointer) // <-- This re-referencing so the return is Age 32

}
```

Next: [Pointers as values](./04-pointers-as-values.md)