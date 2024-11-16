package main

import "fmt"

func SearchUniqElement(s1, s2 []string) []string {
	var result []string
	m := make(map[string]string)
	for _, v := range s2 {
		m[v] = ""
	}
	for _, v := range s1 {
		if _, ok := m[v]; ok {
			continue
		}
		result = append(result, v)
	}
	return result
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	fmt.Println(SearchUniqElement(slice1, slice2))
}
