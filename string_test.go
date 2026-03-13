package slicex_test

import (
	"fmt"
	"strconv"

	"github.com/ndx-technologies/slicex"
)

type V int

func (s *V) UnmarshalText(b []byte) error {
	v, err := strconv.Atoi(string(b))
	*s = V(v)
	return err
}

func ExampleSliceParser() {
	var vs []V
	parse := slicex.SliceParser(&vs)
	parse("1,2,3")
	fmt.Print(vs)
	// Output: [1 2 3]
}

func ExampleSliceParser_empty() {
	var vs []V
	parse := slicex.SliceParser(&vs)
	parse("")
	fmt.Print(vs)
	// Output: []
}

func ExampleSliceParser_single() {
	var vs []V
	parse := slicex.SliceParser(&vs)
	parse("1")
	fmt.Print(vs)
	// Output: [1]
}
