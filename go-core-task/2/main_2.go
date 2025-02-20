package main

import "fmt"

func sliceExample(slice []int) []int {
	out := make([]int, 0)

	for _, val := range slice {
		if val%2 == 0 {
			out = append(out, val)
		}
	}
	return out
}

func addElements(slice []int, num int) []int {
	slice = append(slice, num)
	return slice
}

func copySlice(slice []int) []int {
	out := make([]int, len(slice))
	copy(out, slice)
	return out
}

func removeElement(slice []int, ind int) []int {
	slice = append(slice[:ind], slice[ind+1:]...)
	return slice
}

// go run go-core-task/2/main_2.go
func main() {
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	exampleSlice := sliceExample(originalSlice)
	fmt.Println(originalSlice)
	fmt.Println(exampleSlice)

	copiedSlice := copySlice(originalSlice)
	fmt.Println("copiedSlice до добавления эдемента в originalSlice", copiedSlice)

	originalSlice = addElements(originalSlice, 11)
	fmt.Println("После добавления 11:", originalSlice)
	fmt.Println("copiedSlice после добавления эдемента в originalSlice", copiedSlice)

	originalSlice = removeElement(originalSlice, 1)
	fmt.Println("После удаления элемента:", originalSlice)
}
