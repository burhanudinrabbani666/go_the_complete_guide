# making maps

```go
func main() {

	coursesRatings := make(map[string]float64, 3)

	coursesRatings["Go"] = 4.7
	coursesRatings["React"] = 4.8
	coursesRatings["Angular"] = 4.8

	fmt.Println(coursesRatings)
}
```

Next: [Working type aliases](./15-working-type-aliases.md)