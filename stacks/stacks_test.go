package stacks

import (
	"testing"
)

var stack *Stack
var size = 2

func TestInitializeStack(t *testing.T) {
	var err error
	stack, err = InitializeStack(size)
	if err != nil {
		t.Fatalf("stack. initialize failed")
	}

}

func TestPush(t *testing.T) {
	var err error
	var testTables = []struct {
		item interface{}
		fail bool
	}{
		{2, false},
		{3, false},
		{4, true},
	}

	for _, testRecord := range testTables {
		err = stack.Push(testRecord.item)
		if (err != nil) != testRecord.fail {
			t.Fatalf("stack. push failed. expected : [%v], found [%v]",
				testRecord.fail, err != nil)
		}
	}
}

func TestSize(t *testing.T) {
	stackSize := stack.Size()
	if stackSize != size {
		t.Fatalf("stack. size check failed. Expected %v, found %v", size, stackSize)
	}
}

func TestPeek(t *testing.T) {
	expected := 3
	var expectedErr error
	top, err := stack.Peek()
	if err != nil {
		t.Fatalf("stack peek failed. Expected %v, found %v", expectedErr, err)
	}
	if top != expected {
		t.Fatalf("stack peek failed. Expected %v, found %v", expected, top)
	}
}

func TestPop(t *testing.T) {
	var testTables = []struct {
		item interface{}
		fail bool
	}{
		{3, false},
		{2, false},
		{4, true},
	}

	for _, testRecord := range testTables {
		item, err := stack.Pop()
		if (err != nil) != testRecord.fail {
			t.Fatalf("stack pop failed. expected [%v, %v], found [%v, %v]",
				testRecord.item, item, nil, err)
		}
	}
}
