package slicex_test

import (
	"fmt"
	"time"

	"github.com/ndx-technologies/slicex"
)

func ExampleAny() {
	ints := []int{1, 2, 3}
	fmt.Print(slicex.Any(ints))
	// Output: [1 2 3]
}

func ExampleAny_empty() {
	var empty []int
	fmt.Print(slicex.Any(empty))
	// Output: []
}

func ExampleStrings() {
	items := []time.Duration{time.Second, time.Minute, time.Hour}
	fmt.Print(slicex.Strings(items))
	// Output: [1s 1m0s 1h0m0s]
}

func ExampleStrings_empty() {
	var empty []time.Duration
	fmt.Print(slicex.Strings(empty))
	// Output: []
}

func ExampleIsAllZero() {
	items := []time.Time{{}, {}, {}}
	fmt.Print(slicex.IsAllZero(items))
	// Output: true
}

func ExampleIsAllZero_mixed() {
	items := []time.Time{{}, time.Now(), {}}
	fmt.Print(slicex.IsAllZero(items))
	// Output: false
}

func ExampleIsAllZero_empty() {
	var empty []time.Time
	fmt.Print(slicex.IsAllZero(empty))
	// Output: true
}

func ExampleHasZero() {
	items := []time.Time{{}, {}, {}}
	fmt.Print(slicex.HasZero(items))
	// Output: true
}

func ExampleHasZero_noZeros() {
	items := []time.Time{time.Now(), time.Now(), time.Now()}
	fmt.Print(slicex.HasZero(items))
	// Output: false
}

func ExampleHasZero_empty() {
	var empty []time.Time
	fmt.Print(slicex.HasZero(empty))
	// Output: false
}
