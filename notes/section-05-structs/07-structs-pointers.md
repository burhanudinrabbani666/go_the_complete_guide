# Structs pointers

```go
	outputUserDetail(&user)

  // ...


func outputUserDetail(user *User) {

	// Why in this dont need *?
	// because is shortcuted by Golang
	// so that dont need *user.
	fmt.Println("First Name: ", user.FirstName)
	fmt.Println("Last Name: ", user.LastName)
	fmt.Println("Birthdate: ", user.Birthdate)
	fmt.Println("Created At: ", user.CreatedAt)
}

```

Next: [Introducing methods](./08-introducing-methods.md)