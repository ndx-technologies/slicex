package slicex_test

import (
	"fmt"

	"github.com/ndx-technologies/slicex"
)

func ExampleCount() {
	count := slicex.Count([]int{1, 2, 3, 1, 2, 3}, 2)
	fmt.Print(count)
	// Output: 2
}
