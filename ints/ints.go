// Package ints contains functions and constants to do with integers and generating sequences of integers.
package ints

import "strconv"

const (
	MaxUint = ^uint(0)
	MinUint = 0
	MinInt = -MaxInt - 1
	MaxInt = int(MaxUint >> 1)
)

// Range generates an integer array with indices from start to end with the given step value.
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

// Max returns the maximum of all the given integers.
//
// If no numbers are given 0 is returned.
func Max(numbers... int) (max int) {
	if len(numbers) > 0 {
		max = MinInt
		for _, n := range numbers {
			if n > max {
				max = n
			}
		}
	}
	return max
}

// Min returns the minimum of all the given integers.
//
// If no numbers are given 0 is returned.
func Min(numbers... int) (min int) {
	if len(numbers) > 0 {
		min = MaxInt
		for _, n := range numbers {
			if n < min {
				min = n
			}
		}
	}
	return min
}

// Ordinal gives you the input number in a rank/ordinal format.
//
// Ordinal(3) -> 3rd. Straight from the "go-humanize" library: https://github.com/dustin/go-humanize.
func Ordinal(x int) string {
	suffix := "th"
	switch x % 10 {
	case 1:
		if x%100 != 11 {
			suffix = "st"
		}
	case 2:
		if x%100 != 12 {
			suffix = "nd"
		}
	case 3:
		if x%100 != 13 {
			suffix = "rd"
		}
	}
	return strconv.Itoa(x) + suffix
}
