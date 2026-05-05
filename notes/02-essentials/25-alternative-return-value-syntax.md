# alternative return value syntax


Explicit return
```go
func calculateFutureValues(invesmentAmount, expectedReturnRate, years float64) (futureValue float64, realFeatureValue float64) {
	futureValue = invesmentAmount * math.Pow((1+expectedReturnRate/100), years)
	realFeatureValue = futureValue / math.Pow(1+INFLATION_RATE/100, years)

	return futureValue, realFeatureValue
}

```

Next: []()