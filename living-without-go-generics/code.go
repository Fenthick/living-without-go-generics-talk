package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {}

// built-ins OMIT

// keys returns slice of keys of given map.
func keys(m map[string]int) []string { // HL
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

// built-ins# OMIT
// built-ins2 OMIT

// keys2 returns slice of keys of given map.
func keys2(m map[int]string) []int { // HL
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

// built-ins2# OMIT
// reader OMIT

type Reader interface {
	Read(p []byte) (n int, err error)
}

// reader# OMIT
// file OMIT

type File struct {
	// ...
}

func (f *File) Read(b []byte) (n int, err error) {
	// ...
	return 0, nil // OMIT
}

// file# OMIT
// file usage OMIT

func readAndDoSomething(f *os.File) { // HL
	var b []byte
	n, err := f.Read(b)
	// do something
	fmt.Print(n, err) // OMIT
}

// file usage# OMIT
// file usage2 OMIT

func readAndDoSomething2(f io.Reader) { // HL
	var b []byte
	n, err := f.Read(b)
	// do something
	fmt.Print(n, err) // OMIT
}

// file usage2# OMIT
// student OMIT

type Student struct{}

func (s *Student) AverageScore() float64 {
	// calculate average score
	return 0 // OMIT
}

func sortStudentsByAverageScore(students []Student) {
	// needs to be implemented // HL
}

// student# OMIT
// sort OMIT

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

func Sort(data sort.Interface) {
	// implementation omitted
}

// sort# OMIT
// sort solution OMIT

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Less(i, j int) bool {
	return s[i].AverageScore() < s[j].AverageScore()
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func sortStudentsByAverageScore2(students []Student) {
	sort.Sort(Students(students))
}

// sort solution# OMIT
// sort2 OMIT
// slice OMIT

// Slice sorts the provided slice given the provided less function.
// The function panics if the provided interface is not a slice.
func Slice(slice interface{}, less func(i, j int) bool) {
	// implementation omitted
}

// slice# OMIT

func sortStudentsByAverageScore3(students []Student) {
	sort.Slice(students, func(i int, j int) bool {
		return students[i].AverageScore() < students[j].AverageScore()
	})
}

// sort2# OMIT
