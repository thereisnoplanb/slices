package slices

import "github.com/thereisnoplanb/delegate"

// Applies an accumulate function over a sequence. The specified seed value is used as the initial accumulator value.
//
// # Parameters
//
//	source []TObject
//
// A slice to aggregate over.
//
//	seed TAccumulator
//
// The initial accumulator value.
//
//	accumulate Accumulator[TObject, TAccumulator]
//
// An accumulate function to be invoked on each element.
//
// # Returns
//
//	result TAccumulator
//
// The final accumulator value.
func Aggreagate[TSource ~[]TObject, TObject any, TAccumulator any](source TSource, seed TAccumulator, accumulate delegate.Accumulator[TObject, TAccumulator]) (result TAccumulator) {
	result = seed
	for _, item := range source {
		result = accumulate(result, item)
	}
	return result
}
