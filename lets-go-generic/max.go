package main

import (
	"errors"
	"fmt"
	"reflect"

	"golang.org/x/exp/constraints"
)

func main() {}

// max-int OMIT

func MaxInt(s []int) int { // HL
	if len(s) == 0 {
		return 0
	}

	max := s[0]
	for _, v := range s[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func ExampleMaxInt() {
	m := MaxInt([]int{4, -8, 15}) // HL
	fmt.Println(m)                // 15
}

// max-int# OMIT
// max-float64 OMIT

func MaxFloat64(s []float64) float64 { // HL
	if len(s) == 0 {
		return 0
	}

	max := s[0]
	for _, v := range s[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func ExampleMaxFloat64() {
	m := MaxFloat64([]float64{4.1, -8.1, 15.1}) // HL
	fmt.Println(m)                              // 15.1
}

// max-float64# OMIT
// max-number1 OMIT

func MaxNumber(s []interface{}) (interface{}, error) { // HL
	if len(s) == 0 {
		return nil, errors.New("no values given")
	}
	// ...
	// max-number1# OMIT
	// max-number2 OMIT
	// ...
	switch first := s[0].(type) { // HL
	case int: // HL
		max := first
		for _, rawV := range s[1:] {
			v := rawV.(int) // HL
			if v > max {
				max = v
			}
		}
		return max, nil
	case float64: // HL
		max := first
		for _, rawV := range s[1:] {
			v := rawV.(float64) // HL
			if v > max {
				max = v
			}
		}
		return max, nil
	default:
		return nil, fmt.Errorf("unsupported element type of given slice: %T", first)
	}
}

// max-number2# OMIT
// max-number-usage OMIT

func ExampleMaxNumber() {
	m1, err1 := MaxNumber([]interface{}{4, -8, 15})       // HL
	m2, err2 := MaxNumber([]interface{}{4.1, -8.1, 15.1}) // HL
	fmt.Println(err1, err2)                               // <nil> <nil>
	fmt.Println(m1, m2)                                   // 15 15.1
}

// max-number-usage# OMIT
// max-number-defined-if1 OMIT

// ComparableSlice is a slice that holds elements that can be compared.
type ComparableSlice interface { // HL
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with index i has lower value than the element with index j.
	Less(i, j int) bool
	// Elem returns the element with index i.
	Elem(i int) interface{}
}

func MaxNumber2(s ComparableSlice) (interface{}, error) { // HL
	if s.Len() == 0 {
		return nil, errors.New("no values given")
	}

	max := s.Elem(0)
	for i := 1; i < s.Len(); i++ {
		if s.Less(i-1, i) {
			max = s.Elem(i)
		}
	}

	return max, nil
}

// max-number-defined-if1# OMIT
// max-number-defined-if2 OMIT

type ComparableIntSlice []int

func (s ComparableIntSlice) Len() int               { return len(s) }
func (s ComparableIntSlice) Less(i, j int) bool     { return s[i] < s[j] }
func (s ComparableIntSlice) Elem(i int) interface{} { return s[i] }

type ComparableFloat64Slice []float64

func (s ComparableFloat64Slice) Len() int               { return len(s) }
func (s ComparableFloat64Slice) Less(i, j int) bool     { return s[i] < s[j] }
func (s ComparableFloat64Slice) Elem(i int) interface{} { return s[i] }

func ExampleMaxNumber2() {
	m1, err1 := MaxNumber2(ComparableIntSlice([]int{4, -8, 15}))               // HL
	m2, err2 := MaxNumber2(ComparableFloat64Slice([]float64{4.1, -8.1, 15.1})) // HL
	fmt.Println(err1, err2)                                                    // <nil> <nil>
	fmt.Println(m1, m2)                                                        // 15 15.1
}

// max-number-defined-if2# OMIT

// max-number-reflect1 OMIT

func MaxNumber3(s []interface{}) (interface{}, error) { // HL
	if len(s) == 0 {
		return nil, errors.New("no values given")
	}

	first := reflect.ValueOf(s[0])
	if first.CanInt() {
		max := first.Int()
		for _, ifV := range s[1:] {
			v := reflect.ValueOf(ifV)
			if v.CanInt() {
				intV := v.Int()
				if intV > max {
					max = intV
				}
			}
		}
		return max, nil
	}

	// max-number-reflect2 OMIT
	// ...
	// max-number-reflect1# OMIT

	if first.CanFloat() {
		max := first.Float()
		for _, ifV := range s[1:] {
			v := reflect.ValueOf(ifV)
			if v.CanFloat() {
				intV := v.Float()
				if intV > max {
					max = intV
				}
			}
		}
		return max, nil
	}

	return nil, fmt.Errorf("unsupported element type of given slice: %T", s[0])
}

// max-number-reflect2# OMIT
// max-number-reflect-usage OMIT

func ExampleMaxNumber3() {
	m1, err1 := MaxNumber3([]interface{}{4, -8, 15})       // HL
	m2, err2 := MaxNumber3([]interface{}{4.1, -8.1, 15.1}) // HL
	fmt.Println(err1, err2)                                // <nil> <nil>
	fmt.Println(m1, m2)                                    // 15 15.1
}

// max-number-reflect-usage# OMIT

// max-generic OMIT

func Max[T constraints.Ordered](s []T) T { // HL
	if len(s) == 0 {
		return *new(T)
	}

	max := s[0]
	for _, v := range s[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func ExampleMax() {
	m1 := Max[int]([]int{4, -8, 15})      // HL
	m2 := Max([]float64{4.1, -8.1, 15.1}) // HL
	fmt.Println(m1, m2)                   // 15 15.1
}

// max-generic# OMIT
