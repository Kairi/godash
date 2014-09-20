package godash

// this struct is used in stack.go and queue.go
type Element struct {
	value    interface{}
	next     *Element
	previous *Element
}
