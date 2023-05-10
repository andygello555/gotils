package maps

import (
	"fmt"
)

// Deep copying a map using recursion via CopyMap.
//
// CopyMap was designed for use with deserialised JSONs hence the map[string]any signature.
func ExampleCopyMap() {
	original := map[string]any{
		"hello": "world",
		"age":   20,
		"bald":  false,
		"friends": []any{
			"Bob",
			"Jane",
			"John",
			"Mark",
			map[string]any{
				"name": "Gregor",
				"age":  31,
				"friends": []any{
					"Bill",
					"Bob",
					"Sarah",
				},
			},
		},
	}
	// Clone the above map.
	clone := CopyMap(original)

	fmt.Println("Original:", original)
	fmt.Println("Clone:", clone)
	// Output:
	// Original: map[age:20 bald:false friends:[Bob Jane John Mark map[age:31 friends:[Bill Bob Sarah] name:Gregor]] hello:world]
	// Clone: map[age:20 bald:false friends:[Bob Jane John Mark map[age:31 friends:[Bill Bob Sarah] name:Gregor]] hello:world]
}

// Retrieve the keys from a map.
func ExampleKeys() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(Keys(m))
	// Output:
	// [a b c]
}

// Retrieve the values from a map.
func ExampleValues() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(Values(m))
	// Output:
	// [1 2 3]
}

// Retrieve the keys from the map in order.
func ExampleOrderedKeys() {
	m := map[string]int{
		"c": 3,
		"b": 2,
		"a": 1,
	}
	fmt.Println(OrderedKeys(m))
	// Output:
	// [a b c]
}

// Range over a map and display the index, key, and value of the current key-value pair.
func ExampleRangeKeys() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	RangeKeys(m, func(i int, key string, val int) bool {
		fmt.Printf("%d-%s-%d\n", i, key, val)
		return true
	})
	// Output:
	// 0-a-1
	// 1-b-2
	// 2-c-3
}

// Range over a map, in key order, and display the index, key, and value of the current key-value pair.
func ExampleRangeOrderedKeys() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	RangeOrderedKeys(m, func(i int, key string, val int) bool {
		fmt.Printf("%d-%s-%d\n", i, key, val)
		return true
	})
	// Output:
	// 0-a-1
	// 1-b-2
	// 2-c-3
}
