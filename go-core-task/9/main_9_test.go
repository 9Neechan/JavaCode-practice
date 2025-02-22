package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {
	ch1 := make(chan uint8)
	ch2 := make(chan float64)

	go func() {
		for num := range ch1 {
			ch2 <- math.Pow(float64(num), 2)
		}
		close(ch2)
	}()

	go func() {
		for i := uint8(0); i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	// Ожидаемые результаты
	expected := []float64{0, 1, 4, 9, 16, 25, 36, 49, 64, 81}

	var results []float64
	for val := range ch2 {
		results = append(results, val)
	}

	require.Equal(t, expected, results, "Выходные значения не совпадают с ожидаемыми")
}
