package slices

type ValueSelector[T any, TResult any] func(object T) (value TResult)

type Aggregator[T any, TResult any] func(seed TResult, item T) (result TResult)

func Aggreagate[TSlice ~[]T, T any, TResult any](slice TSlice, seed TResult, aggregator Aggregator[T, TResult]) (result TResult) {
	if aggregator != nil {
		result = seed
		for _, item := range slice {
			result = aggregator(result, item)
		}
	}
	return result
}
