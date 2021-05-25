package maps

import "fmt"

// Deep copying a map using recursion via CopyMap.
//
// CopyMap was designed for use with deserialised JSONs hence the map[string]interface{} signature.
func ExampleCopyMap() {
	original := map[string]interface{} {
		"hello": "world",
		"age": 20,
		"bald": false,
		"friends": []interface{} {
			"Bob",
			"Jane",
			"John",
			"Mark",
			map[string]interface{} {
				"name": "Gregor",
				"age": 31,
				"friends": []interface{} {
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
