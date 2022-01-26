package main

import (
	"fmt"
	"sort"
)

// sort OMIT

func Sort(data sort.Interface) {} // HL

type Interface interface { // HL
	// Len is the number of elements in the collection.
	Len() int

	// Less reports whether the element with index i
	// must sort before the element with index j.
	Less(i, j int) bool

	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

// sort# OMIT
// sort-usage OMIT

type IntSlice []int // HL

func (s IntSlice) Len() int { // HL
	return len(s)
}

func (s IntSlice) Less(i, j int) bool { // HL
	return s[i] < s[j]
}

func (s IntSlice) Swap(i, j int) { // HL
	s[i], s[j] = s[j], s[i]
}

func ExampleIntSlice() {
	s := []int{4, -8, 15}
	sort.Sort(IntSlice(s)) // HL
	fmt.Println(s)         // [-8 4 15]
}

// sort-usage# OMIT
