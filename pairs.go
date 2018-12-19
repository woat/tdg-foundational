package main

import (
	"fmt"
	"reflect"
)

func solution(a []string) map[string]string {
	m := make(map[string]string)
	for _, e := range a {
		m[string(e[0])] = string(e[len(e)-1])
	}
	return m
}

func main() {
	m := solution([]string{"man", "moon", "good", "night"})
	ex := map[string]string{"m": "n", "g": "d", "n": "t"}
	ans := reflect.DeepEqual(m, ex)
	if !ans {
		fmt.Printf("FAIL: expected %v, got %v\n", ex, m)
	} else {
		fmt.Printf("PASS: got %v\n", m)
	}
}
