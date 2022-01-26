package main

import "fmt"

// stringable-float OMIT

type StringableFloat interface { // HL
	~float32 | ~float64 // union of types
	String() string
}

type MyFloat float64 // satisfies StringableFloat type constraints

func (m MyFloat) String() string {
	return fmt.Sprintf("%e", m)
}

// stringable-float# OMIT
// stringable-float-usage OMIT

func stringifyFloat[T StringableFloat](f T) string { // HL
	return f.String()
}

func ExampleMyFloat() {
	var f MyFloat = 48151623.42
	s := stringifyFloat[MyFloat](f) // HL
	fmt.Println(s)                  // 4.815162e+07
}

// stringable-float-usage# OMIT
// lesser OMIT

// lesser# OMIT
