package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Создай/реализуй модель "производитель-потребитель" с использованием каналов.
// Создай одну горутину-производителя, которая генерирует данные и отправляет их в канал, 
// и несколько горутин-потребителей, которые извлокают данные из канала и обрабатывают их.

// Количество потребителей
const numConsumers = 3

// горутина-производитель, генерирует данные и отправляет их в канал
func producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик при завершении
	defer close(ch)

	for i := 0; i < 10; i++ {
		num := rand.Intn(100) // Генерируем случайное число
		fmt.Printf("[Производитель] Сгенерировал: %d\n", num)
		ch <- num                          // Отправляем в канал
		//time.Sleep(time.Millisecond * 500) // Имитация задержки
	}
	 // Закрываем канал после завершения
	fmt.Println("[Производитель] Завершил работу")
}

func consumer(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик при завершении

	for num := range ch { // Читаем из канала
		fmt.Printf("  [Потребитель %d] Обрабатывает: %d\n", id, num)
		//time.Sleep(time.Millisecond * 700) // Имитация обработки
	}

	fmt.Printf("  [Потребитель %d] Завершил работу\n", id)
}

func task6_test() {
	rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел
	ch := make(chan int)

	var wg sync.WaitGroup

	// Запускаем производителя
	wg.Add(1)
	go producer(ch, &wg)

	// Запускаем несколько потребителей
	for i := 1; i <= numConsumers; i++ {
		wg.Add(1)
		go consumer(i, ch, &wg)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()
	fmt.Println("[Главный] Все горутины завершены.")
}
