package main

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomNum(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int // ожидаемое количество чисел в канале
	}{
		{
			name:     "Generate 5 random numbers",
			n:        5,
			expected: 5,
		},
		{
			name:     "Generate 0 random numbers",
			n:        0,
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ch := make(chan int) 

			go generateRandomNum(ch, test.n)

			var count int
			for num := range ch {
				count++
				require.True(t, num >= 0 && num < 100, "Generated number out of range: %d", num)
				if count == test.expected {
					break
				}
			}

			require.Equal(t, test.expected, count)
		})
	}
}
