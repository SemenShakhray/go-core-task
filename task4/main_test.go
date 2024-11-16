package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	expected := []string{"apple", "cherry", "43", "lead", "gno1"}

	assert.Equal(t, expected, SearchUniqElement(slice1, slice2))
}
