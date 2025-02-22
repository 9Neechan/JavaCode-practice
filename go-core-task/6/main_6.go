package main

import (
	"fmt"

	"math/rand"
)

func generateRandomNum(ch chan int, n int) {
	for i := 0; i < n; i++ {
		num := rand.Intn(100)
		ch <- num
	}
	close(ch)
}

// go run go-core-task/6/main_6.go
func main() {
	ch := make(chan int)
	n := 10

	go generateRandomNum(ch, n)

	for i := 0; i < n; i++ {
		num := <-ch
		fmt.Println(num)
	}
}
