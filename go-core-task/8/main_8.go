package main

import (
	"fmt"
	"sync"
)

type MyWaitGroup struct {
	counter int        // Счётчик горутин
	mu      sync.Mutex // Мьютекс для защиты счётчика
	cond    *sync.Cond // Условная переменная для ожидания
}

func NewMyWaitGroup() *MyWaitGroup {
	wg := &MyWaitGroup{}
	wg.cond = sync.NewCond(&wg.mu)
	return wg
}

func (wg *MyWaitGroup) Add(delta int) {
	wg.mu.Lock()

	wg.counter += delta
	if wg.counter == 0 {
		wg.cond.Broadcast() // Если все горутины завершились, уведомляем ожидающих
	}
	if wg.counter < 0 {
		panic("counter < 0")
	}

	wg.mu.Unlock()
}

func (wg *MyWaitGroup) Done() {
	wg.Add(-1)
}

func (wg *MyWaitGroup) Wait() {
	wg.mu.Lock()
	for wg.counter > 0 {
		wg.cond.Wait() // Ожидаем завершения горутин
	}
	wg.mu.Unlock()
}

func main() {
	wg := NewMyWaitGroup()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Printf("Горутина %d начала работу\n", i)
			wg.Done()
			fmt.Printf("Горутина %d завершила работу\n", i)
		}(i)
	}

	wg.Wait()
	fmt.Println("Все горутины завершили работу")
}
