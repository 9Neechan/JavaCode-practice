package main

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestCompareSlices(t *testing.T) {
	tests := []struct {
		name     string
		s1       []int
		s2       []int
		expected bool
		result   []int
	}{
		{
			name:     "Both slices are empty",
			s1:       []int{},
			s2:       []int{},
			expected: false,
			result:   []int{},
		},
		{
			name:     "First slice is empty",
			s1:       []int{},
			s2:       []int{1, 2, 3},
			expected: false,
			result:   []int{},
		},
		{
			name:     "Second slice is empty",
			s1:       []int{1, 2, 3},
			s2:       []int{},
			expected: false,
			result:   []int{},
		},
		{
			name:     "No common elements",
			s1:       []int{1, 2, 3},
			s2:       []int{4, 5, 6},
			expected: false,
			result:   []int{},
		},
		{
			name:     "Some common elements",
			s1:       []int{1, 2, 3},
			s2:       []int{2, 3, 4},
			expected: true,
			result:   []int{2, 3},
		},
		{
			name:     "Common elements with duplicates in s2",
			s1:       []int{1, 2, 3},
			s2:       []int{2, 2, 3, 3, 3},
			expected: true,
			result:   []int{2, 3},
		},
		{
			name:     "Common elements with duplicates in s1",
			s1:       []int{1, 1, 2, 2, 3},
			s2:       []int{2, 2, 3},
			expected: true,
			result:   []int{2, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, result := compareSlices(test.s1, test.s2)
			require.Equal(t, test.expected, got)
			require.ElementsMatch(t, test.result, result)
		})
	}
}
