package slices

import "fmt"

func ExampleSameElements() {
	// Two slices with the same elements but different orders
	arr1 := []interface{}{1, 2, 3}
	arr2 := []interface{}{2, 1, 3}
	fmt.Printf("SameElements(%v, %v) = %t\n", arr1, arr2, SameElements(arr1, arr2))

	// Two slices with the different elements
	arr1 = []interface{}{1, 2, 4}
	arr2 = []interface{}{2, 3, 1}
	fmt.Printf("SameElements(%v, %v) = %t\n", arr1, arr2, SameElements(arr1, arr2))

	// Output:
	// SameElements([1 2 3], [2 1 3]) = true
	// SameElements([1 2 4], [2 3 1]) = false
}

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

// Perform list comprehension on a given array/slice.
func ExampleComprehension() {
	intArr := []int{1, 2, 3, 4, 5}
	fmt.Println("Before:", intArr)

	strArr := Comprehension[int, string](intArr, func(idx int, value int, arr []int) string {
		return fmt.Sprintf("%d", idx*value)
	})
	fmt.Println("After:", strArr)
	// Output:
	// Before: [1 2 3 4 5]
	// After: [0 2 6 12 20]
}
