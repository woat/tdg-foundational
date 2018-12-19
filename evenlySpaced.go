package main

import (
	"fmt"
	"math"
	"sort"
)

func solution(a, b, c int) bool {
	ans := []int{a, b, c}
	sort.Ints(ans)
	return math.Abs(float64(ans[0]-ans[1])) == math.Abs(float64(ans[1]-ans[2]))
}

func main() {
	ans := solution(2, 4, 6)
	if !ans {
		fmt.Printf("FAIL: expected true, got %v\n", ans)
	} else {
		fmt.Printf("PASS: got %v\n", ans)
	}
}
