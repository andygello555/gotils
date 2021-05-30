// String based data structures and string manipulation functions.
package strings

import (
	"github.com/andygello555/gotils/slices"
	"reflect"
	"sort"
	"strings"
	"unicode"
)

// A priority queue which sorts strings lexicographically.
type StringHeap []string

// The length of the StringHeap.
func (spq StringHeap) Len() int { return len(spq) }

// Compares two string elements in the StringHeap.
//
// Uses strings.Compare internally.
func (spq StringHeap) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return strings.Compare(spq[i], spq[j]) <= 0
}

// Swaps the two elements indicated via the given indices.
func (spq StringHeap) Swap(i, j int) { spq[i], spq[j] = spq[j], spq[i] }

// Pushes the given interface{} to the heap.
//
// Will panic if given interface{} is not a string.
func (spq *StringHeap) Push(x interface{}) { *spq = append(*spq, x.(string)) }

// Pops the tail of the queue.
func (spq *StringHeap) Pop() interface{} {
	old := *spq
	n := len(old)
	str := old[n-1]
	*spq = old[0 : n-1]
	return str
}

// Given a string, will strip all whitespace from it and return a new string without any whitespace.
func StripWhitespace(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

// Will replace all characters at the given indices with the new strings. Returns a new string.
//
// Indices are all the character indices with which to replace with the new string. Each occurrence will be replaced by
// the string new[occCount % len(new)]. Where occCount is the current count of the indices that have been replaced.
//
// The indices slice can contain duplicates and doesn't need to be sorted.
func ReplaceCharIndex(old string, indices []int, new... string) string {
	if len(indices) > 0 && len(new) > 0 {
		// Lets sort the indices and make them unique so that we can pop them in ascending order
		slices.RemoveDuplicatesAndSort(&indices)
		// Pop the first element
		var currIdx int
		currIdx, indices = indices[0], indices[1:]
		occCount := 0

		var b strings.Builder
		for idx, val := range old {
			if currIdx == idx {
				// If we have reached an index to replace then write the new string and pop the new idx
				b.WriteString(new[occCount % len(new)])
				occCount++
				if len(indices) > 0 {
					currIdx, indices = indices[0], indices[1:]
				}
			} else {
				// Otherwise write the current character
				b.WriteString(string(val))
			}
		}
		return b.String()
	}
	// If there is nothing to replace then return the old string
	return old
}

// Similar to ReplaceCharIndex but takes multiple index ranges in the form of [start, end].
//
// The length of new strings must be less than or equal to the length of the indices slice. The length of indices must
// also be greater than 0. If any of these conditions are not met the old string shall be returned.
//
// The indices slice can contain duplicates and doesn't need to be sorted. It's worth bearing in mind that removing the
// duplicates from the indices slice is O(n^2).
func ReplaceCharIndexRange(old string, indices [][]int, new... string) string {
	if len(indices) > 0 && len(new) <= len(indices) {
		// Remove duplicates from the indices slice
		// FIXME: Find a more efficient way of doing this. Wrapper for 2D []int with equality?
		newIndices := make([][]int, 0)
		for _, ran := range indices {
			for _, inSet := range newIndices {
				if ran[0] == inSet[0] && ran[1] == inSet[1] {
					// We continue onto the next element in indices without appending to newIndices
					goto skip
				}
			}
			newIndices = append(newIndices, ran)
			skip:
				continue
		}
		indices = newIndices

		// Sort the indices by ascending end values
		sort.SliceStable(indices, func(i, j int) bool {
			return indices[i][1] < indices[j][1]
		})

		// Pop the first element
		var currRange []int
		currRange, indices = indices[0], indices[1:]
		idxCount := 0

		var b strings.Builder
		idx := 0
		for idx < len(old) {
			if idx == currRange[0] {
				// Write the new string if we have just stumbled upon the start of the current range
				b.WriteString(new[idxCount % len(new)])
				idxCount++
				idx += currRange[1] - currRange[0]
				// Pop the new range if we still can
				if len(indices) > 0 {
					currRange, indices = indices[0], indices[1:]
				}
				continue
			}
			b.WriteString(string(old[idx]))
			idx++
		}
		return b.String()
	}
	// If there is nothing to replace then return the old string
	return old
}

// The TypeName of the given interface{}.
//
// If i is nil, "<nil>" will be returned.
func TypeName(i interface{}) string {
	if i == nil {
		return "<nil>"
	}
	return reflect.TypeOf(i).String()
}
