package queue

import "testing"

// TODO: clean this up and utilize Table Driven Tests
func Test(t *testing.T) {
	q := New()
	if q.Length() != 0 {
		t.Errorf("initialized queue size was not 0!")
	}
	if q.fst != nil {
		t.Errorf("first element was not nil. got=%v", q.fst)
	}
	if q.lst != nil {
		t.Errorf("first element was not nil. got=%v", q.lst)
	}

	q.Enqueue(1)
	if q.Length() != 1 {
		t.Errorf("Queue size did not increment on Enqueue")
	}
	if q.fst != nil {
		t.Errorf("Enqueue placed item in wrong order")
	}
	if q.lst == nil {
		t.Errorf("Enqueue did not place item in the last position")
	}
	if q.lst.value.(int) != 1 {
		t.Errorf("Object put into queue was malformed. expected=%v. got=%v", 1, q.lst.value)
	}

	q.Enqueue("testing")
	if q.Length() != 2 {
		t.Errorf("Queue size did not increment on Enqueue")
	}
	if q.fst == nil {
		t.Errorf("Enqueue placed item in wrong position")
	}
	if q.lst == nil {
		t.Errorf("Enqueue erased the contents of lst")
	}
	if q.lst.value.(int) != 1 {
		t.Errorf("First object put into queue was malformed. expected=%v. got=%v", 1, q.lst.value)
	}
	if q.fst.value.(string) != "testing" {
		t.Errorf("Object put into queue was malformed. expected=%v. got=%v", 1, q.fst.value)
	}

	q.Enqueue(2)
	if q.Length() != 3 {
		t.Errorf("Queue size did not increment on Enqueue")
	}
	if q.fst == nil {
		t.Errorf("Enqueue placed item in wrong position")
	}
	if q.lst == nil {
		t.Errorf("Enqueue erased the contents of lst")
	}
	if q.lst.value.(int) != 1 {
		t.Errorf("First object put into queue was malformed. expected=%v. got=%v", 1, q.lst.value)
	}
	if q.fst.value.(int) != 2 {
		t.Errorf("Object put into queue was malformed. expected=%v. got=%v", 1, q.fst.value)
	}

	item := q.Peek()
	if item.(int) != 1 {
		t.Errorf("Did not get expected object. expected=%v. got=%v", 1, item.(int))
	}
	if q.Length() != 3 {
		t.Errorf("Queue size decremented on call to Peek. expected=%d. got=%d", 3, q.size)
	}

	item = q.Dequeue()
	if q.Length() != 2 {
		t.Errorf("Queue size did not decrement on Dequeue")
	}
	if item.(int) != 1 {
		t.Errorf("Did not get expected object. expected=%v. got=%v", 1, item.(int))
	}

	item = q.Dequeue()
	if q.Length() != 1 {
		t.Errorf("Queue size did not decrement on Dequeue")
	}
	if item.(string) != "testing" {
		t.Errorf("Did not get expected object. expected=testing. got=%v", item.(string))
	}

	item = q.Dequeue()
	if q.Length() != 0 {
		t.Errorf("Queue size did not decrement on Dequeue")
	}
	if item.(int) != 2 {
		t.Errorf("Did not get expected object. expected=%v. got=%v", 1, item.(int))
	}

	item = q.Dequeue()
	if q.Length() != 0 {
		t.Errorf("Queue size was not 0 when Dequeue on empty Queue")
	}
	if item != nil {
		t.Errorf("Did not return nil when empty. expected=nil. got=%v", item.(int))
	}

}
