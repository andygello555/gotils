// Package maps contains functions related to maps such as deep copying and equality testing.
package maps

import (
	"container/heap"
	"fmt"
	"github.com/andygello555/gotils/v2/structs"
	"github.com/go-test/deep"
	"golang.org/x/exp/constraints"
	"strings"
	"testing"
)

// CopyMap clones a map deeply using recursion.
func CopyMap(m map[string]any) map[string]any {
	cp := make(map[string]any)
	for k, v := range m {
		vm, ok := v.(map[string]any)
		if ok {
			cp[k] = CopyMap(vm)
		} else {
			cp[k] = v
		}
	}

	return cp
}

// MapRangeFunc is the signature passed to RangeKeys, and RangeOrderedKeys. It is passed the key-value pair, and
// should return whether you want to keep iterating over the map.
type MapRangeFunc[K comparable, V any] func(i int, key K, val V) bool

// RangeOrderedKeys calls the given MapRangeFunc on each index-key-value triple. Triples are ordered by their keys.
func RangeOrderedKeys[K constraints.Ordered, V any](m map[K]V, fun MapRangeFunc[K, V]) {
	keys := make(structs.Heap[K], 0)
	heap.Init(&keys)
	for key := range m {
		heap.Push(&keys, key)
	}

	i := 0
	for keys.Len() > 0 {
		key := heap.Pop(&keys).(K)
		val := m[key]
		if !fun(i, key, val) {
			break
		}
		i++
	}
}

// RangeKeys calls the given MapRangeFunc on each index-key-value triple. Triples are unordered.
func RangeKeys[K comparable, V any](m map[K]V, fun MapRangeFunc[K, V]) {
	i := 0
	for key, val := range m {
		if !fun(i, key, val) {
			break
		}
		i++
	}
}

// OrderedKeys returns the ordered keys for a given map.
func OrderedKeys[K constraints.Ordered, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	RangeOrderedKeys(m, func(i int, key K, val V) bool {
		keys[i] = key
		return true
	})
	return keys
}

// Keys returns the keys within a given map.
func Keys[K comparable, V any](m map[K]V) []K {
	i := 0
	keys := make([]K, len(m))
	for key := range m {
		keys[i] = key
		i++
	}
	return keys
}

// Values returns the values within a given map.
func Values[K comparable, V any](m map[K]V) []V {
	i := 0
	values := make([]V, len(m))
	for _, val := range m {
		values[i] = val
		i++
	}
	return values
}

// Filter takes a map and runs the given predicate function on each index-key-value triple. If the predicate
// returns false for an element, then that element will be removed from the given map.
func Filter[K comparable, V any](m map[K]V, fun func(i int, key K, val V) bool) {
	i := 0
	for key, val := range m {
		if !fun(i, key, val) {
			delete(m, key)
		}
	}
}

// Union takes merges the source map into the destination map, overriding any matching keys.
func Union[K comparable, V any](dst map[K]V, src map[K]V) {
	for key, val := range src {
		dst[key] = val
	}
}

// Difference removes every key-value pair in m that also exists in n.
func Difference[K comparable, V any](m map[K]V, n map[K]V) {
	for key := range n {
		delete(m, key)
	}
}

// JsonMapEqualTest used in tests to check equality between two anys.
//
// This takes into account orderings of slices.
func JsonMapEqualTest(t *testing.T, actual, expected any, forString string) {
	if diff := deep.Equal(actual, expected); diff != nil {
		var errB strings.Builder
		errB.WriteString(fmt.Sprintf("Difference between actual and expected for %s (Left = Actual, Right = Expected)\n", forString))
		for _, d := range diff {
			errB.WriteString(fmt.Sprintf("\t%s\n", d))
		}
		t.Error(errB.String())
	}
}
