package main

import (
	"fmt"
	"reflect"
)

func solution(a []string) map[string]int {
	m := make(map[string]int)
	for _, e := range a {
		if _, ok := m[string(e[0])]; !ok {
			m[string(e[0])] = 0
		}
		m[string(e[0])]++
	}
	return m
}

func main() {
	m := solution([]string{"a", "a", "b"})
	ex := map[string]int{"a": 2, "b": 1}
	ans := reflect.DeepEqual(m, ex)
	if !ans {
		fmt.Printf("FAIL: expected %v, got %v\n", ex, m)
	} else {
		fmt.Printf("PASS: got %v\n", m)
	}
}
