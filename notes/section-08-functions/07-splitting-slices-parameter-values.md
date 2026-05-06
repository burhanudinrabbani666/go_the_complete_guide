# splitting slices parameter values

```go
func main() {

	numbers := []int{1, 10, 15}
	sum := sumUp(numbers...) // <-- This is Spliting Parameter Values
	sum2 := sumUp(2, 4, 6, 6)
	

	fmt.Println(sum)
	fmt.Println(sum2)

}
```