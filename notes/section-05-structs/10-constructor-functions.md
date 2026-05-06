# Constructor functions

```go
func NewUser(firstName, lastName, birthdate string) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Birthdate: birthdate,
		CreatedAt: time.Now(),
	}

}
```

```go
	user := NewUser("Burhanudin", "Rabbani", "11/14/2002")
```


Next: [Constructor Validation](./11-constructor-validation.md)