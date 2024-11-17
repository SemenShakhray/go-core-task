package main

import (
	"sync"
)

func Collector(channels []chan int) chan int {
	chOut := make(chan int)
	var wg sync.WaitGroup
	for _, ch := range channels {
		wg.Add(1)
		go func(in chan int) {
			defer wg.Done()
			for n := range in {
				chOut <- n
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(chOut)
	}()
	return chOut
}
