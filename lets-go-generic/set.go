package main

import "fmt"

// set1 OMIT

// Set implements generic set data structure backed by a hash table.
// It is not thread safe.
type Set[T comparable] struct { // HL
	values map[T]struct{}
}

func NewSet[T comparable](values ...T) *Set[T] { // HL
	m := make(map[T]struct{}, len(values))
	for _, v := range values {
		m[v] = struct{}{}
	}
	return &Set[T]{
		values: m,
	}
}

// set1# OMIT
// set2 OMIT

func (s *Set[T]) Add(values ...T) { // HL
	for _, v := range values {
		s.values[v] = struct{}{}
	}
}

func (s *Set[T]) Remove(values ...T) { // HL
	for _, v := range values {
		delete(s.values, v)
	}
}

func (s *Set[T]) Contains(values ...T) bool { // HL
	for _, v := range values {
		_, ok := s.values[v]
		if !ok {
			return false
		}
	}
	return true
}

// set2# OMIT
// set3 OMIT

func (s *Set[T]) Union(other *Set[T]) *Set[T] { // HL
	result := NewSet[T](s.Values()...)
	for _, v := range other.Values() {
		if !result.Contains(v) {
			result.Add(v)
		}
	}
	return result
}

func (s *Set[T]) Intersect(other *Set[T]) *Set[T] { // HL
	// pass smaller set first for optimization
	if s.Size() < other.Size() {
		return intersect(s, other)
	}
	return intersect(other, s)
}

// set3# OMIT
// set4 OMIT

// intersect returns intersection of given sets. It iterates over smaller set for optimization.
func intersect[T comparable](smaller, bigger *Set[T]) *Set[T] { // HL
	result := NewSet[T]()
	for k, _ := range smaller.values {
		if bigger.Contains(k) {
			result.Add(k)
		}
	}
	return result
}

func (s *Set[T]) Values() []T { // HL
	return s.toSlice()
}

// set4# OMIT
// set5 OMIT

func (s *Set[T]) Size() int { // HL
	return len(s.values)
}

func (s *Set[T]) Clear() { // HL
	s.values = map[T]struct{}{}
}

func (s *Set[T]) String() string { // HL
	return fmt.Sprint(s.toSlice())
}

func (s *Set[T]) toSlice() []T {
	result := make([]T, 0, len(s.values))
	for k := range s.values {
		result = append(result, k)
	}
	return result
}

// set5# OMIT
// set-usage OMIT

func ExampleSet() {
	s1 := NewSet(4, 4, -8, 15)               // HL
	s2 := NewSet("foo", "foo", "bar", "baz") // HL
	fmt.Println(s1.Size(), s2.Size())        // 3, 3

	s1.Add(-16)
	s2.Add("hoge")
	fmt.Println(s1.Size(), s2.Size())                  // 4, 4
	fmt.Println(s1.Contains(-16), s2.Contains("hoge")) // true, true

	s1.Remove(15)
	s2.Remove("baz")
	fmt.Println(s1.Size(), s2.Size()) // 3, 3

	fmt.Println(len(s1.Values()), len(s2.Values())) // 3, 3

	s3 := NewSet("hoge", "dragon", "fly") // HL
	fmt.Println(s2.Union(s3).Size())      // 5
	fmt.Println(s2.Intersect(s3))         // [hoge]

	s1.Clear()
	s2.Clear()
	fmt.Println(s1.Size(), s2.Size()) // 0, 0
}

// set-usage# OMIT
