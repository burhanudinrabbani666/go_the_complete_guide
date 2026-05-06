# more slice selection

```go
func main() {

	pricesArray := [4]float64{10, 11, 12, 13}
	featuredPrices := pricesArray[1:]     // [11, 12, 13]
	highlightPrices := featuredPrices[:1] // Slices can create reference slice

	fmt.Println(pricesArray)
	fmt.Println(featuredPrices)
	fmt.Println(highlightPrices)

}
```

Next: [Diving deeper slices](./05-diving-deeper-slices.md)