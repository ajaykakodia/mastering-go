// define how to use sort.Interface for slice to sort data
/*
	type Interface interface {
		// Length for number of element in given slice
		Len() int
		// Less reports whether the element with index i should sort before the element with index j
		Less(i, j int) bool
		// Swap swaps the elements with indexes i and j
		Swap(i, j int)
	}
*/
package main

import (
	"fmt"
	"sort"
)

type S1 struct {
	f1 int
	f2 string
	f3 int
}

type S2 struct {
	f1 int
	f2 string
	f3 S1
}

type S2Slice []S2

func (s S2Slice) Len() int {
	return len(s)
}

func (s S2Slice) Less(i, j int) bool {
	return s[i].f3.f1 < s[j].f3.f1
}

func (s S2Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	data := []S2{
		{1, "One", S1{1, "S1_1", 10}},
		{2, "Two", S1{2, "S1_1", 20}},
		{3, "Three", S1{-1, "S1_1", -20}},
	}
	fmt.Println("Before: ", data)
	sort.Sort(S2Slice(data))
	fmt.Println("After Sort: ", data)

	// Reverse sorting works automatically
	sort.Sort(sort.Reverse(S2Slice(data)))
	fmt.Println("Reverse: ", data)
}
