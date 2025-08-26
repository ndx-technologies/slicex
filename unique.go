package slicex

func Unique[T comparable](vs []T) []T { return UniqueBy(vs, func(v T) T { return v }) }

// UniqueBy takes first occurrence of each element
func UniqueBy[T any, K comparable](vs []T, k func(T) K) []T {
	idx := make(map[K]int, len(vs))
	for i, v := range vs {
		k := k(v)
		if _, ok := idx[k]; !ok {
			idx[k] = i
		}
	}

	if len(idx) == len(vs) {
		return vs
	}

	result := make([]T, 0, len(idx))
	for i := range vs {
		if idx[k(vs[i])] != i {
			continue
		}
		result = append(result, vs[idx[k(vs[i])]])
	}

	return result
}
