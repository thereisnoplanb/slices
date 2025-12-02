package slices

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
//	accumulate accumulate[TAccumulator, TObject]
//
// An accumulate function to be invoked on each element.
//
// # Returns
//
//	result TAccumulator
//
// The final accumulator value.
func Aggreagate[TSource ~[]TObject, TObject any, TAccumulator any](source TSource, seed TAccumulator, accumulate accumulate[TAccumulator, TObject]) (result TAccumulator) {
	result = seed
	for i := range source {
		result = accumulate(result, source[i])
	}
	return result
}
