package utils

import "sync"

func MapAsync[T any, R any](items []T, transform func(T) R) []R {
	result := make([]R, len(items))

	var wg sync.WaitGroup
	wg.Add(len(items))

	for i, item := range items {
		go func(i int, item T) {
			defer wg.Done()
			result[i] = transform(item)
		}(i, item)
	}

	wg.Wait()
	return result
}
