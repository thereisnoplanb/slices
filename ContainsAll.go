package slices

import "github.com/thereisnoplanb/generic"

// Determines whether a sequence contains all specified elements by using the default equality comparer.
//
// Parameters
//
//	source []TObject - A sequence in which to locate a values.
//	values TObject - The values to locate in the sequence.
//
// Return Value
//
//	result bool - true if the source sequence contains all elements that have values ​​equal to all of the values ​​in the specified values; false otherwise.
func ContainsAll[TSource ~[]TObject, TObject generic.Equatable](source TSource, values ...TObject) (result bool) {
	for _, item := range values {
		if !Contains(source, item) {
			return false
		}
	}
	return true
}
