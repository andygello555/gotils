package numbers

import "fmt"

// Showcases scaling a value to a linear range, and an inverse linear range.
func ExampleScaleRange() {
	// Linear range
	fmt.Println(ScaleRange(100.0, 0.0, 200.0, 0.0, 1.0))
	// Inverse linear range
	fmt.Println(ScaleRange(100.0, 0.0, 100.0, 1.0, -1.0))
	// Output:
	// 0.5
	// -1
}

// Prints out an array of 8-bit integers that are clamped to a maximum of int8(2).
func ExampleClamp() {
	x := []int8{1, 2, 3, 4}
	for i, val := range x {
		x[i] = Clamp(val, 2)
	}
	fmt.Println(x)
	// Output:
	// [1 2 2 2]
}

// Prints out an array of 8-bit integers that are clamped to a minimum of int8(2).
func ExampleClampMin() {
	x := []int8{1, 2, 3, 4}
	for i, val := range x {
		x[i] = ClampMin(val, 2)
	}
	fmt.Println(x)
	// Output:
	// [2 2 3 4]
}

// Prints out an array of 8-bit integers that are clamped to a minimum of int8(2), and a maximum of int8(3).
func ExampleClampMinMax() {
	x := []int8{1, 2, 3, 4}
	for i, val := range x {
		x[i] = ClampMinMax(val, 2, 3)
	}
	fmt.Println(x)
	// Output:
	// [2 2 3 3]
}

// Prints out an array of positive numbers from an array of positive and negative numbers.
func ExampleAbs() {
	x := []float32{-1.0, 1.0, -0.001, 0.001}
	for i, val := range x {
		x[i] = Abs(val)
	}
	fmt.Println(x)
	// Output:
	// [1 1 0.001 0.001]
}

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

// Generates the ordinals only for 1, 2 and 3.
func ExampleOrdinalOnly() {
	fmt.Println(OrdinalOnly(1), OrdinalOnly(2), OrdinalOnly(3))
	// Output: st nd rd
}
