package utils

func Reverse[T any](source []T) []T {
	length := len(source)
	result := make([]T, length)
	for i, item := range source {
		result[length-i-1] = item
	}
	return result
}

func AreEqual[T comparable](s1 []T, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i, v1 := range s1 {
		if s2[i] != v1 {
			return false
		}
	}

	return true
}
