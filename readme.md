# Data structures

This is a collection of data structures. To run tests:

```
go test ./... -v
```

---

## 1. Stacks

```go
var s Stack
var err error
var size = 4

s, err = stacks.InitializeStack()
if err != nil {
    fmt.Fatal("unable to initialize stack")
}

s.Push("Name")
s.Push("Hello")

fmt.Printf("Stack size is : %v", stack.Size())
// Stack size is 2

top, err := stack.Peek()
if err != nil {
    fmt.Fatal("unable to peek stack because : %v", err)
}
fmt.Printf("Top : %v", top)

item, err := stack.Pop() // pops "Hello"
item, err := stack.Pop() // pops "Name"
item, err := stack.Pop() // err is not nil because the stack is now empt
```

---

Cheers! > [_rapando](https://twitter.com/_rapando)