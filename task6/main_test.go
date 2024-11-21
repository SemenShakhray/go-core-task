package main

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerator(t *testing.T) {
	var numbers []int
	ch := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go Generator(ctx, ch)

	for v := range ch {
		numbers = append(numbers, v)
	}

	assert.Equal(t, 10, len(numbers))
}
