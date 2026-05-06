# Understanding Pointers

- What is "pointers"? Variables that store value addresses instead of values.
- Why "pointers"?
  - Avoid Unnecessary value copies
    - By default, Go cretes a copy when passing values to functions
    - For very large & complex values, this may take up too much memory space unnecessarily
    - With pointers, only one value is stored memory (and the address is passed around) 
  - Directly mutate values
    - Pass a pointer (address) instead of a value to a function
    - The function can then directly the underlying value - no return value is required
    - Can lead to less code

Next: [writing code without pointers](./02-writing-code-without-pointers.md)