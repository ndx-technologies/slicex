package slicex

import "strings"

func Parser[V any, P interface {
	*V
	UnmarshalText(b []byte) error
}](vs *[]V) func(string) error {
	return func(s string) error {
		if len(s) == 0 {
			return nil
		}
		for p := range strings.SplitSeq(s, ",") {
			var v V
			if err := P(&v).UnmarshalText([]byte(p)); err != nil {
				return err
			}
			*vs = append(*vs, v)
		}
		return nil
	}
}
