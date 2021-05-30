package strings

import (
	"container/heap"
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

// How to create and use a StringHeap.
//
// Just uses the usual standard heap package functions.
func ExampleStringHeap() {
	// Create a StringHeap.
	stringHeap := make(StringHeap, 0)
	heap.Init(&stringHeap)

	// Push some strings. Make sure to use heap.Push rather than stringHeap.Push.
	heap.Push(&stringHeap, "Crisps")
	heap.Push(&stringHeap, "Egg")
	heap.Push(&stringHeap, "Bananas")
	heap.Push(&stringHeap, "Doughnut")
	heap.Push(&stringHeap, "Apple")
	heap.Push(&stringHeap, "Fried chicken")
	heap.Push(&stringHeap, "Orange")
	heap.Push(&stringHeap, "Grapefruit")

	// We can get the length using stringHeap.Len.
	fmt.Println("Length before:", stringHeap.Len())

	// Pop them off.
	for stringHeap.Len() > 0 {
		fmt.Println(heap.Pop(&stringHeap).(string))
	}

	fmt.Println("Length after:", stringHeap.Len())
	// Output:
	// Length before: 8
	// Apple
	// Bananas
	// Crisps
	// Doughnut
	// Egg
	// Fried chicken
	// Grapefruit
	// Orange
	// Length after: 0
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
