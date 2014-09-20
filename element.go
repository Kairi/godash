package godash

type Element struct {
	value    interface{}
	next     *Element
	previous *Element
}
