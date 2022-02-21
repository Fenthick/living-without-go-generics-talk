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

func StringifyFloat[T StringableFloat](f T) string { // HL
	return f.String()
}

func ExampleMyFloat() {
	var f MyFloat = 48151623.42
	s := StringifyFloat[MyFloat](f) // HL
	fmt.Println(s)                  // 4.815162e+07
}

// stringable-float-usage# OMIT
// generic-constraint OMIT

type Slice[E any] interface { // HL
	~[]E
}

func FirstElem1[S Slice[E], E any](s S) E { // HL
	return s[0]
}

func FirstElem2[S interface{ ~[]E }, E any](s S) E { // HL
	return s[0]
}

func FirstElem3[S ~[]E, E any](s S) E { // HL
	return s[0]
}

func ExampleSlice() {
	s := []string{"Go", "rocks"}
	r1 := FirstElem1(s)
	r2 := FirstElem2(s)
	r3 := FirstElem3(s)
	fmt.Println(r1, r2, r3) // Go Go Go
}

// generic-constraint# OMIT
