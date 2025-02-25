package main

import (
	"fmt"
	"sync"
)

// Реализуйте функцию, которая принимает массив строк и строку для поиска.
// Разделите массив на несколько частей и выполните поиск в каждой части параллельно.
// Верните индексы найденных строк.
func parallelSearch(strs []string, s string, n int) []int {
	indices := make([]int, 0)
	results := make(chan int, len(strs))
	chunkSize := (len(strs) + n - 1) / n 

	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(strs) {
			end = len(strs)
		}

		go func(start, end int) {
			defer wg.Done()
			for j := start; j < end; j++ {
				if strs[j] == s {
					results <- j
				}
			}
		}(start, end)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for val := range results {
		indices = append(indices, val)
	}

	return indices
}

func parallel_search_test() {
	data := []string{"apple", "banana", "cherry", "date", "fig", "grape", "banana"}
	search := "banana"
	found := parallelSearch(data, search, 3)
	fmt.Println("Found indices:", found)
}
