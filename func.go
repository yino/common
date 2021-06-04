package common

import (
	"reflect"
)

// Search element whether in array and return bool
func InArray(search interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == search {
				return true
			}
		}
	}

	return false
}

// Search element whether in array and returns index and bool
func ArraySearch(search interface{}, target []interface{})(int,bool){
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == search {
				return i, true
			}
		}
	}
	return -1, false
}

// Search key by value and return key
func MapSearch(search interface{}, target map[interface{}]interface{})(interface{}, bool){
	var result interface{}
	ok := false
	for k,v := range target{
		if v == search {
			result = k
			ok = true
			break
		}
	}
	return result, ok
}
