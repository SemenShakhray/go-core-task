package main

import (
	"context"
	"math/rand"
	"time"
)

func Generator(ctx context.Context, ch chan int) {

	for {
		select {
		case ch <- rand.Intn(100):
		case <-ctx.Done():
			close(ch)
			return
		}
		time.Sleep(200 * time.Millisecond)
	}
}
