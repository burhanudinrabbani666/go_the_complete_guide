# functions return values Scope 

```go
func calculateFutureValues(invesmentAmount, expectedReturnRate, years float64) (float64, float64) {
	futureValue := invesmentAmount * math.Pow((1+expectedReturnRate/100), years)
	realFeatureValue := futureValue / math.Pow(1+INFLATION_RATE/100, years)

	return futureValue, realFeatureValue
}
```

```go
	futureValue, futureRealValue := calculateFutureValues(invesmentAmount, expectedReturnRate, years)
```

Next: [Alternative return value syntax](./25-alternative-return-value-syntax.md)