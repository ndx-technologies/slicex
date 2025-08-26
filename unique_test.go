package slicex_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/ndx-technologies/slicex"
)

func TestUniqueByPerson(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	tests := []struct {
		name   string
		vs     []Person
		unique []Person
	}{
		{
			name: "no duplicates",
			vs: []Person{
				{Name: "Alice", Age: 25},
				{Name: "Bob", Age: 30},
				{Name: "Charlie", Age: 35},
			},
			unique: []Person{
				{Name: "Alice", Age: 25},
				{Name: "Bob", Age: 30},
				{Name: "Charlie", Age: 35},
			},
		},
		{
			name: "duplicates present",
			vs: []Person{
				{Name: "Alice", Age: 25},
				{Name: "Bob", Age: 30},
				{Name: "Alice", Age: 26},
				{Name: "Charlie", Age: 35},
				{Name: "Bob", Age: 31},
			},
			unique: []Person{
				{Name: "Alice", Age: 25},
				{Name: "Bob", Age: 30},
				{Name: "Charlie", Age: 35},
			},
		},
		{
			name: "all duplicates",
			vs: []Person{
				{Name: "Alice", Age: 25},
				{Name: "Alice", Age: 26},
				{Name: "Alice", Age: 27},
			},
			unique: []Person{
				{Name: "Alice", Age: 25},
			},
		},
		{
			name:   "empty slice",
			vs:     nil,
			unique: nil,
		},
		{
			name: "single element",
			vs: []Person{
				{Name: "Alice", Age: 25},
			},
			unique: []Person{
				{Name: "Alice", Age: 25},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vs := slicex.UniqueBy(tt.vs, func(p Person) string { return p.Name })

			if !slices.Equal(vs, tt.unique) {
				t.Error(vs, tt.unique)
			}
		})
	}
}

func generateDataString(size int, duplicates bool) []string {
	data := make([]string, size)
	if duplicates {
		for i := range size {
			data[i] = fmt.Sprintf("item-%d", i/2)
		}
	} else {
		for i := range size {
			data[i] = fmt.Sprintf("item-%d", i)
		}
	}
	return data
}

func generateDataInt(size int, duplicates bool) []int {
	data := make([]int, size)
	if duplicates {
		for i := range size {
			data[i] = i / 2
		}
	} else {
		for i := range size {
			data[i] = i
		}
	}
	return data
}

func BenchmarkUniqueBy(b *testing.B) {
	tests := []struct {
		size       int
		duplicates bool
	}{
		{size: 1, duplicates: false},
		{size: 3, duplicates: false},
		{size: 5, duplicates: false},
		{size: 7, duplicates: false},
		{size: 9, duplicates: false},
		{size: 10, duplicates: false},
		{size: 100, duplicates: false},
		{size: 100, duplicates: true},
		{size: 1000000, duplicates: false},
		{size: 1000000, duplicates: true},
	}
	for _, tc := range tests {
		b.Run(fmt.Sprintf("string/size=%d/duplicates=%t", tc.size, tc.duplicates), func(b *testing.B) {
			data := generateDataString(tc.size, tc.duplicates)

			for b.Loop() {
				slicex.UniqueBy(data, func(s string) string { return s })
			}
		})

		b.Run(fmt.Sprintf("int/size=%d/duplicates=%t", tc.size, tc.duplicates), func(b *testing.B) {
			data := generateDataInt(tc.size, tc.duplicates)

			for b.Loop() {
				slicex.UniqueBy(data, func(s int) int { return s })
			}
		})
	}
}
