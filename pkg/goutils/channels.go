package goutils

func Collect[T any](c chan T) []T {
	retSlice := make([]T, 0)
	for elem := range c {
		retSlice = append(retSlice, elem)
	}
	return retSlice
}
