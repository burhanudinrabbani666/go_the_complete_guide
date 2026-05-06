# Working type aliases

```go
func (floatMap FloatMap) output() {
	fmt.Println(floatMap)
}

func main() {

	coursesRatings := make(FloatMap, 3)

	coursesRatings["Go"] = 4.7
	coursesRatings["React"] = 4.8
	coursesRatings["Angular"] = 4.8

	coursesRatings.output()
}
```

Next: [For loops arrays slices maps](./16-for-loops-arrays-slices-maps.md)