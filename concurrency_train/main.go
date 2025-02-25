package main

import (
	"fmt"
	"sync"
)

// Создайте программу, которая принимает массив чисел
// и создает N горутин для вычисления квадрата каждого числа.
// Результаты должны собираться в отдельный массив, который будет выведен в конце.
func task1() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := len(nums)
	results := make(chan int, n)
	out := make([]int, 0, n)

	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			results <- nums[i] * nums[i]
		}(i)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for val := range results {
		out = append(out, val)
	}

	fmt.Println(out)
}

// Создайте программу, которая запускает несколько горутин,
// каждая из которых увеличивает общий счетчик
func task2() {
	counter := 0
	const numGoroutines = 10

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	var mu sync.Mutex

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			counter++
		}()
	}

	wg.Wait()

	fmt.Println("Final counter:", counter)
}

// программу, которая разбивает массив чисел на N частей
// и использует N горутин для вычисления суммы квадрмтов каждой части.
// В конце программа должна вывести общую сумму квадратов всех частей.
func task3() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // 385
	n := 3
	chunkSize := (len(nums) + n - 1) / n
	out := 0

	var wg sync.WaitGroup
	wg.Add(n)

	sums := make(chan int, n)

	for i := 0; i < n; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(nums) {
			end = len(nums)
		}

		go func(start, end int) {
			defer wg.Done()

			sum := 0
			for j := start; j < end; j++ {
				sum += nums[j] * nums[j]
			}
			sums <- sum
		}(start, end)
	}

	go func() {
		wg.Wait()
		close(sums)
	}()

	for val := range sums {
		out += val
	}

	fmt.Println(out)
}

// go run .
func main() {
	//map_cache_test()
	//count_words_test()
	//parallel_search_test()

	//task1()
	//task2()
	// task3()

	task6_test()
	

}
