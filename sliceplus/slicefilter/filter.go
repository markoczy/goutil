package slicefilter

import (
	"github.com/markoczy/goutil/sliceplus"
)

// IntFilterDuplicates ...
func IntFilterDuplicates(elem int, idx int, slice sliceplus.IntSlice) bool {
	for _, e := range slice {
		if e == elem {
			return false
		}
	}
	return true
}
