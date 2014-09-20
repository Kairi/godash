package queue

type Element struct {
	value          interface{}
	next, previous *Element
}

type Queue struct {
	top  *Element
	tail *Element
	size int
}

func NewQueue() *Queue {
	return new(Queue)
}

func (q *Queue) IsEmpty() bool {
	if q.size > 0 {
		return false
	}
	return true
}

func (q *Queue) Len() int {
	return q.size
}

// alias to Push()
func (q *Queue) Enqueue(value interface{}) {
	q.Push(value)
}

func (q *Queue) Push(value interface{}) {
	q.size++
	if q.size == 1 { // when first Element
		q.top = &Element{value: value, next: nil, previous: nil}
		q.tail = q.top
		return
	}
	q.tail.previous = &Element{value: value, next: nil, previous: nil}
	q.tail = q.tail.previous

}

// alias to Pop()
func (q *Queue) Dequeue() interface{} {
	return q.Pop()
}

func (q *Queue) Pop() interface{} {
	if !q.IsEmpty() {
		value := q.top.value
		q.top = q.top.previous
		q.size--
		return value
	}
	return nil
}

func (q *Queue) Peep() interface{} {
	if !q.IsEmpty() {
		return q.top
	}
	return nil
}
