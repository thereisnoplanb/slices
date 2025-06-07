package slices

import "github.com/thereisnoplanb/compare"

// Computes the sum of the values in a sequence.
//
// # Parameters
//
//	source []TObject
//
// A slice of values to calculate the sum of.
//
// # Returns
//
//	result TObject
//
// The sum of the values in the sequence.
func Sum[TSource ~[]TObject, TObject compare.Number | compare.String](source TSource) (result TObject) {
	for _, item := range source {
		result += item
	}
	return result
}
