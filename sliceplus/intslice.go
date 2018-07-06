package sliceplus

// IntSlice ...
type IntSlice []int

// Filter a Slice
func (s IntSlice) Filter(filter func(elem int, idx int, slice IntSlice) bool) IntSlice {
	ret := IntSlice{}
	for i, e := range s {
		if filter(e, i, ret) {
			ret = append(ret, e)
		}
	}
	return ret
}

// AnyMatch if any element matches
func (s IntSlice) AnyMatch(predicate func(int) bool) bool {
	for _, e := range s {
		if predicate(e) {
			return true
		}
	}
	return false
}

// AllMatch if all element match
func (s IntSlice) AllMatch(predicate func(int) bool) bool {
	for _, e := range s {
		if !predicate(e) {
			return false
		}
	}
	return true
}

// Map elements by function
func (s IntSlice) Map(f func(int) int) IntSlice {
	ret := IntSlice{}
	for _, e := range s {
		ret = append(ret, f(e))
	}
	return ret
}

// Collect ...
func (s IntSlice) Collect(acc func(int)) {
	for _, e := range s {
		acc(e)
	}
}
