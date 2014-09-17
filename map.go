package godash

import (
	"reflect"
	"errors"
)


// exported function
func Map(lst interface{}, fnc interface{}) (interface{}, error) {
	lstVal := reflect.ValueOf(lst)
	err := checkIterable(lstVal)
	if err != nil {
		return nil, err
	}
	fncVal := reflect.ValueOf(fnc)
	if lstVal.Kind() == reflect.Map {
		if fncVal.Type().NumIn() != 2 {
			return nil, errors.New("fnc NOT one argument")
		}
		if fncVal.Type().NumOut() != 2 {
			return nil, errors.New("fnc NOT return exactly one value")
		}
	} else {
		if fncVal.Type().NumIn() != 1 {
			return nil, errors.New("fnc must take exactly one argument")
		}
		if fncVal.Type().NumOut() != 1 {
			return nil, errors.New("fnc must return exactly one value")
		}
	}
	
	if lstVal.Kind() == reflect.String {
		buf := make([]rune, lstVal.Len())
		f, ok := fnc.(func(rune) rune)
		if !ok {
			return nil, errors.New("invalid filter function")
		}
		for i, c := range lst.(string) {
			buf[i] = f(c)
		}
		return string(buf), nil
	} else if lstVal.Kind() == reflect.Map {
		res := reflect.MakeMap(lstVal.Type())
		for _, k := range lstVal.MapKeys() {
			v := lstVal.MapIndex(k)
			out := fncVal.Call([]reflect.Value{k, v})
			res.SetMapIndex(out[0], out[1])
		}
		return res.Interface(), nil
	} else {
		res := reflect.MakeSlice(reflect.SliceOf(fncVal.Type().Out(0)), lstVal.Len(), lstVal.Cap())
		for i, l := 0, lstVal.Len(); i < l; i++ {
			v := lstVal.Index(i)
			res.Index(i).Set(fncVal.Call([]reflect.Value{v})[0])
		}
		return res.Interface(), nil
	}
}

func MapSlice() {
}

func MapString() {
}

func MapArray() {
}

//FIX:NAME
func MapMap() {
}
