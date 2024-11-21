package main

import "fmt"

func SearhIntersections(a, b []int) (bool, []int) {
	m := make(map[int]string)
	var result []int
	for _, v := range a {
		m[v] = ""
	}
	for _, v := range b {
		if _, ok := m[v]; ok {
			result = append(result, v)
		}
	}
	if result == nil {
		return false, nil
	}
	return true, result
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	ok, res := SearhIntersections(a, b)
	fmt.Println(ok, res)
}
