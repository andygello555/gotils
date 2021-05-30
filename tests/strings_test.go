package tests

import (
	"github.com/andygello555/gotils/strings"
	"testing"
)

func TestReplaceCharIndex(t *testing.T) {
	for _, test := range []struct{
		old            string
		indices        []int
		new            []string
		expectedOutput string
	}{
		{
			"Hello world! Hope you are having a good day today.",
			[]int{1, 2, 7, 9, 10, 14, 20, 41},
			[]string{"Fizz", "Buzz", "Bing"},
			"HFizzBuzzlo wBingrFizzBuzz! HBingpe yoFizz are having a good dBuzzy today.",
		},
		{
			"Let's replace all @s with lollipops, @ are tasty when they are @ flavour. I love @.",
			[]int{37, 63, 81},
			[]string{"lollipops"},
			"Let's replace all @s with lollipops, lollipops are tasty when they are lollipops flavour. I love lollipops.",
		},
		{
			"Hello",
			[]int{0, 1, 2, 3, 4},
			[]string{"One", "Two", "Three", "Four", "Five"},
			"OneTwoThreeFourFive",
		},
		{
			"Hello",
			[]int{},
			[]string{},
			"Hello",
		},
	} {
		newString := strings.ReplaceCharIndex(test.old, test.indices, test.new...)
		if newString != test.expectedOutput {
			t.Errorf("Got: \"%v\", expected: \"%v\"", newString, test.expectedOutput)
		}
	}
}

func TestReplaceCharIndexRange(t *testing.T) {
	for _, test := range []struct{
		old            string
		indices        [][]int
		new            []string
		expectedOutput string
	}{
		{
			"Hello world! Hope you are having a good day today.",
			[][]int{{1, 2}, {7, 9}, {10, 14}, {20, 41}},
			[]string{"Fizz", "Buzz", "Bing"},
			"HFizzllo wBuzzlBingope yoFizzay today.",
		},
		{
			"Let's replace all 'boop's with lollipops, boop are tasty when they are boop flavour. I love boop.",
			[][]int{{42, 46}, {71, 75}, {92, 96}},
			[]string{"lollipops"},
			"Let's replace all 'boop's with lollipops, lollipops are tasty when they are lollipops flavour. I love lollipops.",
		},
		{
			"Hello",
			[][]int{{0, 5}},
			[]string{"olleH"},
			"olleH",
		},
		{
			"Hello",
			[][]int{{0, 4}},
			[]string{"One", "Two", "Three", "Four", "Five"},
			"Hello",
		},
		{
			"Hello",
			[][]int{},
			[]string{},
			"Hello",
		},
	} {
		newString := strings.ReplaceCharIndexRange(test.old, test.indices, test.new...)
		if newString != test.expectedOutput {
			t.Errorf("Got: \"%v\", expected: \"%v\"", newString, test.expectedOutput)
		}
	}
}

func TestTypeName(t *testing.T) {
	for _, test := range []struct{
		i              interface{}
		expectedOutput string
	}{
		{
			"hello world",
			"string",
		},
		{
			map[string]interface{} {
				"hello": "world",
			},
			"map[string]interface {}",
		},
		{
			10,
			"int",
		},
		{
			float64(10),
			"float64",
		},
		{
			[]int{1, 2, 3},
			"[]int",
		},
		{
			[]string{"hello", "world"},
			"[]string",
		},
	} {
		typeName := strings.TypeName(test.i)
		if typeName != test.expectedOutput {
			t.Errorf("Got: \"%s\", expected: \"%s\"", typeName, test.expectedOutput)
		}
	}
}
