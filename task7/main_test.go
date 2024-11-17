package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollector(t *testing.T) {
	ch := make([]chan int, 5)
	for i := 0; i < 5; i++ {
		ch[i] = make(chan int, 5)
		for j := 0; j < 5; j++ {
			ch[i] <- j
		}
		close(ch[i])

	}
	out := Collector(ch)

	var slice []int
	for i := range out {
		slice = append(slice, i)
	}
	assert.Equal(t, 25, len(slice))

	ex := []int{0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4}
	sort.Ints(slice)
	assert.Equal(t, ex, slice)
}
