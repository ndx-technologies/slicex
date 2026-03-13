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

func ExampleParser() {
	var vs []V
	f := slicex.Parser(&vs)
	f("1,2,3")
	fmt.Print(vs)
	// Output: [1 2 3]
}

func ExampleParser_empty() {
	var vs []V
	f := slicex.Parser(&vs)
	err := f("")
	fmt.Print(vs, err)
	// Output: [] <nil>
}

func ExampleParser_single() {
	var vs []V
	f := slicex.Parser(&vs)
	f("1")
	fmt.Print(vs)
	// Output: [1]
}

func ExampleParser_error() {
	var vs []V
	f := slicex.Parser(&vs)
	err := f("a")
	fmt.Print(vs, err != nil)
	// Output: [] true

}
