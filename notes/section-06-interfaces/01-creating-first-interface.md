# Creating first interface

An interface type is defined as a set of method signatures. A value of interface type can hold any value that implements those methods.


```go
type saver interface {
	Save() error
}
```

Next: [Using interface](./02-using-interface.md)