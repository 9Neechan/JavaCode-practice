package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Глобальный WaitGroup
var wg sync.WaitGroup

// Обработчик запроса: Возвращает текущее время
func timeHandler(w http.ResponseWriter, r *http.Request) {
	defer wg.Done() // Уменьшаем счетчик горутин по завершению

	currentTime := time.Now().Format("15:04:05")
	fmt.Fprintf(w, "Текущее время: %s\n", currentTime)
}

func task5() {
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		wg.Add(1) // Увеличиваем счетчик горутин
		go timeHandler(w, r) // Запускаем обработчик в отдельной горутине
	})

	server := &http.Server{Addr: ":8080"}

	fmt.Println("Сервер запущен на http://localhost:8080")
	go server.ListenAndServe() // Запускаем сервер в отдельной горутине

	// Ожидаем завершения всех горутин перед завершением программы
	wg.Wait()
	fmt.Println("Сервер завершает работу.")
}
