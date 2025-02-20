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
		{[]int{10, 15, 20, 25, 30}, []int{10, 20, 30}},
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
		{[]int{5, 6}, 7, []int{5, 6, 7}},
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
		{[]int{10, 20, 30}, []int{10, 20, 30}},
		{[]int{}, []int{}}, 
	}

	for _, test := range tests {
		t.Run("Testing copySlice", func(t *testing.T) {
			result := copySlice(test.input)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		input    []int
		ind      int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 4, 5}}, // Удаление элемента с индексом 2 (значение 3)
		{[]int{10, 20, 30}, 1, []int{10, 30}},          // Удаление элемента с индексом 1 (значение 20)  
		{[]int{1}, 0, []int{}},                         // Удаление единственного элемента
	}

	for _, test := range tests {
		t.Run("Testing removeElement", func(t *testing.T) {
			result := removeElement(test.input, test.ind)
			require.Equal(t, test.expected, result)
		})
	}
}
