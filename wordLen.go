package main

import (
	"fmt"
	"reflect"
)

func solution(a []string) map[string]int {
	m := make(map[string]int)
	for _, e := range a {
		m[e] = len(e)
	}
	return m
}

func main() {
	m := solution([]string{"a", "bb", "a", "bb"})
	ex := map[string]int{"a": 1, "bb": 2}
	ans := reflect.DeepEqual(m, ex)
	if !ans {
		fmt.Printf("FAIL: expected %v, got %v\n", ex, m)
	} else {
		fmt.Printf("PASS: got %v\n", m)
	}
}
