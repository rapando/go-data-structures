# Data structures

This is a collection of data structures. To run tests:

```
go test ./... -v
```

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

## 2. Queues

```go
var q *queues.Queue
var err error

q, err = queues.InitializeQueue(4)
if err != nil {
	fmt.Println(err)
	return
}

q.Enqueue(2)
f, _ := q.Peek()
fmt.Println("Queued :", 2, "Peek : ", f)

q.Enqueue(4)
f, _ = q.Peek()
fmt.Println("Queued :", 4, "Peek : ", f)

q.Enqueue(6)
f, _ = q.Peek()
fmt.Println("Queued :", 6, "Peek : ", f)


i, _ := q.Dequeue()
fmt.Println("Dequeued : ", i)

i, _ = q.Dequeue()
fmt.Println("Dequeued : ", i)

i, _ = q.Dequeue()
fmt.Println("Dequeued : ", i)

i, _ = q.Dequeue()
fmt.Println("Dequeued : ", i)

// Queued : 2 Peek :  2
// Queued : 4 Peek :  2
// Queued : 6 Peek :  2
// Dequeued :  2
// Dequeued :  4
// Dequeued :  6
// Dequeued :  <nil>
```
Cheers! 

> [_rapando](https://twitter.com/_rapando)