# Embedded Interfaces

```go
type Saver interface {
	Save() error
}

type Outputable interface {
	Saver
	Display()
}
```

Next: [Any value allowed type](./04-any-value-allowed-type.md)