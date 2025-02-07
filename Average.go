package slices

import "github.com/thereisnoplanb/generic"

// Computes the average of a sequence of values.
//
// Parameters
//
//	source []TObject - A slice of values to calculate the average of.
//
// Return Value
//
//	result float64 - The average of the sequence of values.
//	err error - ErrSourceSequenceIsEmpty if source contains no elements.
func Average[TSource ~[]TObject, TObject generic.Real](source TSource) (result float64, err error) {
	if len(source) == 0 {
		return 0, ErrSourceSequenceIsEmpty
	}
	return float64(Sum(source)) / float64(len(source)), nil
}

// Computes the average of a sequence of values that are obtained by invoking a transform function on each element of the input sequence.
//
// Parameters
//
//	source []TObject - A sequence of values to calculate the average of.
//	valueSelector ValueSelector[TObject, float64] - A transform function to apply to each element.
//
// Return Value
//
//	result float64 - The average of the sequence of values.
//	err error - ErrSourceSequenceIsEmpty if source contains no elements.
func AverageBy[TSource ~[]TObject, TObject any](source TSource, valueSelector generic.ValueSelector[TObject, float64]) (result float64, err error) {
	if len(source) == 0 {
		return 0, ErrSourceSequenceIsEmpty
	}
	return SumBy(source, valueSelector) / float64(len(source)), nil
}
