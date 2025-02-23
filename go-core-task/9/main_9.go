package main

import (
	"fmt"
	"math"
)

// go run go-core-task/9/main_9.go
func main() {
	ch1 := make(chan uint8)
	ch2 := make(chan float64)

	go func() {
		for num := range ch1 {
			ch2 <- math.Pow(float64(num), 2)
		}
		close(ch2)
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- uint8(i)
		}
		close(ch1)
	}()

	for val := range ch2 {
		fmt.Println(val)
	}
}
