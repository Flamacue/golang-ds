package stack

type Stack struct {
	top  *node
	size int
}

type node struct {
	value interface{}
	prev  *node
}

func New() *Stack {
	return &Stack{nil, 0}
}

func (s *Stack) Peek() interface{} {
	if s.size == 0 {
		return nil
	}
	return s.top.value
}

func (s *Stack) Pop() interface{} {
	if s.size == 0 {
		return nil
	}
	s.size--
	ret := s.top
	s.top = s.top.prev
	return ret.value
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Push(val interface{}) {
	n := &node{val, s.top}
	s.top = n
	s.size++
}
