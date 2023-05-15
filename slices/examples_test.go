package slices

import (
	"fmt"
	"math"
	"reflect"
)

func ExampleSameElements() {
	// Two slices with the same elements but different orders
	arr1 := []any{1, 2, 3}
	arr2 := []any{2, 1, 3}
	fmt.Printf("SameElements(%v, %v) = %t\n", arr1, arr2, SameElements(arr1, arr2))

	// Two slices with the different elements
	arr1 = []any{1, 2, 4}
	arr2 = []any{2, 3, 1}
	fmt.Printf("SameElements(%v, %v) = %t\n", arr1, arr2, SameElements(arr1, arr2))

	// Output:
	// SameElements([1 2 3], [2 1 3]) = true
	// SameElements([1 2 4], [2 3 1]) = false
}

// Add the given element at the given indices.
func ExampleAddElems() {
	arr := []int{1, 2, 3}
	fmt.Println("Before:", arr)

	// All duplicate indices will be removed.
	// Here the new length of the array will be 7 as it is greater than len(arr) + len(unique indices).
	arr = AddElems(arr, []int{4}, 0, 0, 3, 7, 1)
	fmt.Println("After:", arr)

	fmt.Println(AddElems([]string{"Hello", "Gertrude"}, []string{"my", "name", "is"}, 1, 2, 3))
	// Output:
	// Before: [1 2 3]
	// After: [4 4 1 4 2 3 0 4]
	// [Hello my name is Gertrude]
}

// Remove the given indices from an array.
func ExampleRemoveElems() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Before:", arr)

	// All duplicate indices will be removed.
	arr = RemoveElems(arr, 4, 4, 2, 1, 2)
	fmt.Println("After:", arr)
	// Output:
	// Before: [1 2 3 4 5]
	// After: [1 4]
}

// Multiplies a list of integers by their relative index, then converts the resulting integer to its string
// representation.
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

// Multiples each integer in each given array by their relative index within the joined array. The result of this
// arithmetic will then be converted to the integer's string representation.
func ExampleJoinedComprehension() {
	intArrs := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("Before:", intArrs)
	fmt.Println("Before (joined):", Join(intArrs...))

	strArr := JoinedComprehension[int, string](func(idx int, value int, arr []int) string {
		return fmt.Sprintf("%d", idx*value)
	}, intArrs...)
	fmt.Println("After:", strArr)
	// Output:
	// Before: [[1 2 3] [4 5 6] [7 8 9]]
	// Before (joined): [1 2 3 4 5 6 7 8 9]
	// After: [0 2 6 12 20 30 42 56 72]
}

// Joins 3 lists of 3 elements each together into a new list of 9 elements.
func ExampleJoin() {
	fmt.Println(Join([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}))
	// Output:
	// [1 2 3 4 5 6 7 8 9]
}

// Reverses a slice of integers in-place.
func ExampleReverse() {
	intArr := []int{1, 2, 3, 4, 5}
	fmt.Println("Before:", intArr)
	Reverse(intArr)
	fmt.Println("After:", intArr)
	// Output:
	// Before: [1 2 3 4 5]
	// After: [5 4 3 2 1]
}

// Reverses a slice of integers, then converts it to a slice of float64s.
func ExampleReverseOut() {
	intArr := []int{1, 2, 3, 4, 5}
	fmt.Printf("Before (intArr): %v, type = %s\n", intArr, reflect.TypeOf(intArr).String())
	// ReverseOut returns a reference to the slice so that it can be nested.
	floatArr := Comprehension(ReverseOut(intArr), func(idx int, value int, arr []int) float64 {
		return float64(value)
	})
	fmt.Printf("After (intArr): %v, type = %s\n", intArr, reflect.TypeOf(intArr).String())
	fmt.Printf("After (floatArr): %v, type = %s\n", floatArr, reflect.TypeOf(floatArr).String())
	// Output:
	// Before (intArr): [1 2 3 4 5], type = []int
	// After (intArr): [5 4 3 2 1], type = []int
	// After (floatArr): [5 4 3 2 1], type = []float64
}

// Filters a list of integers.
func ExampleFilter() {
	fmt.Println(
		"square numbers (0-9):",
		Filter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(idx int, value int, arr []int) bool {
			if value >= 0 {
				sr := int(math.Sqrt(float64(value)))
				return sr*sr == value
			}
			return false
		},
		))
	// Output:
	// square numbers (0-9): [1 4 9]
}

// Shows a few examples on how to use the Any function, and how the default evaluation functions change depending on the
// type of the input array.
func ExampleAny() {
	checkStringLenEq10 := func(idx int, value string, arr []string) bool {
		return len(value) == 10
	}

	checkStringEqB := func(idx int, value string, arr []string) bool {
		return value == "b"
	}

	bools := []bool{false, true, false}
	numbers := []uint32{0, 0, 1, 3}
	strings := []string{"hello", "a", "world", "b"}
	emptyStrings := []string{"", "", ""}
	arrayOfArrays := [][]any{{}, {}, {1, 2, 3}}

	fmt.Printf("Any(%v) = %t\n", bools, Any(bools))
	fmt.Printf("Any(%v) = %t\n", numbers, Any(numbers))
	fmt.Printf("Any(%v, checkStringLenEq10, checkStringEqB) = %t\n", strings, Any(strings, checkStringLenEq10, checkStringEqB))
	fmt.Printf("Any(%v) = %t\n", emptyStrings, Any(emptyStrings))
	fmt.Printf("Any(%v) = %t\n", arrayOfArrays, Any(arrayOfArrays))
	// Output:
	// Any([false true false]) = true
	// Any([0 0 1 3]) = true
	// Any([hello a world b], checkStringLenEq10, checkStringEqB) = true
	// Any([  ]) = false
	// Any([[] [] [1 2 3]]) = true
}

// Shows a few examples on how to use the All function, and how the default evaluation functions change depending on the
// type of the input array.
func ExampleAll() {
	checkStringLenEq10 := func(idx int, value string, arr []string) bool {
		return len(value) == 10
	}

	checkStringEqB := func(idx int, value string, arr []string) bool {
		return value == "b"
	}

	bools := []bool{true, true, true}
	numbers := []uint32{0, 0, 1, 3}
	strings := []string{"helloworld", "b", "helloworld", "b"}
	arrayOfArrays := [][]any{{1}, {1, 2}, {1, 2, 3}}

	fmt.Printf("All(%v) = %t\n", bools, All(bools))
	fmt.Printf("All(%v) = %t\n", numbers, All(numbers))
	fmt.Printf("All(%v, checkStringLenEq10, checkStringEqB) = %t\n", strings, All(strings, checkStringLenEq10, checkStringEqB))
	fmt.Printf("All(%v) = %t\n", arrayOfArrays, All(arrayOfArrays))
	// Output:
	// All([true true true]) = true
	// All([0 0 1 3]) = false
	// All([helloworld b helloworld b], checkStringLenEq10, checkStringEqB) = true
	// All([[1] [1 2] [1 2 3]]) = true
}

// Orders a few different types of slices that can be ordered, as well as showcasing what happens when a slice cannot be
// ordered.
func ExampleOrder() {
	a := []uint8{4, 2, 1, 3}
	Order(a)
	b := []float32{3.142, 1.23, 2.222, 4.0}
	Order(b)
	c := []string{"3", "4", "2", "1"}
	Order(c)
	d := []bool{false, true}
	Order(d)

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
	// Output:
	// a: [1 2 3 4]
	// b: [1.23 2.222 3.142 4]
	// c: [1 2 3 4]
	// d: [false true]
}
