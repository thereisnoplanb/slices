package slices

import "github.com/thereisnoplanb/generic"

func Intersect[TSource ~[]TObject, TObject generic.Equatable](collection1, collection2 TSource) (result TSource) {
	result = make(TSource, 0, len(collection1))
	for _, item := range collection1 {
		if Contains(collection2, item) {
			result = append(result, item)
		}
	}
	return result
}

// func IntersectBy[TSource ~[]TObject, T generic.Equatable](collection1, collection2 TSource, equalityComparer EqualityComparer[T]) (result TSource) {
// 	result = make(TSource, 0, len(collection1))
// 	for _, item := range collection1 {
// 		if ContainsBy(collection2, item, equalityComparer) {
// 			result = append(result, item)
// 		}
// 	}
// 	return result
// }

// func IntersectBy2[TSource ~[]TObject, T IEquatable[T]](collection1, collection2 TSource) (result TSource) {
// 	result = make(TSource, 0, len(collection1))
// 	for _, item := range collection1 {
// 		if ContainsBy2(collection2, item) {
// 			result = append(result, item)
// 		}
// 	}
// 	return result
// }
