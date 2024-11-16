package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMethod(t *testing.T) {
	exampleStruct := StringIntMap{
		m: make(map[string]int),
	}

	exampleStruct.Add("home", 12)
	assert.NotEmpty(t, exampleStruct.m)

	exampleStruct.Remove("home")
	assert.Empty(t, exampleStruct.m)

	exampleStruct.Add("street", 10)
	assert.True(t, exampleStruct.Exists("street"))
	assert.False(t, exampleStruct.Exists("city"))

	assert.Equal(t, exampleStruct.m, exampleStruct.Copy())

	p, flag := exampleStruct.Get("street")
	assert.Equal(t, 10, p)
	assert.True(t, flag)

	_, flag1 := exampleStruct.Get("city")
	assert.False(t, flag1)

}
