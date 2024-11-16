package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var originalSlice []int
	l := 10
	for i := 0; i < l; i++ {
		originalSlice = append(originalSlice, rand.Intn(100)-rand.Intn(50))
	}
	fmt.Println(originalSlice)
	fmt.Println("Slice of even numbers: ", sliceExample(originalSlice))
	fmt.Println("Add numbers: ", addElements(originalSlice, 11))
	fmt.Println("Copy slice: ", copySlice(originalSlice))
	fmt.Println("Delete element: ", removeElement(originalSlice, 4))
}

func sliceExample(s []int) []int {
	var newSlice []int
	for _, v := range s {
		if v%2 == 0 {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func addElements(s []int, n int) []int {
	s = append(s, n)
	return s
}

func copySlice(s []int) []int {
	return s
}

func removeElement(s []int, n int) []int {
	if n >= len(s) {
		return nil
	}
	s = append(s[:n], s[n+1:]...)
	return s
}
