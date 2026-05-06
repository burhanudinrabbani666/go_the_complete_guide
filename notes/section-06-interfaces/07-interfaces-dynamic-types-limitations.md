# interfaces dynamic types limitations

```go

// This is Ugly 
func add(a,b any) int {
  aInt, aIsInt := a.(int)
  bInt, bIsInt := b.(int)

  if aIsInt && bIsInt{
    return a + b
  }
}

```


Next: [Introducing generics](./08-introducing-generics.md)