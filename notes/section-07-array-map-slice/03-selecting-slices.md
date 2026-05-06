# Selecting slices

```go
func main() {

	pricesArray := [4]float64{10, 11, 12, 13}
	featuredPrices := pricesArray[1:] // [11, 12, 13]

	fmt.Println(featuredPrices)

}
```

Next: [More slice selection](./04-more-slice-selection.md)