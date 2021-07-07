package common

import (
	"reflect"
)

// InArray: Search element whether in array and return bool
// 判断目标元素是否出现在数组中
func InArray(search interface{}, arr interface{}) bool {
	arrValue := getArrValue(arr)
	for i := 0; i < arrValue.Len(); i++ {
		if arrValue.Index(i).Interface() == search {
			return true
		}
	}

	return false
}

// ArraySearch: Search element whether in array and returns index and bool
// 搜索数组中目标元素的位置 并返回下标
func ArraySearch(search interface{}, arr interface{}) (int, bool) {

	arrValue := getArrValue(arr)
	for i := 0; i < arrValue.Len(); i++ {
		if arrValue.Index(i).Interface() == search {
			return i, true
		}
	}
	return -1, false
}

// ArrayCountValues: The array () function counts all the values in the array.
// 统计数组元素出现的次数
func ArrayCountValues(arr interface{}) (map[interface{}]int, bool) {
	result := make(map[interface{}]int)
	arrValue := getArrValue(arr)
	if arrValue.Len() == 0 {
		return result, false
	}
	for i := 0; i < arrValue.Len(); i++ {
		result[arrValue.Index(i).Interface()]++
	}

	return result, true
}

func getArrValue(arr interface{}) reflect.Value {
	objType := reflect.TypeOf(arr).Kind()
	if objType != reflect.Array && objType != reflect.Slice {
		panic("Incorrect parameter type：arr interface unequal to slice or array")
	}

	return reflect.ValueOf(arr)
}
