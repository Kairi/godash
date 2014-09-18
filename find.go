package godash

func FindIndex(slc []int, fnc func(value int)(bool)) int { 
	for i, slcVal := range slc {
		if fnc(slcVal) {
			return i
		} 
	}
	return -1
}
