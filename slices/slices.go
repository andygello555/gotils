// Package slices contains helpers for the manipulation of slices.
package slices

import (
	"fmt"
	"sort"
)

// SameElements checks if two interface slices have the same elements.
//
// Unlike reflect.DeepEqual this will not care about order.
func SameElements(x, y []any) bool {
	if len(x) != len(y) {
		return false
	}
	// create a map of string -> int
	diff := make(map[string]int, len(x))
	for _, _x := range x {
		// 0 value for int is 0, so just increment a counter for the value
		diff[fmt.Sprint(_x)]++
	}
	for _, _y := range y {
		// If the string _y is not in diff bail out early
		if _, ok := diff[fmt.Sprint(_y)]; !ok {
			return false
		}
		diff[fmt.Sprint(_y)] -= 1
		if diff[fmt.Sprint(_y)] == 0 {
			delete(diff, fmt.Sprint(_y))
		}
	}
	if len(diff) == 0 {
		return true
	}
	return false
}

// RemoveDuplicatesAndSort removes duplicates and sort an array of integers in place.
func RemoveDuplicatesAndSort(indices *[]int) {
	actualIndices := make([]int, 0)
	indexSet := make(map[int]struct{})

	for _, index := range *indices {
		// Check if the index is already in the array of actual indices. If not then we can add it
		if _, exists := indexSet[index]; !exists {
			actualIndices = append(actualIndices, index)
			indexSet[index] = struct{}{}
		}
	}

	// Sort the indices
	sort.Ints(actualIndices)
	*indices = actualIndices
}

// AddElems adds the given values at the given indices.
//
// If there is an index which exceeds the length of the given slice plus the number of unique indices given then this
// will result in a new array that's the length of the maximum index in indices. If this happens then each "empty"
// space will be filled by E's zero-value.
//
// Each given value will be inserted in their given order. If the number of given values aren't the same length as the
// number of (unique) indices, then the last given value will be inserted as many times as there are indices left. If
// the number of given values is greater than the number of given indices, all values that reside at an index that is
// greater than or equal to the number of given indices will be ignored. I.e. not inserted into the output slice. If
// there are no values given, then the zero-value will be inserted at all given indices.
func AddElems[E any](slice []E, values []E, indices ...int) []E {
	RemoveDuplicatesAndSort(&indices)
	// Find the bounds of the new array which will contain the appended value. This is either:
	// 1. The maximum index: when it exceeds the limits of the new array which will be the length of the slice plus the number of indices
	// 2. The length of the slice plus the number of indices: otherwise
	var high int
	if indices[len(indices)-1]+1 > len(slice)+len(indices) {
		high = indices[len(indices)-1] + 1
	} else {
		high = len(slice) + len(indices)
	}
	// Construct a new array from the specifications above
	newArr := make([]E, high)
	offset := 0

	var currIdx int
	currIdx, indices = indices[0], indices[1:]

	valuePtr := 0
	if len(values) == 0 {
		var x E
		values = append(values, x)
	}

	// Iterate from 0 to high inserting a value at each index to insert into
	for i := 0; i < high; i++ {
		if currIdx == i {
			if len(indices) > 0 {
				currIdx, indices = indices[0], indices[1:]
			}
			newArr[i] = values[valuePtr]
			if valuePtr < len(values)-1 {
				valuePtr++
			}
			offset += 1
			continue
		}
		if i-offset < len(slice) {
			newArr[i] = slice[i-offset]
		}
	}
	return newArr
}

// RemoveElems removes the elements at the given indices in the given slice and returns a new slice of that type.
//
// The new array will have a length which is the difference between the length of the given slice and the cardinality of
// the given indices as a unique set.
func RemoveElems[E any](slice []E, indices ...int) []E {
	RemoveDuplicatesAndSort(&indices)
	out := make([]E, 0)
	// Simple priority queue structure
	var currIdx int
	currIdx, indices = indices[0], indices[1:]

	for i, elem := range slice {
		if i == currIdx {
			if len(indices) > 0 {
				currIdx, indices = indices[0], indices[1:]
			}
			continue
		}
		out = append(out, elem)
	}
	return out
}

// Comprehension takes a list of elements of any type and runs the given function on each element to construct a new
// list with elements of the given type.
//
// The first parameter of the function is the index of the currently iterated element, the second is the currently
// iterated element's value, and the last is the input array in full.
func Comprehension[IE any, OE any](s []IE, fun func(idx int, value IE, arr []IE) OE) []OE {
	out := make([]OE, len(s))
	for i, ie := range s {
		out[i] = fun(i, ie, s)
	}
	return out
}

// Reverse reverses a slice in-place.
func Reverse[IE any](s []IE) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// ReverseOut reverses a slice in-place, then returns a reference to the slice so that it can be nested within other
// functions.
func ReverseOut[IE any](s []IE) []IE {
	Reverse(s)
	return s
}
