# Constructor validation

```go
func NewUser(firstName, lastName, birthdate string) (*User, error) {

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First Name, Last Name, Birthdate is required")
	}

	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Birthdate: birthdate,
		CreatedAt: time.Now(),
	}, nil

}
```

Next: [Structs packages exports](./12-structs-packages-exports.md)