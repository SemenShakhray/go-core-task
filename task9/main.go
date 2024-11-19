package main

import (
	"fmt"
)

func firstStage(sliceIn []uint8, ch chan uint8) {
	for _, n := range sliceIn {
		ch <- n
	}
	close(ch)
}

func secondStage(ch1 chan uint8, ch2 chan float64) {
	for n := range ch1 {
		s := float64(n)
		ch2 <- s * s * s
	}
	close(ch2)
}

func main() {
	sliceIn := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch1 := make(chan uint8)
	ch2 := make(chan float64)

	go firstStage(sliceIn, ch1)
	go secondStage(ch1, ch2)

	for i := range ch2 {
		fmt.Println(i)
	}

}
