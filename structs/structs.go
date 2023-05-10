package structs

import "golang.org/x/exp/constraints"

// Heap is a priority queue which sorts strings lexicographically.
type Heap[E constraints.Ordered] []E

// Len gives the length of the Heap.
func (h Heap[E]) Len() int { return len(h) }

// Less returns true if the element at the first index is less than the element at the second.
func (h Heap[E]) Less(i, j int) bool {
	return h[i] < h[j]
}

// Swap swaps the two elements indicated via the given indices.
func (h Heap[E]) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push pushes the given element onto the heap.
func (h *Heap[E]) Push(x any) { *h = append(*h, x.(E)) }

// Pop pops the tail of the queue.
func (h *Heap[E]) Pop() any {
	old := *h
	n := len(old)
	str := old[n-1]
	*h = old[0 : n-1]
	return str
}
