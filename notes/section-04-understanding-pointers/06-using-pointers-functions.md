# using pointers functions

```go
func main() {
	age := 32 // Reguler Variable

	// var agePointer *int // <-- this * tell you dealing with pointer
	// agePointer = &age   // <-- This & is pointer

	adultYear := getAdultYears(&age)
	fmt.Println(adultYear)
	fmt.Println(age)

}

func getAdultYears(age *int) int {

	return *age - 18
}

```

Next: [Pointers data mutation](./07-pointers-data-mutation.md)