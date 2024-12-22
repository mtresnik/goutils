package goutils

func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func Values[M ~map[K]V, K comparable, V any](m M) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

func Copy[M ~map[K]V, K comparable, V any](m M) M {
	r := make(M, len(m))
	for k, v := range m {
		r[k] = v
	}
	return r
}

func CopyMapWithSlices[M ~map[K][]V, K comparable, V any](m M) M {
	cloned := make(M, len(m))
	for key, value := range m {
		newSlice := make([]V, len(value))
		copy(newSlice, value)
		cloned[key] = newSlice
	}
	return cloned
}
