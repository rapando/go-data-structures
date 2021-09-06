package queues

import "testing"

var queue *Queue
var size = 2

func TestInitializeQueue(t *testing.T) {
	var err error
	queue, err = InitializeQueue(size)
	if err != nil {
		t.Fatalf("queue. initialize failed")
	}
}

func TestQueue(t *testing.T) {
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
		err = queue.Enqueue(testRecord.item)
		if (err != nil) != testRecord.fail {
			t.Fatalf("queue. queue operation failed : [%v], found [%v]",
				testRecord.fail, err != nil)
		}
	}
}
