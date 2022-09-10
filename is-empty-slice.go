package tutils

func IsEmptySlice[T any](value []T) bool {
	return len(value) == 0
}
