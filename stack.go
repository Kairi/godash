package godash

type Stack struct {
	emptyFlg bool
	top *Element
	size int
}

type Element struct {
	value interface{}
	next *Element
}
	


func (s *Stack) isEmpty() bool {
	return s.emptyFlg
}

func (s *Stack) Len() int {
	return s.size
}

func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
	s.emptyFlg = false
}

func (s *Stack) Pop() interface{} {
	if !s.isEmpty() {
		value := s.top.value
		s.top = s.top.next
		s.size--
		return value
	}
	return nil
}

func (s *Stack) Peep() interface{} {
	if !s.isEmpty() {
		return s.top.value
	}
	return nil
}



