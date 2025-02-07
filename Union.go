package slices

import "github.com/thereisnoplanb/generic"

func Union[TSource ~[]TObject, TObject generic.Equatable](collection1, collection2 TSource) (result TSource) {
	result = append(result, collection1...)
	result = append(result, collection2...)
	return Distinct(result)
}

// func UnionBy[TSource ~[]TObject, TObject any](collection1, collection2 TSource, equalityComparer EqualityComparer[T]) (result TSource) {
// 	result = append(result, collection1...)
// 	result = append(result, collection2...)
// 	return DistinctBy(result, equalityComparer)
// }

// func UnionBy2[TSource ~[]TObject, T IEquatable[T]](collection1, collection2 TSource) (result TSource) {
// 	result = append(result, collection1...)
// 	result = append(result, collection2...)
// 	return DistinctByEquatable(result)
// }

// func Zip[TCollection1 ~[]T1, T1 any, TCollection2 ~[]T2, T2 any](collection1 TCollection1, collection2 TCollection2) (result TCollection3) {
// 	length := min(len(collection1),len(collection2))
// 	result = make(TSource, length)
// 	for i:=0; i < length; i++ {
// 		result[]
// 	}

// 	result = append(result, collection1...)
// 	result = append(result, collection2...)
// 	return Distinct(result)
// }
