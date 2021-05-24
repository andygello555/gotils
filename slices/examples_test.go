package slices

import "fmt"

// Add the given element at the given indices.
func ExampleAddElems() {
	arr := []interface{}{1, 2, 3}
	fmt.Println("Before:", arr)

	// All duplicate indices will be removed.
	// Here the new length of the array will be 7 as it is greater than len(arr) + len(unique indices).
	arr = AddElems(arr, 0, 0, 0, 3, 7, 1)
	fmt.Println("After:", arr)
	// Output:
	// Before: [1 2 3]
	// After: [0 0 1 0 2 3 <nil> 0]
}

// Remove the given indices from an array.
func ExampleRemoveElems() {
	arr := []interface{}{1, 2, 3, 4, 5}
	fmt.Println("Before:", arr)

	// All duplicate indices will be removed.
	arr = RemoveElems(arr, 4, 4, 2, 1, 2)
	fmt.Println("After:", arr)
	// Output:
	// Before: [1 2 3 4 5]
	// After: [1 4]
}
