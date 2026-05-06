# Introducing Methods

```go
type User struct {
	FirstName string
	LastName  string
	Birthdate string
	CreatedAt time.Time
}

func (user User) outputUserDetail() {

	fmt.Printf("First Name: %s\n", user.FirstName)
	fmt.Printf("Last Name: %s\n", user.LastName)
	fmt.Printf("Birthdate: %s\n", user.Birthdate)
	fmt.Printf("Created At: %s\n", user.CreatedAt)
}


// .......

	user.outputUserDetail()

```

Next: [Mutation methods](./09-mutation-methods.md)