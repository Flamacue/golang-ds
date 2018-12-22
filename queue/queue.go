package queue

type Queue struct {
	first *node
	last  *node
	size  int
}

type node struct {
	value interface{}
	nxt   *node
}

func New() *Queue {
	return &Queue{nil, nil, 0}
}

func (q *Queue) Enqueue(val interface{}) {
	if q.last == nil {
		q.last = &node{val, nil}
	} else if q.first == nil {
		q.first = &node{val, nil}
		q.last.nxt = q.first
	} else {
		ins := &node{val, nil}
		q.first.nxt = ins
		q.first = ins
	}
	q.size++
}

func (q *Queue) Dequeue() interface{} {
	if q.size == 0 {
		return nil
	}
	ret := q.last.value
	q.last = q.last.nxt
	q.size--
	return ret
}

func (q *Queue) Length() int {
	return q.size
}

func (q *Queue) Peek() interface{} {
	if q.size == 0 {
		return nil
	}
	return q.last.value
}
