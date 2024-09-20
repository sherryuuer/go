package queue

import "testing"

func TestAddQueue(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		if len(q.items) != i {
			t.Errorf("Incorrect element count: %v, want %v", len(q.items), i)
		}
		if !q.Append(i) {
			t.Errorf("Failed to append item %v", i)
		}
	}
	if q.Append(4) {
		t.Errorf("Should not be able to add item to a full queue")
	}
}

func TestNext(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		q.Append(i)
	}
	for i := 0; i < 3; i++ {
		item, ok := q.Next()
		if !ok {
			t.Errorf("Should be able to get item")
		}
		if item != i {
			t.Errorf("Get item in wrong order: %v, want %v", item, i)
		}
	}
	// Queue is empty at this point
	item, ok := q.Next()
	if ok {
		t.Errorf("Should not be able to get item, get %v", item)
	}
}
