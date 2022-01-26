package main

import (
	"constraints"
	"errors"
	"fmt"
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
