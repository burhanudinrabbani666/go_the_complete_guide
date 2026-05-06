# passing struct arguments

```go
func outputUserDetail(user User) {

	fmt.Println("First Name: ", user.FirstName)
	fmt.Println("Last Name: ", user.LastName)
	fmt.Println("Birthdate: ", user.Birthdate)
	fmt.Println("Created At: ", user.CreatedAt)
}
```

Next: [Structs pointers](./07-structs-pointers.md)