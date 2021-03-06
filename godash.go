/*
  This godash package implements a library for functional programming
  in golang.
  This package is inspired by JavaScript library, "Lo-Dash". (http://www.lodash.com)
*/
package godash

import (
	"errors"
	"reflect"
)

//
func checkIterable(value reflect.Value) error {
	kind := value.Kind()
	if kind != reflect.Array && kind != reflect.Slice && kind != reflect.Map && kind != reflect.String {
		return errors.New("arg1 NOT iterable")
	}
	return nil
}

func checkIterableExceptMap(value reflect.Value) error {
	kind := value.Kind()
	if kind != reflect.Array && kind != reflect.Slice && kind != reflect.String {
		return errors.New("NOT iterable")
	}
	return nil
}
		
	 
