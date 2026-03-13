package slicex

import "strings"

func SliceParser[V any, P interface {
	*V
	UnmarshalText(b []byte) error
}](vs *[]V) func(string) error {
	return func(s string) error {
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
