package slices

import "github.com/thereisnoplanb/compare"

// Computes the average of a sequence of values.
//
// # Parameters
//
//	source []TObject
//
// A slice of values to calculate the average of.
//
// # Returns
//
//	result float64
//
// The average of the sequence of values.
//
//	err error
//
// ErrSourceSequenceIsEmpty - when the source sequence contains no elements.
func Average[TSource ~[]TObject, TObject compare.Real](source TSource) (result float64, err error) {
	if len(source) == 0 {
		return 0, ErrSourceSequenceIsEmpty
	}
	return float64(Sum(source)) / float64(len(source)), nil
}
