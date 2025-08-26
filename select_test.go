package slicex_test

import (
	"fmt"

	"github.com/ndx-technologies/slicex"
)

func ExampleSelectMask() {
	vs := []int{1, 2, 3, 4, 5}
	keep := []bool{true, false, true, false, true}
	fmt.Print(slicex.SelectMask(vs, keep))
	// Output: [1 3 5]
}

func ExampleSelectMask_none() {
	vs := []int{1, 2, 3, 4, 5}
	keep := []bool{false, false, false, false, false}
	fmt.Print(slicex.SelectMask(vs, keep))
	// Output: []
}

func ExampleSelectMask_all() {
	vs := []int{1, 2, 3, 4, 5}
	keep := []bool{true, true, true, true, true}
	fmt.Print(slicex.SelectMask(vs, keep))
	// Output: [1 2 3 4 5]
}

func ExampleSelectMask_brokenMaskKeeps() {
	vs := []int{1, 2, 3, 4, 5}
	keep := []bool{true, true}
	fmt.Print(slicex.SelectMask(vs, keep))
	// Output: [1 2 3 4 5]
}

func ExampleSelect() {
	vs := []int{1, 2, 3, 4, 5}
	fmt.Print(slicex.Select(vs, func(v int) bool { return v%2 == 0 }))
	// Output: [2 4]
}
