package stack

import "testing"

func Test(t *testing.T) {
	s := New()
	if s.size != 0 {
		t.Errorf("Length of the stack cannot be empty")
	}

	if s.top != nil {
		t.Errorf("Node was not initialized to nil")
	}

	s.Push(1)
	if s.size != 1 {
		t.Errorf("Length of the stack did not increase")
	}

	item := s.Peek()
	if item.(int) != 1 {
		t.Errorf("Returned wrong value. expected=1. got=%v", item)
	}

	s.Push("testing")

	item = s.Peek()
	if s.size != 2 {
		t.Errorf("Size did not increment on second Push()")
	}
	if item.(string) != "testing" {
		t.Errorf("Returned wrong value. expected=testing. got=%v", item)
	}

	item = s.Pop()
	if s.size != 1 {
		t.Errorf("Size did not decrement on pop")
	}
	if item.(string) != "testing" {
		t.Errorf("Pop did not return the correct value")
	}

	item = s.Pop()
	if s.size != 0 {
		t.Errorf("Size did not decrement on pop")
	}
	if item.(int) != 1 {
		t.Errorf("Pop did not return the correct value")
	}

}
