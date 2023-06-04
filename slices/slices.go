// Package slices contains helpers for the manipulation of slices.
package slices

import (
	"fmt"
	"github.com/andygello555/gotils/v2/misc"
	"github.com/andygello555/gotils/v2/numbers"
	"reflect"
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
// Each given value will be inserted in their given order. If the number of given values is less than the length of the
// number of (unique) indices, then the last given value will be inserted as many times as there are indices left.
//
// If the number of given values is greater than the number of given indices, all values that reside at an index that
// are greater than or equal to the number of given indices will be ignored. I.e. not inserted into the output slice.
//
// If there are no values given, then the zero-value will be inserted at all given indices.
//
// If there are no indices given, then a copy of slice will be returned.
func AddElems[E any](slice []E, values []E, indices ...int) []E {
	RemoveDuplicatesAndSort(&indices)
	// Find the bounds of the new array which will contain the appended value. This is either:
	// 1. The length of the slice: if there are no indices
	// 2. The maximum index: when it exceeds the limits of the new array which will be the length of the slice plus the number of indices
	// 3. The length of the slice plus the number of indices: otherwise
	var high int
	if len(indices) == 0 {
		high = len(slice)
	} else if indices[len(indices)-1]+1 > len(slice)+len(indices) {
		high = indices[len(indices)-1] + 1
	} else {
		high = len(slice) + len(indices)
	}
	// Construct a new array from the specifications above
	newArr := make([]E, high)
	offset := 0

	currIdx := len(newArr)
	if len(indices) > 0 {
		currIdx, indices = indices[0], indices[1:]
	}

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
//
// If no indices are given, then a copy of the slice will be returned.
func RemoveElems[E any](slice []E, indices ...int) []E {
	RemoveDuplicatesAndSort(&indices)
	out := make([]E, 0)

	// Simple priority queue structure
	currIdx := len(slice)
	if len(indices) > 0 {
		currIdx, indices = indices[0], indices[1:]
	}

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

// Join returns a new array of the given type that is created by joining/concatenating the elements of each given slice
// in their given order.
//
// If no arrays are given, then an empty slice is returned.
func Join[E any](ss ...[]E) []E {
	outLen := numbers.Sum(Comprehension(ss, func(idx int, value []E, arr [][]E) int {
		return len(value)
	})...)
	out := make([]E, outLen, outLen)

	if outLen > 0 {
		start := 0
		for _, s := range ss {
			end := start + len(s)
			copy(out[start:end], s)
			start = end
		}
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

// JoinedComprehension performs Comprehension on an array created from the Join of the given arrays.
func JoinedComprehension[IE any, OE any](fun func(idx int, value IE, arr []IE) OE, ss ...[]IE) []OE {
	return Comprehension(Join(ss...), fun)
}

// Filter takes a list of elements of any type and runs the given predicate function on each element. If the predicate
// returns true for an element, then that element will be added to a new list of the same type as the one provided.
//
// The first parameter of the predicate is the index of the currently iterated element within the given slice. The second
// is the currently iterated element's values, and the last is the input array in full.
func Filter[E any](s []E, fun func(idx int, value E, arr []E) bool) []E {
	out := make([]E, 0, len(s))
	for i, ie := range s {
		if fun(i, ie, s) {
			out = append(out, ie)
		}
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

func emptyFuncsResolve[E any]() (funcs []func(idx int, value E, arr []E) bool) {
	funcs = make([]func(idx int, value E, arr []E) bool, 1, 1)
	e := any(new(E))
	switch e.(type) {
	case *bool:
		funcs[0] = func(idx int, value E, arr []E) bool {
			return any(value).(bool)
		}
	case *string:
		funcs[0] = func(idx int, value E, arr []E) bool {
			return any(value).(string) != ""
		}
	default:
		eVal := reflect.ValueOf(e).Elem()
		eType := eVal.Type()
		switch {
		case eVal.CanInt(), eVal.CanUint(), eVal.CanFloat():
			funcs[0] = func(idx int, value E, arr []E) bool {
				val := reflect.ValueOf(value)
				return val.Convert(reflect.TypeOf(0.0)).Float() > 0
			}
		case eType.Kind() == reflect.Array, eType.Kind() == reflect.Slice,
			eType.Kind() == reflect.Ptr && eType.Elem().Kind() == reflect.Array:
			funcs[0] = func(idx int, value E, arr []E) bool {
				return reflect.ValueOf(value).Len() > 0
			}
		default:
			funcs[0] = func(idx int, value E, arr []E) bool {
				return true
			}
		}
	}
	return
}

// Any returns true if any of the elements in the given array are proved to be truthy according to the given functions.
//
// If "i" is the index of the currently iterated element then we will use the function located at index:
//
//	i % len(funcs)
//
// If true is returned by the function at this index before the last element has been evaluated, then All will return
// true early.
//
// If the given array is empty, then false is returned.
//
// If no functions are provided to Any, then depending on the type of the array, the following will occur, in order of
// priority (type signatures are given as regex to avoid verbosity):
//   - r"\[\]bool" (array of booleans): returns the value of the current boolean element.
//   - r"\[\]string" (array of strings): returns whether the value of the current string element != "".
//   - r"\[\]*?\[\].+" (array of arrays/pointers to arrays): returns whether the current nested array's length is > 0.
//   - r"\[\](u?int(8|16|32|64)?|float(32|64))" (array of numbers): returns true when the element is > 0.
//   - Otherwise, the only function will be one that always returns true.
func Any[E any](s []E, funcs ...func(idx int, value E, arr []E) bool) bool {
	if len(s) == 0 {
		return false
	}

	if len(funcs) == 0 {
		funcs = emptyFuncsResolve[E]()
	}

	for i, e := range s {
		if funcs[i%len(funcs)](i, e, s) {
			return true
		}
	}
	return false
}

// JoinedAny performs Any on an array created from the Join of the given arrays.
func JoinedAny[E any](funcs []func(idx int, value E, arr []E) bool, ss ...[]E) bool {
	return Any(Join(ss...), funcs...)
}

// All returns true if all the elements in the given array are proved to be truthy according to the given functions.
//
// If "i" is the index of the currently iterated element then we will use the function located at index:
//
//	i % len(funcs)
//
// If false is returned by the function at this index before the last element has been evaluated, then All will return
// false early.
//
// If the given array is empty, then false is returned.
//
// If no functions are provided to All, then depending on the type of the array, the following will occur, in order of
// priority (type signatures are given as regex to avoid verbosity):
//   - r"\[\]bool" (array of booleans): returns the value of the current boolean element.
//   - r"\[\]string" (array of strings): returns whether the value of the current string element != "".
//   - r"\[\]*?\[\].+" (array of arrays/pointers to arrays): returns whether the current nested array's length is > 0.
//   - r"\[\](u?int(8|16|32|64)?|float(32|64))" (array of numbers): returns true when the element is > 0.
//   - Otherwise, the only function will be one that always returns true.
func All[E any](s []E, funcs ...func(idx int, value E, arr []E) bool) bool {
	if len(s) == 0 {
		return false
	}

	if len(funcs) == 0 {
		funcs = emptyFuncsResolve[E]()
	}

	for i, e := range s {
		if !funcs[i%len(funcs)](i, e, s) {
			return false
		}
	}
	return true
}

// JoinedAll performs All on an array created from the Join of the given arrays.
func JoinedAll[E any](funcs []func(idx int, value E, arr []E) bool, ss ...[]E) bool {
	return All(Join(ss...), funcs...)
}

type twoSlices[E any] struct {
	values []reflect.Value
	s      []E
	less   func(a, b reflect.Value) bool
}

func (t twoSlices[E]) Len() int { return numbers.Min(len(t.s), len(t.values)) }

func (t twoSlices[E]) Less(i, j int) bool {
	if t.less == nil {
		return false
	}
	return t.less(t.values[i], t.values[j])
}

func (t twoSlices[E]) Swap(i, j int) {
	t.values[i], t.values[j] = t.values[j], t.values[i]
	t.s[i], t.s[j] = t.s[j], t.s[i]
}

func compoundLen(a reflect.Value) int {
	switch a.Kind() {
	case reflect.Array, reflect.Slice:
		return a.Len()
	case reflect.Struct:
		return a.NumField()
	default:
		return 0
	}
}

func compoundIndex(a reflect.Value, i int) reflect.Value {
	switch a.Kind() {
	case reflect.Array, reflect.Slice:
		return a.Index(i)
	case reflect.Struct:
		return a.Field(i)
	default:
		return a
	}
}

func orderCompound(a, b reflect.Value) misc.Ordered {
	aLen, bLen := compoundLen(a), compoundLen(b)
	for i := 0; i < numbers.Min(aLen, bLen); i++ {
		aEl, bEl := compoundIndex(a, i), compoundIndex(b, i)
		if aEl.Type() != bEl.Type() || aEl.Kind() != bEl.Kind() {
			continue
		}

		for {
			switch aEl.Kind() {
			case reflect.Ptr, reflect.Interface:
				aEl, bEl = aEl.Elem(), bEl.Elem()
				continue
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if aInt, bInt := aEl.Int(), bEl.Int(); aInt != bInt {
					return misc.Compare(aEl.Int(), bEl.Int())
				}
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				if aUint, bUint := aEl.Uint(), bEl.Uint(); aUint != bUint {
					return misc.Compare(aEl.Uint(), bEl.Uint())
				}
			case reflect.Float32, reflect.Float64:
				if aFloat, bFloat := aEl.Float(), bEl.Float(); aFloat != bFloat {
					return misc.Compare(aEl.Float(), bEl.Float())
				}
			case reflect.String:
				if aString, bString := aEl.String(), bEl.String(); aString != bString {
					return misc.Compare(aEl.String(), bEl.String())
				}
			case reflect.Array, reflect.Slice, reflect.Struct:
				if o := orderCompound(aEl, bEl); o != misc.Equal {
					return o
				}
			}
			break
		}
	}
	return misc.Compare(aLen, bLen)
}

func newTwoSlices[E any](s []E) twoSlices[E] {
	ts := twoSlices[E]{s: s}
	eType := reflect.TypeOf(new(E)).Elem()
	for {
		switch eType.Kind() {
		case reflect.Ptr, reflect.Interface:
			eType = eType.Elem()
			continue
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			ts.less = func(a, b reflect.Value) bool { return a.Int() < b.Int() }
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			ts.less = func(a, b reflect.Value) bool { return a.Uint() < b.Uint() }
		case reflect.Float32, reflect.Float64:
			ts.less = func(a, b reflect.Value) bool { return a.Float() < b.Float() }
		case reflect.String:
			ts.less = func(a, b reflect.Value) bool { return a.String() < b.String() }
		case reflect.Struct:
			ts.less = func(a, b reflect.Value) bool {
				return orderCompound(a, b) == misc.Less
			}
		case reflect.Array, reflect.Slice:
			ts.less = func(a, b reflect.Value) bool {
				return orderCompound(a, b) == misc.Less
			}
		default:
			ts.less = nil
		}
		break
	}

	if ts.less != nil {
		ts.values = Comprehension(s, func(idx int, value E, arr []E) reflect.Value {
			v := reflect.ValueOf(value)
			for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
				v = v.Elem()
			}
			return v
		})
	}
	return ts
}

// Order will order any slice of elements that can be ordered (integers, floats, and strings), as well as structs and
// arrays/slices, in place.
//
// It is useful for sorting arrays whose elements are constraints.Ordered. The reason why we accept a slice with any
// element type is for completeness: a slice whose elements are unordered will be left untouched/the order is preserved.
//
// Structs are ordered lexicographically. When given instances of a struct of type A and a struct of type B the fields
// of each struct are both iterated up until minimum of A.NumFields() and B.NumFields().
//   - If the fields' types are pointer/interface types then they're iteratively dereferenced until they are not.
//   - If the fields' types are not the same then they are skipped.
//   - If the fields' types are constraints.Ordered, and the values of them are found not to be equal: A field's value
//     is compared with B field's value and the result is returned. If they are equal then the next field is checked.
//     Fields that are Struct, Array, and Slice types are recursively compared. Try to avoid this, as this takes a lot
//     of reflection.
//   - If the fields' types are not constraints.Ordered, they are skipped.
//
// If iteration has finished without a return, then the number of fields for A and B are compared and returned.
//
// Arrays/slices are also ordered lexicographically in a similar way to Structs. Given an instance A of an array/slice
// type and an instance B of an array/slice type, the elements of each array/slice are iterated up until the minimum of
// A.Len() and B.Len().
//   - If the elements' types are pointer/interface types then they're iteratively dereferenced until they are not.
//   - If the elements' types are not the same then they are skipped
//   - If the elements' types are constraints.Ordered, and the values of them are found not to be equal: A element's
//     value is compared with B element's value and the result is returned. If they are equal, then the next element is
//     checked. Elements that are Struct, Array, and Slice types are recursively compared. Try to avoid this, as this
//     takes a lot of reflection.
//   - If the elements' types are not constraints.Ordered, they are skipped.
//
// If iteration has finished without a return, then the number of elements for A and B are compared and returned.
//
// Order incurs quite a lot more overhead than a call to sort.Slice or sort.Sort, as the entire array first needs to be
// converted to reflect.Value(s). Then sort.Sort will be called, sorting both the array of reflect.Value and the input
// slice at the same time.
func Order[E any](s []E) {
	ts := newTwoSlices(s)
	if ts.less != nil {
		sort.Sort(ts)
	}
}
