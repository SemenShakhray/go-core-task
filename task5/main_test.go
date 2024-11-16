package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	expected := []int{64, 3}

	ok, res := SearhIntersections(a, b)

	assert.True(t, ok)
	assert.Equal(t, expected, res)
}
