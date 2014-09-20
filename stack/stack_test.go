package stack

import (
	"testing"
)


func checkStackSize(t *testing.T, s *Stack, size int) bool {
	if n := s.Len(); n != size {
		t.Errorf("s.size() = %d, want %d", n, size)
		return false
	}
	return true
}

func checkIsEmpty(t *testing.T, s *Stack, flg bool) bool {
	if isEmpty := s.IsEmpty(); isEmpty != flg {
		t.Errorf("s.IsEmpty = %t, want %t", isEmpty, flg)
		return false
	}
	return true
}


	

func TestStack(t *testing.T) {
	s := NewStack()
	checkIsEmpty(t, s, true)
	checkStackSize(t, s, 0)
	s.Push("1")
	checkStackSize(t, s, 1)
	checkIsEmpty(t, s, false)
	s.Push("2")
	checkStackSize(t, s, 2)
	checkIsEmpty(t, s, false)
	v, _ := s.Pop().(string)
	if v != "2" {
		t.Errorf("s.Pop = %s, want %s", v, "2")
	}
	checkStackSize(t, s, 1)
	checkIsEmpty(t, s, false)
	
	v, _ = s.Pop().(string)
	if v != "1" {
		t.Errorf("s.Pop = %s, want %s", v, "1")
	}
	checkStackSize(t, s, 0)
	checkIsEmpty(t, s, true)
}
