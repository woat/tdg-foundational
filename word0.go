package main

import "fmt"
import "reflect"

func solution(a []string) map[string]int {
	m := make(map[string]int)
	for _, e := range a {
		m[e] = 0
	}
	return m
}

func main() {
	m := solution([]string{"a", "b", "a", "c"})
	ex := map[string]int{"a": 0, "b": 0, "c": 0}
	ans := reflect.DeepEqual(m, ex)
	if !ans {
		fmt.Printf("FAIL: expected %v, got %v\n", ex, m)
	} else {
		fmt.Printf("PASS: got %v\n", m)
	}
}
