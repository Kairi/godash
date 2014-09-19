package godash


/*
  delete function.
  in Golang, delete function not implemented
*/
func Delete(slc []int, index int) []int {
	return append(slc[:(index-1)], slc[index:]...)
}
	
