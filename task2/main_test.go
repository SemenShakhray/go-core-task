package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperation(t *testing.T) {
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	add := 11
	delete := 4

	expectedExapmple := []int{2, 4, 6, 8, 10}
	expectedAdd := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	expectedDelete := []int{1, 2, 3, 4, 6, 7, 8, 9, 10}

	assert.Equal(t, expectedExapmple, sliceExample(originalSlice))
	assert.Equal(t, expectedAdd, addElements(originalSlice, add))
	assert.Equal(t, originalSlice, copySlice(originalSlice))
	assert.Equal(t, expectedDelete, removeElement(originalSlice, delete))

}
