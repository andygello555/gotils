// Package numbers contains functions and constants to do with integers and generating sequences of integers.
package numbers

import (
	"golang.org/x/exp/constraints"
	"strconv"
)

const (
	MaxUint = ^uint(0)
	MinUint = 0
	MinInt  = -MaxInt - 1
	MaxInt  = int(MaxUint >> 1)
)

// Number is the set of both SignedNumber and UnsignedNumber types.
type Number interface {
	SignedNumber | UnsignedNumber
}

// SignedNumber is the set of all floating point, and signed integer types.
type SignedNumber interface {
	constraints.Float | constraints.Signed
}

// UnsignedNumber is the set of all unsigned integer types.
type UnsignedNumber interface {
	constraints.Unsigned
}

// ScaleRange scales a value x, that is between xMin and xMax, to be between yMin and yMax.
func ScaleRange[N Number](x, xMin, xMax, yMin, yMax N) N {
	return (x-xMin)/(xMax-xMin)*(yMax-yMin) + yMin
}

// Clamp clamps the given Number at the given maximum value.
func Clamp[N Number](x, max N) N {
	if x > max {
		x = max
	}
	return x
}

// ClampMin clamps the given Number at the given minimum value.
func ClampMin[N Number](x, min N) N {
	if x < min {
		x = min
	}
	return x
}

// ClampMinMax clamps the given Number at the given minimum and maximum values.
func ClampMinMax[N Number](x, min, max N) N {
	if x < min {
		x = min
	} else if x > max {
		x = max
	}
	return x
}

// Abs returns the absolute value of a Number.
func Abs[N SignedNumber](n N) N {
	if n < N(0) {
		n = n * N(-1)
	}
	return n
}

// Range generates a SignedNumber array with indices from start to end with the given step value.
//
// Returns an empty array if step is equal to 0 or end is less than start and step is a positive SignedNumber.
func Range[N SignedNumber](start, end, step N) []N {
	if step == 0 || (end < start && step > 0) {
		return []N{}
	}

	// For checking if the predicate still holds.
	keepGoing := func(s, e N) bool {
		if step < 0 {
			return e <= s
		} else {
			return s <= e
		}
	}

	s := make([]N, 0, int(N(1)+Abs(end-start)/Abs(step)))
	for keepGoing(start, end) {
		s = append(s, start)
		start += step
	}
	return s
}

// Max returns the maximum of all the given Number(s).
//
// If no Number(s) are given, 0 is returned.
func Max[N Number](numbers ...N) (max N) {
	if len(numbers) > 0 {
		max = numbers[0]
		for _, n := range numbers {
			if n > max {
				max = n
			}
		}
	}
	return max
}

// Min returns the minimum of all the given Number(s).
//
// If no Number(s) are given, 0 is returned.
func Min[N Number](numbers ...N) (min N) {
	if len(numbers) > 0 {
		min = numbers[0]
		for _, n := range numbers {
			if n < min {
				min = n
			}
		}
	}
	return min
}

// Ordinal gives you the input integer with a rank/ordinal format.
//
// Ordinal(3) -> "3rd". Straight from the "go-humanize" library: https://github.com/dustin/go-humanize.
func Ordinal(x int) string {
	return strconv.Itoa(x) + OrdinalOnly(x)
}

// OrdinalOnly gives you the input integer's rank/ordinal.
//
// OrdinalOnly(3) -> "rd".
func OrdinalOnly(x int) string {
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
	return suffix
}
