package slicex

func SelectMask[T any](vs []T, keep []bool) []T {
	if len(vs) != len(keep) {
		return vs
	}

	count := Count(keep, true)

	if count == 0 {
		return nil
	}

	if len(vs) == count {
		return vs
	}

	res := make([]T, 0, count)
	for i := range vs {
		if keep[i] {
			res = append(res, vs[i])
		}
	}

	return res
}

func Select[T any](vs []T, keep func(v T) bool) []T {
	mask := make([]bool, len(vs))
	for i, v := range vs {
		mask[i] = keep(v)
	}
	return SelectMask(vs, mask)
}
