package slices

func Filter[T any](arr []T, cb func(T, int) bool) []T {
	newArr := make([]T, 0)

	for i := 0; i < len(arr); i++ {
		if cb(arr[i], i) {
			newArr = append(newArr, arr[i])
		}
	}

	return newArr
}
