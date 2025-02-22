package main

import (
	"sync"
)

func mergeChannels[T any](chans []chan T) chan T {
	out := make(chan T)
	var wg sync.WaitGroup

	for _, ch := range chans {
		wg.Add(1)
		go func(ch chan T) {
			defer wg.Done()
			for val := range ch {
				out <- val
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
