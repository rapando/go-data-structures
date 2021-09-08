package linkedlists

import "testing"

var linkedList *LinkedList

func TestInitializeLinkedList(t *testing.T) {
	data := 2
	linkedList = InitializeLinkedList(data)
	tail := linkedList.GetTail().data
	if data != tail {
		t.Fatalf("linked list. initialize failed, expected : %v, found : %v",
			data, tail)
	}
}

func TestAddNode(t *testing.T) {
	linkedList.AddNode(4)
	linkedList.AddNode(6)
	linkedList.AddNode(8)
	tail := linkedList.GetTail().data
	if tail != 8 {
		t.Fatalf("linked list. add node failed. expected %v, found %v", 8, tail)
	}
}

func TestTraverse(t *testing.T) {
	expected := "2 4 6 8"
	if found := linkedList.Traverse(); found != expected {
		t.Fatalf("linked list. traverse failed. expected [%v], found [%v]", expected, found)
	}
}

func TestSearch(t *testing.T) {
	var testTable = []struct {
		item     int
		found    bool
		position int
	}{
		{2, true, 0},
		{4, true, 1},
		{6, true, 2},
		{8, true, 3},
		{10, false, 0},
	}

	for _, testRecord := range testTable {
		found, position := linkedList.Search(testRecord.item)
		if found != testRecord.found || position != testRecord.position {
			t.Fatalf("linked list. search failed. Expected [%v,%v] Found [%v,%v]",
				testRecord.found, testRecord.position, found, position)
		}
	}
}
