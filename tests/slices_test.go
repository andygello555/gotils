package tests

import (
	"github.com/andygello555/gotils/v2/slices"
	"reflect"
	"testing"
)

func TestRemoveDuplicatesAndSort(t *testing.T) {
	for _, test := range []struct {
		input          []int
		expectedOutput []int
	}{
		{
			[]int{0, 0, 1, 4, 6, 2, 3, 3, 7, 5},
			[]int{0, 1, 2, 3, 4, 5, 6, 7},
		},
		{
			[]int{1238, 7216846, 28, 28, 23, 1238, 983, 847, 23, 983},
			[]int{23, 28, 847, 983, 1238, 7216846},
		},
		{
			[]int{1, 12, 123, 1234, 12, 21, 34, 65, 79, 100},
			[]int{1, 12, 21, 34, 65, 79, 100, 123, 1234},
		},
		{
			[]int{5, 6, 10, 0, 1, 1, 2, 6, 5, 9, 10, 3, 3, 4, 7, 8},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	} {
		slices.RemoveDuplicatesAndSort(&test.input)
		if !reflect.DeepEqual(test.input, test.expectedOutput) {
			t.Errorf("Got: \"%v\", expected: \"%v\"", test.input, test.expectedOutput)
		}
	}
}

func TestAddElems(t *testing.T) {
	for _, test := range []struct {
		slice          []any
		value          []any
		indices        []int
		expectedOutput []any
	}{
		{
			[]any{"hello", "world", "insert", "here", "->", "<-", "nice"},
			[]any{": )"},
			[]int{5},
			[]any{"hello", "world", "insert", "here", "->", ": )", "<-", "nice"},
		},
		{
			[]any{1, 2, 3},
			[]any{0},
			[]int{0, 0, 3, 7, 1},
			[]any{0, 0, 1, 0, 2, 3, nil, 0},
		},
		{
			[]any{
				map[string]any{
					"name": "Bob",
					"age":  29,
					"egg":  true,
				},
				[]any{
					"Jill",
					"Bill",
					"Tom",
				},
				"Sarah",
				"John",
			},
			[]any{map[string]string{
				"hello": "world",
			}},
			[]int{0, 1, 3, 9, 3, 1},
			[]any{
				map[string]string{
					"hello": "world",
				},
				map[string]string{
					"hello": "world",
				},
				map[string]any{
					"name": "Bob",
					"age":  29,
					"egg":  true,
				},
				map[string]string{
					"hello": "world",
				},
				[]any{
					"Jill",
					"Bill",
					"Tom",
				},
				"Sarah",
				"John",
				nil,
				nil,
				map[string]string{
					"hello": "world",
				},
			},
		},
		{
			[]any{1, 2, 3},
			[]any{4},
			[]int{3},
			[]any{1, 2, 3, 4},
		},
		{
			[]any{1, 2, 3},
			[]any{},
			[]int{1},
			[]any{1, nil, 2, 3},
		},
		{
			[]any{1, 2, 3},
			[]any{4, 5, 6, 7},
			[]int{3, 4, 5},
			[]any{1, 2, 3, 4, 5, 6},
		},
		{
			[]any{1, 2, 3},
			[]any{4, 5, 6},
			[]int{},
			[]any{1, 2, 3},
		},
		{
			[]any{},
			[]any{1, 2, 3},
			[]int{0, 1, 2},
			[]any{1, 2, 3},
		},
		{
			[]any{},
			[]any{},
			[]int{},
			[]any{},
		},
		{
			[]any{},
			[]any{},
			[]int{0, 1, 2},
			[]any{nil, nil, nil},
		},
	} {
		newSlice := slices.AddElems(test.slice, test.value, test.indices...)
		if !reflect.DeepEqual(newSlice, test.expectedOutput) {
			t.Errorf("Got: \"%v\", expected: \"%v\"", newSlice, test.expectedOutput)
		}
	}
}

func TestRemoveElems(t *testing.T) {
	for _, test := range []struct {
		slice          []any
		indices        []int
		expectedOutput []any
	}{
		{
			[]any{"hello", "world", "delete", "->", "this", "<-", "nice"},
			[]int{4},
			[]any{"hello", "world", "delete", "->", "<-", "nice"},
		},
		{
			[]any{1, 2, 3},
			[]int{0, 0, 2},
			[]any{2},
		},
		{
			[]any{
				map[string]any{
					"name": "Bob",
					"age":  29,
					"egg":  true,
				},
				[]any{
					"Jill",
					"Bill",
					"Tom",
				},
				"Sarah",
				"John",
			},
			[]int{1, 3, 3, 1},
			[]any{
				map[string]any{
					"name": "Bob",
					"age":  29,
					"egg":  true,
				},
				"Sarah",
			},
		},
		{
			[]any{1, 2, 3},
			[]int{2},
			[]any{1, 2},
		},
		{
			[]any{1, 2, 3},
			[]int{},
			[]any{1, 2, 3},
		},
		{
			[]any{},
			[]int{1, 2, 3},
			[]any{},
		},
	} {
		newSlice := slices.RemoveElems(test.slice, test.indices...)
		if !reflect.DeepEqual(newSlice, test.expectedOutput) {
			t.Errorf("Got: \"%v\", expected: \"%v\"", newSlice, test.expectedOutput)
		}
	}
}

func TestSameElements(t *testing.T) {
	for _, test := range []struct {
		slice1         []any
		slice2         []any
		expectedOutput bool
	}{
		{
			[]any{1, 2, 3},
			[]any{2, 1, 3},
			true,
		},
		{
			[]any{1, 2, 3},
			[]any{1, 2, 3},
			true,
		},
		{
			[]any{
				map[string]any{
					"name": "Jim",
					"age":  20,
				},
				map[string]any{
					"name": "Bob",
					"age":  38,
				},
			},
			[]any{
				map[string]any{
					"name": "Bob",
					"age":  38,
				},
				map[string]any{
					"name": "Jim",
					"age":  20,
				},
			},
			true,
		},
		{
			[]any{
				map[string]any{
					"name": "Jim",
					"age":  20,
				},
				map[string]any{
					"name": "Bob",
					"age":  38,
				},
			},
			[]any{
				map[string]any{
					"name": "Bob",
					"age":  38,
				},
				// Age is not 20
				map[string]any{
					"name": "Jim",
					"age":  21,
				},
			},
			false,
		},
	} {
		actual := slices.SameElements(test.slice1, test.slice2)
		if actual != test.expectedOutput {
			t.Errorf("Got: \"%v\", expected: \"%v\"", actual, test.expectedOutput)
		}
	}
}

func TestJoin(t *testing.T) {
	for testNo, test := range []struct {
		arrays   [][]any
		expected []any
	}{
		{
			arrays:   [][]any{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expected: []any{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			arrays:   [][]any{},
			expected: []any{},
		},
		{
			arrays:   [][]any{{1, 2, 3}},
			expected: []any{1, 2, 3},
		},
	} {
		actual := slices.Join(test.arrays...)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("%d: Got %v, expected %v", testNo+1, actual, test.expected)
		}
	}
}

func TestAny(t *testing.T) {
	for testNo, test := range []struct {
		array    []bool
		funcs    []func(idx int, value bool, arr []bool) bool
		expected bool
	}{
		{
			array:    []bool{false, true, false},
			funcs:    []func(idx int, value bool, arr []bool) bool{},
			expected: true,
		},
		{
			array:    []bool{false, false, false},
			funcs:    []func(idx int, value bool, arr []bool) bool{},
			expected: false,
		},
		{
			array:    []bool{},
			funcs:    []func(idx int, value bool, arr []bool) bool{},
			expected: false,
		},
		{
			array: []bool{true, true, true},
			funcs: []func(idx int, value bool, arr []bool) bool{
				func(idx int, value bool, arr []bool) bool {
					return !value
				},
			},
			expected: false,
		},
		{
			array: []bool{true, false, true, false},
			funcs: []func(idx int, value bool, arr []bool) bool{
				func(idx int, value bool, arr []bool) bool {
					return !value
				},
				func(idx int, value bool, arr []bool) bool {
					return value
				},
			},
			expected: false,
		},
	} {
		actual := slices.Any(test.array, test.funcs...)
		if actual != test.expected {
			t.Errorf("[]bool no. %d (%v): Got %t, expected %t", testNo+1, test, actual, test.expected)
		}
	}

	for testNo, test := range []struct {
		array    []float64
		funcs    []func(idx int, value float64, arr []float64) bool
		expected bool
	}{
		{
			array:    []float64{0.0, 1.0, 0.0},
			funcs:    []func(idx int, value float64, arr []float64) bool{},
			expected: true,
		},
		{
			array:    []float64{0.0, 0.0, 0.0},
			funcs:    []func(idx int, value float64, arr []float64) bool{},
			expected: false,
		},
		{
			array:    []float64{},
			funcs:    []func(idx int, value float64, arr []float64) bool{},
			expected: false,
		},
		{
			array: []float64{1.0, 2.0, 3.0},
			funcs: []func(idx int, value float64, arr []float64) bool{
				func(idx int, value float64, arr []float64) bool {
					return value == 0
				},
			},
			expected: false,
		},
		{
			array: []float64{1.0, 0.0, 2.0, 0.0},
			funcs: []func(idx int, value float64, arr []float64) bool{
				func(idx int, value float64, arr []float64) bool {
					return value == 0
				},
				func(idx int, value float64, arr []float64) bool {
					return value > 0
				},
			},
			expected: false,
		},
	} {
		actual := slices.Any(test.array, test.funcs...)
		if actual != test.expected {
			t.Errorf("[]float64 no. %d (%v): Got %t, expected %t", testNo+1, test, actual, test.expected)
		}
	}

	for testNo, test := range []struct {
		array    []string
		funcs    []func(idx int, value string, arr []string) bool
		expected bool
	}{
		{
			array:    []string{"", "+", ""},
			funcs:    []func(idx int, value string, arr []string) bool{},
			expected: true,
		},
		{
			array:    []string{"", "", ""},
			funcs:    []func(idx int, value string, arr []string) bool{},
			expected: false,
		},
		{
			array:    []string{},
			funcs:    []func(idx int, value string, arr []string) bool{},
			expected: false,
		},
		{
			array: []string{"1", "2", "3"},
			funcs: []func(idx int, value string, arr []string) bool{
				func(idx int, value string, arr []string) bool {
					return value == ""
				},
			},
			expected: false,
		},
		{
			array: []string{"1", "", "2", ""},
			funcs: []func(idx int, value string, arr []string) bool{
				func(idx int, value string, arr []string) bool {
					return value == ""
				},
				func(idx int, value string, arr []string) bool {
					return value != ""
				},
			},
			expected: false,
		},
	} {
		actual := slices.Any(test.array, test.funcs...)
		if actual != test.expected {
			t.Errorf("[]string no. %d (%v): Got %t, expected %t", testNo+1, test, actual, test.expected)
		}
	}

	for testNo, test := range []struct {
		array    []*[]any
		funcs    []func(idx int, value *[]any, arr []*[]any) bool
		expected bool
	}{
		{
			array:    []*[]any{{}, {1}, {}},
			funcs:    []func(idx int, value *[]any, arr []*[]any) bool{},
			expected: true,
		},
		{
			array: []*[]any{{}, {}, {}},
			funcs: []func(idx int, value *[]any, arr []*[]any) bool{},
			// This returns true because the empty arrays above are actually slices
			expected: true,
		},
		{
			array:    []*[]any{},
			funcs:    []func(idx int, value *[]any, arr []*[]any) bool{},
			expected: false,
		},
		{
			array: []*[]any{{1}, {2}, {3}},
			funcs: []func(idx int, value *[]any, arr []*[]any) bool{
				func(idx int, value *[]any, arr []*[]any) bool {
					return len(*value) == 0
				},
			},
			expected: false,
		},
		{
			array: []*[]any{{1}, {}, {2}, {}},
			funcs: []func(idx int, value *[]any, arr []*[]any) bool{
				func(idx int, value *[]any, arr []*[]any) bool {
					return len(*value) == 0
				},
				func(idx int, value *[]any, arr []*[]any) bool {
					return len(*value) > 0
				},
			},
			expected: false,
		},
	} {
		actual := slices.Any(test.array, test.funcs...)
		if actual != test.expected {
			t.Errorf("[]*[]any no. %d (%v): Got %t, expected %t", testNo+1, test, actual, test.expected)
		}
	}

	for testNo, test := range []struct {
		array    []float64
		funcs    []func(idx int, value float64, arr []float64) bool
		expected bool
	}{
		{
			array:    []float64{0.0, 1.0, 0.0},
			funcs:    []func(idx int, value float64, arr []float64) bool{},
			expected: true,
		},
		{
			array:    []float64{0.0, 0.0, 0.0},
			funcs:    []func(idx int, value float64, arr []float64) bool{},
			expected: false,
		},
		{
			array:    []float64{},
			funcs:    []func(idx int, value float64, arr []float64) bool{},
			expected: false,
		},
		{
			array: []float64{1.0, 2.0, 3.0},
			funcs: []func(idx int, value float64, arr []float64) bool{
				func(idx int, value float64, arr []float64) bool {
					return value == 0
				},
			},
			expected: false,
		},
		{
			array: []float64{1.0, 0.0, 2.0, 0.0},
			funcs: []func(idx int, value float64, arr []float64) bool{
				func(idx int, value float64, arr []float64) bool {
					return value == 0
				},
				func(idx int, value float64, arr []float64) bool {
					return value > 0
				},
			},
			expected: false,
		},
	} {
		actual := slices.Any(test.array, test.funcs...)
		if actual != test.expected {
			t.Errorf("[]float64 no. %d (%v): Got %t, expected %t", testNo+1, test, actual, test.expected)
		}
	}
}

func TestAll(t *testing.T) {
	for testNo, test := range []struct {
		array    []bool
		funcs    []func(idx int, value bool, arr []bool) bool
		expected bool
	}{
		{
			array:    []bool{true, true, true},
			funcs:    []func(idx int, value bool, arr []bool) bool{},
			expected: true,
		},
		{
			array:    []bool{false, false, false},
			funcs:    []func(idx int, value bool, arr []bool) bool{},
			expected: false,
		},
		{
			array:    []bool{},
			funcs:    []func(idx int, value bool, arr []bool) bool{},
			expected: false,
		},
		{
			array: []bool{true, true, true},
			funcs: []func(idx int, value bool, arr []bool) bool{
				func(idx int, value bool, arr []bool) bool {
					return !value
				},
			},
			expected: false,
		},
		{
			array: []bool{true, false, true, false},
			funcs: []func(idx int, value bool, arr []bool) bool{
				func(idx int, value bool, arr []bool) bool {
					return !value
				},
				func(idx int, value bool, arr []bool) bool {
					return value
				},
			},
			expected: false,
		},
	} {
		actual := slices.All(test.array, test.funcs...)
		if actual != test.expected {
			t.Errorf("[]bool no. %d (%v): Got %t, expected %t", testNo+1, test, actual, test.expected)
		}
	}

	for testNo, test := range []struct {
		array    []float64
		funcs    []func(idx int, value float64, arr []float64) bool
		expected bool
	}{
		{
			array:    []float64{1.0, 1.0, 1.0},
			funcs:    []func(idx int, value float64, arr []float64) bool{},
			expected: true,
		},
		{
			array:    []float64{0.0, 0.0, 0.0},
			funcs:    []func(idx int, value float64, arr []float64) bool{},
			expected: false,
		},
		{
			array:    []float64{},
			funcs:    []func(idx int, value float64, arr []float64) bool{},
			expected: false,
		},
		{
			array: []float64{1.0, 2.0, 3.0},
			funcs: []func(idx int, value float64, arr []float64) bool{
				func(idx int, value float64, arr []float64) bool {
					return value == 0
				},
			},
			expected: false,
		},
		{
			array: []float64{1.0, 0.0, 2.0, 0.0},
			funcs: []func(idx int, value float64, arr []float64) bool{
				func(idx int, value float64, arr []float64) bool {
					return value == 0
				},
				func(idx int, value float64, arr []float64) bool {
					return value > 0
				},
			},
			expected: false,
		},
	} {
		actual := slices.All(test.array, test.funcs...)
		if actual != test.expected {
			t.Errorf("[]float64 no. %d (%v): Got %t, expected %t", testNo+1, test, actual, test.expected)
		}
	}

	for testNo, test := range []struct {
		array    []string
		funcs    []func(idx int, value string, arr []string) bool
		expected bool
	}{
		{
			array:    []string{"1", "2", "3"},
			funcs:    []func(idx int, value string, arr []string) bool{},
			expected: true,
		},
		{
			array:    []string{"", "", ""},
			funcs:    []func(idx int, value string, arr []string) bool{},
			expected: false,
		},
		{
			array:    []string{},
			funcs:    []func(idx int, value string, arr []string) bool{},
			expected: false,
		},
		{
			array: []string{"1", "2", "3"},
			funcs: []func(idx int, value string, arr []string) bool{
				func(idx int, value string, arr []string) bool {
					return value == ""
				},
			},
			expected: false,
		},
		{
			array: []string{"1", "", "2", ""},
			funcs: []func(idx int, value string, arr []string) bool{
				func(idx int, value string, arr []string) bool {
					return value == ""
				},
				func(idx int, value string, arr []string) bool {
					return value != ""
				},
			},
			expected: false,
		},
	} {
		actual := slices.All(test.array, test.funcs...)
		if actual != test.expected {
			t.Errorf("[]string no. %d (%v): Got %t, expected %t", testNo+1, test, actual, test.expected)
		}
	}

	for testNo, test := range []struct {
		array    []*[]any
		funcs    []func(idx int, value *[]any, arr []*[]any) bool
		expected bool
	}{
		{
			array:    []*[]any{{1}, {2}, {3}},
			funcs:    []func(idx int, value *[]any, arr []*[]any) bool{},
			expected: true,
		},
		{
			array: []*[]any{{}, {}, {}},
			funcs: []func(idx int, value *[]any, arr []*[]any) bool{},
			// This returns true because the empty arrays above are actually slices
			expected: true,
		},
		{
			array:    []*[]any{},
			funcs:    []func(idx int, value *[]any, arr []*[]any) bool{},
			expected: false,
		},
		{
			array: []*[]any{{1}, {2}, {3}},
			funcs: []func(idx int, value *[]any, arr []*[]any) bool{
				func(idx int, value *[]any, arr []*[]any) bool {
					return len(*value) == 0
				},
			},
			expected: false,
		},
		{
			array: []*[]any{{1}, {}, {2}, {}},
			funcs: []func(idx int, value *[]any, arr []*[]any) bool{
				func(idx int, value *[]any, arr []*[]any) bool {
					return len(*value) == 0
				},
				func(idx int, value *[]any, arr []*[]any) bool {
					return len(*value) > 0
				},
			},
			expected: false,
		},
	} {
		actual := slices.All(test.array, test.funcs...)
		if actual != test.expected {
			t.Errorf("[]*[]any no. %d (%v): Got %t, expected %t", testNo+1, test, actual, test.expected)
		}
	}

	for testNo, test := range []struct {
		array    []float64
		funcs    []func(idx int, value float64, arr []float64) bool
		expected bool
	}{
		{
			array:    []float64{1.0, 2.0, 3.0},
			funcs:    []func(idx int, value float64, arr []float64) bool{},
			expected: true,
		},
		{
			array:    []float64{0.0, 0.0, 0.0},
			funcs:    []func(idx int, value float64, arr []float64) bool{},
			expected: false,
		},
		{
			array:    []float64{},
			funcs:    []func(idx int, value float64, arr []float64) bool{},
			expected: false,
		},
		{
			array: []float64{1.0, 2.0, 3.0},
			funcs: []func(idx int, value float64, arr []float64) bool{
				func(idx int, value float64, arr []float64) bool {
					return value == 0
				},
			},
			expected: false,
		},
		{
			array: []float64{1.0, 0.0, 2.0, 0.0},
			funcs: []func(idx int, value float64, arr []float64) bool{
				func(idx int, value float64, arr []float64) bool {
					return value == 0
				},
				func(idx int, value float64, arr []float64) bool {
					return value > 0
				},
			},
			expected: false,
		},
	} {
		actual := slices.All(test.array, test.funcs...)
		if actual != test.expected {
			t.Errorf("[]float64 no. %d (%v): Got %t, expected %t", testNo+1, test, actual, test.expected)
		}
	}
}
