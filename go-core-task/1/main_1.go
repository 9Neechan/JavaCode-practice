package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"
)

func makeString[T any](arr []T) string {
	str_arr := make([]string, 0)
	for _, val := range arr {
		str_arr = append(str_arr, fmt.Sprintf("%v", val))
	}

	return strings.Join(str_arr, "")
}

func hashRunes(runes []rune) string {
	salt := "go-2024" 
	mid := len(runes) / 2

	// Вставляем соль в середину
	newRunes := append([]rune{}, runes[:mid]...)
	newRunes = append(newRunes, []rune(salt)...)
	newRunes = append(newRunes, runes[mid:]...)

	str := string(newRunes)
	hash := sha256.Sum256([]byte(str))

	// Возвращаем хэш в виде строки (hex)
	return hex.EncodeToString(hash[:])
}


// go run go-core-task/1/main_1.go
func main() {
	// 1) Создает несколько переменных различных типов данных:
	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатиричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64*/

	arr := []any{numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum}

	// 2) Определяет тип каждой переменной и выводит его на экран.
	for _, val := range arr {
		fmt.Printf("%v has type: %s\n", val, reflect.TypeOf(val))
	}

	// 3) Преобразует все переменные в строковый тип и объединяет их в одну строку.
	str := makeString(arr)

	// 4) Преобразовать эту строку в срез рун.
	runes := []rune(str)

	// 5) Захэшировать этот срез рун SHA256, добавив в середину соль "go-2024" и вывести результат.
	hashed := hashRunes(runes)
	fmt.Println(hashed)
}
