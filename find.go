package godash

import (
	"reflect"
	"errors"
)


func FindIndex(lst interface{}, fnc interface{}) (int, error) {
	lstVal := reflect.ValueOf(lst)
	err := checkIterableExceptMap(lstVal)
	if err != nil {
		return -1, err
	}

	fncVal := reflect.ValueOf(lst)
	if fncVal.Type().NumIn() != 1 {
		return -1, errors.New("callback NOT take one argument")
	}
	if fncVal.Type().NumOut() != 1 {
		return -1, errors.New("callback NOT return one argument")
	}
/*
	// if string
	for i, slcVal := range lstVal {
		if fnc(slcVal) {
			return i
		} 
	}
*/
	return -1, nil
	// if slice
}

/*
func FindIndex(slc []int, fnc func(value int)(bool)) int { 
	for i, slcVal := range slc {
		if fnc(slcVal) {
			return i
		} 
	}
	return -1
}
*/

func FindLastIndex(slc []int, fnc func(value int)(bool)) int {
	lastIndex := -1
	for i, slcVal := range slc {		
		if fnc(slcVal) {
			lastIndex = i
		}
	}
	return lastIndex
}
