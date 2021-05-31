package ints

import "fmt"

// Counting down from 10 in intervals of 2.
//
// The Range function is similar to the range function in Python.
func ExampleRange() {
	r := Range(10, 0, -2)
	fmt.Println(r)
	// Output:
	// [10 8 6 4 2 0]
}

// Generates the ordinals for 1, 2 and 3.
func ExampleOrdinal() {
	fmt.Println(Ordinal(1), Ordinal(2), Ordinal(3))
	// Output: 1st 2nd 3rd
}
