package sliceplus

// Slice ...
type Slice []interface{}

// Filter ...
func (s Slice) Filter(filter func(interface{}) bool) Slice {
	ret := []interface{}{}
	for _, e := range s {
		if filter(e) {
			ret = append(ret, e)
		}
	}
	return ret
}

// AnyMatch if any element matches
func (s Slice) AnyMatch(predicate func(interface{}) bool) bool {
	for _, e := range s {
		if predicate(e) {
			return true
		}
	}
	return false
}

// AllMatch if all element match
func (s Slice) AllMatch(predicate func(interface{}) bool) bool {
	for _, e := range s {
		if !predicate(e) {
			return false
		}
	}
	return true
}

// Map elements by function
func (s Slice) Map(f func(interface{}) interface{}) Slice {
	ret := Slice{}
	for _, e := range s {
		ret = append(ret, f(e))
	}
	return ret
}

// MapToInt map slice to int slice
func (s Slice) MapToInt(f func(interface{}) int) IntSlice {
	ret := IntSlice{}
	for _, e := range s {
		ret = append(ret, f(e))
	}
	return ret
}

// MapToIntG map slice to int slice (generic function)
func (s Slice) MapToIntG(f func(interface{}) interface{}) IntSlice {
	ret := IntSlice{}
	for _, e := range s {
		ret = append(ret, f(e).(int))
	}
	return ret
}

// Collect ...
func (s Slice) Collect(acc func(interface{})) {
	for _, e := range s {
		acc(e)
	}
}

// const Accumulators = {

// }
// func NoDupes(v interface{}, acc func(interface{})))  {

// }

// func (s Slice) RemoveDuplicates(interface{}) bool {

// }
