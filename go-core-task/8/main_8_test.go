package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMyWaitGroup_SingleGoroutine(t *testing.T) {
	wg := NewMyWaitGroup()
	wg.Add(1)

	go func() {
		defer wg.Done()
	}()

	wg.Wait()
}

func TestMyWaitGroup_MultipleGoroutines(t *testing.T) {
	numGoroutines := 10
	result := 0

	wg := NewMyWaitGroup()
	wg.Add(numGoroutines)

	var mu sync.Mutex

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			result++
			mu.Unlock()
		}()
	}

	wg.Wait()

	require.Equal(t, numGoroutines, result)
}

func TestMyWaitGroup_AddNegative(t *testing.T) {
	wg := NewMyWaitGroup()
	require.Panics(t, func(){ wg.Add(-1) })
}
