package structs

import (
	"container/heap"
	"fmt"
)

// How to create and use a Heap.
//
// Just uses the usual standard heap package functions.
func ExampleHeap() {
	// Create a Heap.
	stringHeap := make(Heap[string], 0)
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
