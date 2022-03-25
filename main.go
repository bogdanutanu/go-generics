package main

import "fmt"

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func SumGeneric[K comparable, V int64 | float64 | float32](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// Type constraint: interface, it's used to allow any type which implements the
// interface (in Go interfaces are implemented implicitly, unlike many other languages,
// where they have to be explicitly specified).
type Number interface {
	int64 | float64
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {

	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	floats32 := map[string]float32{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums:                       %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic sums with explicit type params: %v and %v\n",
		SumGeneric[string, int64](ints),
		SumGeneric[string, float64](floats))

	// The function can also be called without explicit type params. This is only possible
	// when the compiler is able to infer the type arguments from the types of function arguments.
	fmt.Printf("Generic sums:                           %v and %v and %v\n",
		SumGeneric(ints),
		SumGeneric(floats),
		SumGeneric(floats32))

	fmt.Printf("Generic Sums with Constraint:           %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))

	j()
}
