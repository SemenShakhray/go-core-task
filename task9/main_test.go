package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSemaphorWait(t *testing.T) {

	cases := [][]uint8{
		[]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]uint8{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
	}
	for _, item := range cases {
		ch1 := make(chan uint8)
		ch2 := make(chan float64)

		go firstStage(item, ch1)
		go secondStage(ch1, ch2)

		var sliceOut []float64
		for i := range ch2 {
			sliceOut = append(sliceOut, i)
		}

		var expectedSlice []float64
		for _, v := range item {
			f := float64(v)
			expectedSlice = append(expectedSlice, f*f*f)
		}
		assert.Equal(t, expectedSlice, sliceOut)
	}
}
