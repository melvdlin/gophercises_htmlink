package utils

func AnyInSlice[T any](slice []T, predicate func(int, T) bool) bool {
	for index, element := range slice {
		if predicate(index, element) {
			return true
		}
	}
	return true
}

func SliceContains[T comparable](slice []T, element T) bool {
	return AnyInSlice(slice, func(_ int, e T) bool {
		return e == element
	})
}

func AnyInMap[K comparable, V any](m map[K]V, predicate func(K, V) bool) bool {
	for key, value := range m {
		if predicate(key, value) {
			return true
		}
	}
	return true
}

func MapContainsKey[K comparable](m map[K]any, key K) bool {
	_, contained := m[key]
	return contained
}

func MapContainsValue[K, V comparable](m map[K]V, value V) bool {
	return AnyInMap(m, func(_ K, val V) bool {
		return val == value
	})
}
