package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSliceExample(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{2, 4}},
		{[]int{}, []int{}}, 
	}

	for _, test := range tests {
		t.Run("Testing sliceExample", func(t *testing.T) {
			result := sliceExample(test.input)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestAddElements(t *testing.T) {
	tests := []struct {
		input    []int
		num      int
		expected []int
	}{
		{[]int{1, 2, 3}, 4, []int{1, 2, 3, 4}},
		{[]int{}, 1, []int{1}}, 
	}

	for _, test := range tests {
		t.Run("Testing addElements", func(t *testing.T) {
			result := addElements(test.input, test.num)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestCopySlice(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{}, []int{}}, 
	}

	for _, test := range tests {
		t.Run("Testing copySlice", func(t *testing.T) {
			result := copySlice(test.input)
			require.Equal(t, test.expected, result)
			require.NotSame(t, &test.input, &result)
		})
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		input    []int
		ind      int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 4, 5}}, 
		{[]int{1}, 0, []int{}},                        
	}

	for _, test := range tests {
		t.Run("Testing removeElement", func(t *testing.T) {
			result := removeElement(test.input, test.ind)
			require.Equal(t, test.expected, result)
		})
	}
}
