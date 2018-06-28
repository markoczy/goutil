package sliceplus

import "reflect"

// OfInts Create Slice of Int
func OfInts(data []int) IntSlice {
	return IntSlice(data)
}

// Of Create Slice of anything (less performance)
func Of(data interface{}) Slice {
	v := reflect.ValueOf(data)
	arr := make([]interface{}, v.Len())
	for i := 0; i < v.Len(); i++ {
		arr[i] = v.Index(i).Interface()
	}
	return Slice(arr)
}
