package stack

// Element is an element of a stack.
type Element struct {
	// The value stored with this element.
	value          interface{}
	
	// The stack to which this element belongs.
	stack *Stack

	next *Element
}

type Stack struct {
	top  *Element
	size int
}

func New() *Stack {
	return new(Stack)
}

func (s *Stack) IsEmpty() bool {
	if s.size > 0 {
		return false
	}
	return true
}

func (s *Stack) Len() int {
	return s.size
}

func (s *Stack) Push(value interface{}) {
	s.top = &Element{value: value, stack:s, next: s.top}
	s.size++
}

func (s *Stack) Pop() interface{} {
	if s.size > 0 {
		value := s.top.value
		s.top = s.top.next
		s.size--
		return value
	}
	return nil
}

func (s *Stack) Peep() interface{} {
	if s.size > 0 {
		return s.top.value
	}
	return nil
}
