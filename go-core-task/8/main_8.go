package main

import (
	"fmt"
	"time"
)

// CustomWaitGroup - кастомная реализация WaitGroup на основе канала.
type CustomWaitGroup struct {
	sem chan struct{}
}

// NewCustomWaitGroup - инициализация CustomWaitGroup с количеством горутин, которые нужно ожидать.
func NewCustomWaitGroup(count int) *CustomWaitGroup {
	return &CustomWaitGroup{
		sem: make(chan struct{}, count),
	}
}

// Add - добавление новой горутины для ожидания (за счет записи в канал).
func (wg *CustomWaitGroup) Add(n int) {
	for i := 0; i < n; i++ {
		wg.sem <- struct{}{} // Записываем в канал, сигнализируя о начале работы горутины
	}
}

// Done - сигнализирует, что горутина завершила свою работу.
func (wg *CustomWaitGroup) Done() {
	<-wg.sem // Получаем из канала, сигнализируя о завершении работы горутины
}

// Wait - блокирует выполнение до тех пор, пока все горутины не завершат работу.
func (wg *CustomWaitGroup) Wait() {
	// Ждем, пока все элементы не будут получены из канала
	for len(wg.sem) > 0 {
		<-wg.sem
	}
}

func main() {
	var wg CustomWaitGroup

	// Запускаем 3 горутины
	wg.Add(3)

	go func() {
		defer wg.Done()
		fmt.Println("Горутина 1: Начало работы")
		time.Sleep(2 * time.Second)
		fmt.Println("Горутина 1: Завершение работы")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Горутина 2: Начало работы")
		time.Sleep(3 * time.Second)
		fmt.Println("Горутина 2: Завершение работы")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Горутина 3: Начало работы")
		time.Sleep(1 * time.Second)
		fmt.Println("Горутина 3: Завершение работы")
	}()

	// Ожидаем завершения всех горутин
	wg.Wait()
	fmt.Println("Все горутины завершены.")
}
