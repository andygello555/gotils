package strings

import (
	"fmt"
)

// Will replace the given character indices within the given string with the given new strings in the order given.
func ExampleReplaceCharIndex() {
	old := "Hello world! Hope you are having a good day today."
	fmt.Println(ReplaceCharIndex(old, []int{1, 2, 7, 9, 10, 14, 20, 41}, "Fizz", "Buzz", "Bing"))
	// Output:
	// HFizzBuzzlo wBingrFizzBuzz! HBingpe yoFizz are having a good dBuzzy today.
}

// Will replace "world" with "me", "you" with "I" and "are" with "am".
func ExampleReplaceCharIndexRange() {
	old := "Hello world! Hope you are having a good day today."
	fmt.Println(ReplaceCharIndexRange(old, [][]int{{6, 11}, {18, 21}, {22, 25}}, "me", "I", "am"))
	// Output:
	// Hello me! Hope I am having a good day today.
}

// Get the type name of a string then nil.
func ExampleTypeName() {
	val := "hello world"
	fmt.Printf("The type of \"%v\" is \"%s\"\n", val, TypeName(val))
	fmt.Printf("Type type of \"%v\" is \"%s\"\n", nil, TypeName(nil))
	// Output:
	// The type of "hello world" is "string"
	// Type type of "<nil>" is "<nil>"
}

// Checks whether "abc" and "abc?" contain letters only.
func ExampleIsAlpha() {
	fmt.Println(IsAlpha("abc"))
	fmt.Println(IsAlpha("abc?"))
	// Output:
	// true
	// false
}

// Checks whether "123" and "123abc" contain numbers only.
func ExampleIsNumeric() {
	fmt.Println(IsNumeric("123"))
	fmt.Println(IsNumeric("123abc"))
	// Output:
	// true
	// false
}

// Checks whether "abc123" and "abc123?!" are alphanumeric strings.
func ExampleIsAlphaNumeric() {
	fmt.Println(IsAlphaNumeric("abc123"))
	fmt.Println(IsAlphaNumeric("abc123?!"))
	// Output:
	// true
	// false
}

// Splits a string which is in camelcase at each hump.
func ExampleSplitCamelcase() {
	fmt.Println(SplitCamelcase("HelloWorld"))
	// Output: [Hello World]
}

// Joins each hump of the camelcase string with the given separator.
func ExampleJoinCamelcase() {
	fmt.Println(JoinCamelcase("HelloWorld", ", "))
	// Output: Hello, World
}
