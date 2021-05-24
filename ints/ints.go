package ints

const (
	MaxUint = ^uint(0)
	MinUint = 0
	MinInt = -MaxInt - 1
	MaxInt = int(MaxUint >> 1)
)

// Generates an integer array with indices from start to end with the given step value.
//
// Returns an empty array if step is equal to 0 or end is less than start and step is a positive number.
func Range(start, end, step int) []int {
	if step == 0 || (end < start && step > 0) {
		return []int{}
	}

	// Gets the absolute of an integer.
	abs := func(x int) int {
		if x < 0 {
			return x * -1
		}
		return x
	}

	// For checking if the iteration still holds.
	keepGoing := func(s, e int) bool {
		if step < 0 {
			return e <= s
		} else {
			return s <= e
		}
	}

	s := make([]int, 0, 1+abs(end-start)/abs(step))
	for keepGoing(start, end) {
		s = append(s, start)
		start += step
	}
	return s
}

// Returns the maximum of all the given integers.
func Max(numbers... int) (max int) {
	max = MinInt
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}
	return max
}

// Returns the minimum of all the given integers.
func Min(numbers... int) (min int) {
	min = MaxInt
	for _, n := range numbers {
		if n < min {
			min = n
		}
	}
	return min
}
