package mrutil

// SetContains O(1)
func SetContains[M ~map[K]bool, K comparable](m M, key K) bool {
	val, ok := m[key]
	return ok && val
}

func ToSet[A ~[]K, K comparable](a A) map[K]bool {
	retMap := make(map[K]bool)
	for _, k := range a {
		retMap[k] = true
	}
	return retMap
}

func NewSet[K comparable]() map[K]bool {
	return make(map[K]bool)
}

// SetToArray O(n)
func SetToArray[M ~map[K]bool, K comparable](m M) []K {
	retArray := make([]K, len(m))
	for k, v := range m {
		if v {
			retArray = append(retArray, k)
		}
	}
	return retArray
}
