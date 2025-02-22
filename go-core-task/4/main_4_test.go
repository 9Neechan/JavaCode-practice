package main

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestCompareSlices(t *testing.T) {
	tests := []struct {
		name     string
		s1       []string
		s2       []string
		expected []string
	}{
		{
			name:     "Both slices are empty",
			s1:       []string{},
			s2:       []string{},
			expected: []string{},
		},
		{
			name:     "First slice is empty",
			s1:       []string{},
			s2:       []string{"apple", "banana", "cherry"},
			expected: []string{"apple", "banana", "cherry"},
		},
		{
			name:     "Second slice is empty",
			s1:       []string{"apple", "banana", "cherry"},
			s2:       []string{},
			expected: []string{},
		},
		{
			name:     "No unique elements in s2",
			s1:       []string{"apple", "banana", "cherry"},
			s2:       []string{"banana", "cherry", "apple"},
			expected: []string{},
		},
		{
			name:     "Some unique elements in s2",
			s1:       []string{"apple", "banana"},
			s2:       []string{"banana", "cherry", "date"},
			expected: []string{"cherry", "date"},
		},
		{
			name:     "Repeating elements in s2",
			s1:       []string{"apple", "banana"},
			s2:       []string{"banana", "banana", "cherry", "cherry", "apple"},
			expected: []string{"cherry"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := compareSlices(test.s1, test.s2)
			require.Equal(t, test.expected, result)
		})
	}
}
