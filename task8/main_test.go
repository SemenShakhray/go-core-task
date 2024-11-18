package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSemaphorWait(t *testing.T) {

	cases := [][]int{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
	}

	for _, item := range cases {
		delta := len(item)

		s := newSemaphor(delta)
		ch := make(chan int)

		for _, n := range item {
			go func() {
				time.Sleep(1000 * time.Millisecond)
				ch <- n
				s.Add(1)
			}()
		}

		go func() {
			s.Done(delta)
			close(ch)
		}()

		var out []int
		for n := range ch {
			out = append(out, n)
		}
		assert.Equal(t, delta, len(out))
	}
}
