package queue

import (
	"testing"
)


func checkQueueSize(t *testing.T, q *Queue, size int) bool {
	if n := q.Len(); n != size {
		t.Errorf("q.Len() = %d, want %d", n, size)
		return false
	}
	return true
}

func checkIsEmpty(t *testing.T, q *Queue, flg bool) bool {
	if isEmpty := q.IsEmpty(); isEmpty != flg {
		t.Errorf("q.IsEmpty = %t, want %t", isEmpty, flg)
		return false
	}
	return true
}


	

func TestStack(t *testing.T) {
	q := NewQueue()
	checkIsEmpty(t, q, true)
	checkQueueSize(t, q, 0)
	q.Push("1")
	checkQueueSize(t, q, 1)
	checkIsEmpty(t, q, false)
	q.Push("2")
	checkQueueSize(t, q, 2)
	checkIsEmpty(t, q, false)
	v, _ := q.Pop().(string)
	if v != "1" {
		t.Errorf("q.Pop = %s, want %s", v, "1")
	}
	checkQueueSize(t, q, 1)
	checkIsEmpty(t, q, false)
	
	v, _ = q.Pop().(string)
	if v != "2" {
		t.Errorf("q.Pop = %s, want %s", v, "2")
	}
	checkQueueSize(t, q, 0)
	checkIsEmpty(t, q, true)
}
