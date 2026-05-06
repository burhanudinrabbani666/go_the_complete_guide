# instantiating structs

```go
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	user := User{
		FirstName: userFirstName,
		LastName:  userLastName,
		Birthdate: userBirthdate,
		CreatedAt: time.Now(),
	}

```

Next: [Struct literal null values](./05-struct-literal-null-values.md)