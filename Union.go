package slices

import "github.com/thereisnoplanb/compare"

// Produces the set union of two sequences by using the default equality comparer.
//
// # Parameters
//
// 	first []TObject
//
// A sequence whose distinct elements form the first set for the union.
//
// 	second []TObject
//
// A sequence whose distinct elements form the second set for the union.
//
// # Returns
//
//	result []TObject
//
// A sequence that contains the elements from both input sequences, excluding duplicates.
func Union[TSource ~[]TObject, TObject comparable](first, second TSource) (result TSource) {
	result = append(result, first...)
	result = append(result, second...)
	return Distinct(result)
}

// Produces the set union of two sequences by using the default equality comparer.
//
// # Parameters
//
// 	first []TObject
//
// A sequence whose distinct elements form the first set for the union.
//
// 	second []TObject
//
// A sequence whose distinct elements form the second set for the union.
//
// # Returns
//
//	result []TObject
//
// A sequence that contains the elements from both input sequences, excluding duplicates.
func UnionEquatable[TSource ~[]TObject, TObject compare.IEquatable[TObject]](first, second TSource) (result TSource) {
	result = append(result, first...)
	result = append(result, second...)
	return DistinctEquatable(result)
}

// Produces the set union of two sequences by using a specified equality comparer.
//
// # Parameters
//
// 	first []TObject
//
// A sequence whose distinct elements form the first set for the union.
//
// 	second []TObject
//
// A sequence whose distinct elements form the second set for the union.
//
//	equality Equality[TObject]
//
// The Equality[TObject] function to compare values for equality.
//
// # Returns
//
//	result []TObject
//
// A sequence that contains the elements from both input sequences, excluding duplicates.
func UnionFunc[TSource ~[]TObject, TObject any](first, second TSource, equality compare.Equality[TObject]) (result TSource) {
	result = append(result, first...)
	result = append(result, second...)
	return DistinctFunc(result, equality)
}
