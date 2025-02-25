package main

import (
	"fmt"
	"strings"
	"sync"
)

func countWords(text string, wg *sync.WaitGroup, result chan<- int) {
	defer wg.Done()
	words := strings.Fields(text)
	result <- len(words)
}

func count_words_test() {
	//text := "This is a sample text with several words. This text is for testing word count."
	parts := []string{
		"This is a sample text with several words.",
		"This text is for testing word count.",
	}

	var wg sync.WaitGroup
	result := make(chan int)

	for _, part := range parts {
		wg.Add(1)
		go countWords(part, &wg, result)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	totalWords := 0
	for count := range result {
		totalWords += count
	}

	fmt.Printf("Total words: %d\n", totalWords)
}
