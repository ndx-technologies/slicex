package slicex

func Count[T comparable](vs []T, v T) int {
	count := 0
	for _, item := range vs {
		if item == v {
			count++
		}
	}
	return count
}
