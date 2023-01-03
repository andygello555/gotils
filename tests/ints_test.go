package tests

import (
	"github.com/andygello555/gotils/v2/ints"
	"reflect"
	"testing"
)

func TestRange(t *testing.T) {
	for _, test := range []struct {
		start          int
		end            int
		step           int
		expectedOutput []int
	}{
		{
			0,
			10,
			2,
			[]int{0, 2, 4, 6, 8, 10},
		},
		{
			0,
			10,
			3,
			[]int{0, 3, 6, 9},
		},
		{
			10,
			0,
			2,
			[]int{},
		},
		{
			10,
			0,
			-2,
			[]int{10, 8, 6, 4, 2, 0},
		},
		{
			10,
			0,
			-3,
			[]int{10, 7, 4, 1},
		},
		{
			1,
			0,
			-2,
			[]int{1},
		},
		{
			0,
			10,
			0,
			[]int{},
		},
		{
			10,
			0,
			1,
			[]int{},
		},
		{
			0,
			10,
			1,
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	} {
		r := ints.Range(test.start, test.end, test.step)
		if !reflect.DeepEqual(r, test.expectedOutput) {
			t.Errorf("Start = \"%v\", End = \"%v\", Step = \"%v\"\nExpected range: \"%v\"\nGot: \"%v\"", test.start, test.end, test.step, test.expectedOutput, r)
		}
	}
}

func TestMax(t *testing.T) {
	for _, test := range []struct {
		input          []int
		expectedOutput int
	}{
		{
			[]int{213, 1, 0, 53, 123, 3, 999, 1000, 12, 5, 9, 8, 123, 3, 1000},
			1000,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			10,
		},
		{
			[]int{0},
			0,
		},
		{
			[]int{},
			0,
		},
		{
			[]int{100, 100},
			100,
		},
	} {
		max := ints.Max(test.input...)
		if max != test.expectedOutput {
			t.Errorf("Got: \"%v\", expected: \"%v\"", max, test.expectedOutput)
		}
	}
}

func TestMin(t *testing.T) {
	for _, test := range []struct {
		input          []int
		expectedOutput int
	}{
		{
			[]int{213, 1, 0, 53, 123, 3, 999, 1000, 12, 5, 9, 8, 123, 3, 1000},
			0,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			1,
		},
		{
			[]int{0},
			0,
		},
		{
			[]int{},
			0,
		},
		{
			[]int{100, 100},
			100,
		},
	} {
		max := ints.Min(test.input...)
		if max != test.expectedOutput {
			t.Errorf("Got: \"%v\", expected: \"%v\"", max, test.expectedOutput)
		}
	}
}
