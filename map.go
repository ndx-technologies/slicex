package slicex

func Any[T any](vs []T) []any {
	if len(vs) == 0 {
		return nil
	}
	res := make([]any, len(vs))
	for i, v := range vs {
		res[i] = v
	}
	return res
}

func Strings[T interface{ String() string }](vs []T) []string {
	if len(vs) == 0 {
		return nil
	}
	ss := make([]string, len(vs))
	for i, v := range vs {
		ss[i] = v.String()
	}
	return ss
}

func IsAllZero[T interface{ IsZero() bool }](vs []T) bool {
	for _, v := range vs {
		if !v.IsZero() {
			return false
		}
	}
	return true
}

func HasZero[T interface{ IsZero() bool }](vs []T) bool {
	for _, v := range vs {
		if v.IsZero() {
			return true
		}
	}
	return false
}
