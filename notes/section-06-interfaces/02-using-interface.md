# Using Interface

```go
type Saver interface {
	Save() error
}


// The data should be have Save method in there own
func saveData(data Saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Failed to sava Data")
		return err
	}

	fmt.Println("Saving the note successed")
	return nil
}

// example: struct

type User struct{
  FirstName string
}

func (user User) Save(){
  // .........
}

// Implement::
newUser := User{"Bani"}
saveData(newUser) // This is correct because User Struct have method Save


```

Next: [Embedded interfaces](./03-embedded-interfaces.md)