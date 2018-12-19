package main

import (
	"fmt"
	"reflect"
)

func solution(raw, cw []string) []string {
	m := make(map[string]string, len(raw))
	ans := make([]string, 0, len(raw))
	i := 0
	for _, r := range raw {
		s := string(r)
		if _, ok := m[s]; !ok {
			m[s] = cw[i]
			i += 1
		}
		ans = append(ans, m[s])
	}
	return ans
}

func main() {
	m := solution([]string{"a", "b", "a", "d"}, []string{"1", "2", "3"})
	ex := []string{"1", "2", "1", "3"}
	ans := reflect.DeepEqual(ex, m)
	if !ans {
		fmt.Printf("FAIL: expected %v, got %v\n", ex, m)
	} else {
		fmt.Printf("PASS: got %v\n", m)
	}
}
