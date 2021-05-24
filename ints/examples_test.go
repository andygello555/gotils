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
