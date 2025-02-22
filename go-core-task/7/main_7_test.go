package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMergeChannelsInts(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 6; i <= 10; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	mergedChannel := mergeChannels([]chan int{ch1, ch2})

	var result []int
	for val := range mergedChannel {
		result = append(result, val)
	}

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	require.ElementsMatch(t, expected, result)
}

func TestMergeChannelsStrings(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for i := 1; i <= 5; i++ {
			ch1 <- strconv.Itoa(i)
		}
		close(ch1)
	}()

	go func() {
		for i := 6; i <= 10; i++ {
			ch2 <- strconv.Itoa(i)
		}
		close(ch2)
	}()

	mergedChannel := mergeChannels([]chan string{ch1, ch2})

	var result []string
	for val := range mergedChannel {
		result = append(result, val)
	}

	expected := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	require.ElementsMatch(t, expected, result)
}
