package slices

import "sort"

// Remove duplicates and sort an array of integers in place.
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

// Adds the given value at the given indices.
//
// If there is an index which exceeds the length of the given slice plus the number of unique indices given then this
// will result in an new array that's the length of the maximum index in indices. If this happens then any "empty"
// space will be filled by default by "nil".
func AddElems(slice []interface{}, value interface{}, indices... int) []interface{} {
	RemoveDuplicatesAndSort(&indices)
	// Find the bounds of the new array which will contain the appended value. This is either:
	// 1. The maximum index: when it exceeds the limits of the new array which will be the length of the slice plus the number of indices
	// 2. The length of the slice plus the number of indices: otherwise
	var high int
	if indices[len(indices) - 1] + 1 > len(slice) + len(indices) {
		high = indices[len(indices) - 1] + 1
	} else {
		high = len(slice) + len(indices)
	}
	// Construct a new array from the specifications above
	newArr := make([]interface{}, high)
	offset := 0

	var currIdx int
	currIdx, indices = indices[0], indices[1:]

	// Iterate from 0 to high inserting a value at each index to insert into
	for i := 0; i < high; i++ {
		if currIdx == i {
			if len(indices) > 0 {
				currIdx, indices = indices[0], indices[1:]
			}
			newArr[i] = value
			offset += 1
			continue
		}
		if i - offset < len(slice) {
			newArr[i] = slice[i - offset]
		}
	}
	return newArr
}

// Removes the elements at the given indices in the given interface slice and returns a new slice.
//
// The new array will have a length which is the difference between the length of the given slice and the length of the
// given indices as a unique set.
func RemoveElems(slice []interface{}, indices... int) []interface{} {
	RemoveDuplicatesAndSort(&indices)
	out := make([]interface{}, 0)
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
