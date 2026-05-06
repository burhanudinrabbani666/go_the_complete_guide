# struct embedding

```go
type User struct {
	FirstName string
	LastName  string
	Birthdate string
	CreatedAt time.Time
}

type Admin struct {
	Email    string
	Password string
	User
}
```


Next: [Summary](./15-structs-summary.md)